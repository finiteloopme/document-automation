package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/finiteloopme/goutils/pkg/log"
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
	fmt.Fprint(w, fmt.Sprint(len(content)))
	log.Info(string(content))
	log.Debug("Done processing FileUpload request")
}

func PersistDocument(content []byte) {
	// ctx := context.Background()
	// client, err := pubsub.NewClient(ctx)
}
