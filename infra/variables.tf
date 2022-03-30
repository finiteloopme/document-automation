variable "project-id" {
    type = string
}

variable "gcp-region"{
    type = string
    default = "us-central1"
}

variable "pubsub-topic" {
    type = string
    default = "document-uploaded-topic"
  
}

variable "repo-name" {
    type = string
    default = "repo-name"
}

variable "gcs-bucket" {
    type = string
    default = "gcs-bucket-4-logs"

}

variable "required-gcp-apis" {
    type = list(string)
    default = [
        "run.googleapis.com",
        "cloudbuild.googleapis.com",
        "documentai.googleapis.com",
        "pubsub.googleapis.com",
        "sourcerepo.googleapis.com",
        "storage-api.googleapis.com"
    ]
}