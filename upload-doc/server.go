package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/finiteloopme/goutils/pkg/log"
	os "github.com/finiteloopme/goutils/pkg/os"

	"cloud.google.com/go/pubsub"
)

func StartServer(hostname, port, serviceName string) {
	log.Info("Starting service: " + serviceName)
	// Register the functions to handle requests
	http.HandleFunc("/", HandleDefaultRequest)
	http.HandleFunc("/upload", HandleFileUpload)
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
	if os.ReadEnvVarOptional("GOOGLE_CLOUD_PROJECT") != "" {
		// Runtime is GCP
		log.Info("Persisting the document...")
		go PersistDocument(content)
	}
	fmt.Fprint(w, fmt.Sprint(len(content)))
	log.Info(string(content))
	log.Debug("Done processing FileUpload request")
}

func PersistDocument(content []byte) {
	pubsubTopic := os.ReadEnvVar("PUBSUB_TOPIC")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.ReadEnvVar("GOOGLE_CLOUD_PROJECT"))
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
