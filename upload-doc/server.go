package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	gcp "github.com/finiteloopme/goutils/pkg/gcp"
	log "github.com/finiteloopme/goutils/pkg/log"
	os "github.com/finiteloopme/goutils/pkg/os"

	documentai "cloud.google.com/go/documentai/apiv1"
	"cloud.google.com/go/pubsub"
	documentaipb "google.golang.org/genproto/googleapis/cloud/documentai/v1"
)

func StartServer(hostname, port, serviceName string) {
	log.Info("Starting service: " + serviceName)
	// Register the functions to handle requests
	http.HandleFunc("/", HandleDefaultRequest)
	http.HandleFunc("/upload", HandleFileUpload)
	http.HandleFunc("/process", ProcessDocument)
	// Start the service
	log.Info("\tListening on port: " + port)
	if err := http.ListenAndServe(hostname+":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func HandleDefaultRequest(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling the request")
	fmt.Fprintf(w, "Hello World")
	log.Debug("Done processing the request")
}

func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
	log.Debug("Handling the FileUpload request")
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "404 not found. ", http.StatusNotFound)
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Check if the runtime is GCP
	// if ReadEnvVarOptional("GOOGLE_CLOUD_PROJECT") != "" {
	if gcp.IsRuntimeGCP() {
		// Runtime is GCP
		log.Info("Persisting the document...")
		go PersistDocument(content)
	}
	fmt.Fprint(w, fmt.Sprint(len(content)))
	log.Debug(string(content))
	log.Info("Done processing FileUpload request")
}

func PersistDocument(content []byte) {
	pubsubTopic := os.ReadEnvVar("PUBSUB_TOPIC")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, gcp.GetProjectID())
	if err != nil {
		log.Fatal(err)
	}
	topic := client.Topic(pubsubTopic)

	if _, err := topic.Publish(ctx, &pubsub.Message{
		Data: content,
	}).Get(ctx); err != nil {
		log.Fatal(err)
	}
	log.Info("Document persisted")
}

func ProcessDocument(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	c, err := documentai.NewDocumentProcessorClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	if r.Method != "POST" {
		http.Error(w, "404 not found. ", http.StatusNotFound)
	}
	// TODO: handle error
	content, _ := ioutil.ReadAll(r.Body)
	req := &documentaipb.ProcessRequest{
		Source: &documentaipb.ProcessRequest_RawDocument{
			RawDocument: &documentaipb.RawDocument{
				Content:  content,
				MimeType: "application/pdf",
			}},
		Name:            "projects/550614207330/locations/us/processors/7b15c915fa37b415",
		SkipHumanReview: true,
	}

	resp, err := c.ProcessDocument(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	labels := make(map[string]string)
	for _, doc_page := range resp.GetDocument().GetPages() {
		for _, field := range doc_page.GetFormFields() {
			labels[field.GetFieldName().GetTextAnchor().GetContent()] = field.GetFieldValue().TextAnchor.GetContent()
			log.Debug("Field name: " + field.GetFieldName().GetTextAnchor().GetContent())
			log.Debug("Field value: " + field.GetFieldValue().TextAnchor.GetContent())
		}
	}
	jsonBytes, _ := json.Marshal(labels)
	log.Info(string(jsonBytes))
	log.Info("Finished processing document")

}
