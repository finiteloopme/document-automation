package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	log "github.com/finiteloopme/goutils/pkg/log"
)

func TestHandleDefaultRequest(t *testing.T) {
	// Create a simple request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleDefaultRequest)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("returned wrong status code.  Got %v, was expecting %v", status, http.StatusOK)
	}

	expectedMsg := "Hello World"
	if rr.Body.String() != expectedMsg {
		t.Errorf("Returned wrong message. Got %v, was expecting %v", rr.Body.String(), expectedMsg)
	}
}

func TestHandleFileUpload(t *testing.T) {

	fileReader, err := os.Open("./data/au_notice_of_assessment_sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer fileReader.Close()

	content, err := ioutil.ReadAll(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	// Create a FileUpload request
	req, err := http.NewRequest("POST", "/upload", bytes.NewBuffer(content))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleFileUpload)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("returned wrong status code.  Got %v, was expecting %v", status, http.StatusOK)
	}

	expectedMsg := fmt.Sprint(len(content))
	if rr.Body.String() != expectedMsg {
		t.Errorf("Returned wrong message. Got %v, was expecting %v", rr.Body.String(), expectedMsg)
	}
}
