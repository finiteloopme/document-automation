data "google_project" "project" {
    project_id  = var.project-id
    # Project number
    # data.google_project.project.number
}

##########################################
# 1. Manage required APIs
module "project-services" {
  source        = "terraform-google-modules/project-factory/google//modules/project_services"
  version       = "10.1.1"

  project_id    = var.project-id
  # APIs to activate
  activate_apis = var.required-gcp-apis
}

##########################################
# 2. Manage Pubsub Topic
module "pubsub" {
  source  = "terraform-google-modules/pubsub/google"
  version = "~> 1.8"

  project_id    = var.project-id
  # Topic to create
  topic         = var.pubsub-topic
}

##########################################
# 3. Registry for managing build artifacts
resource "google_artifact_registry_repository" "doc-automation-rego" {
  provider = google-beta

  project = var.project-id

  location = var.gcp-region
  repository_id = var.repo-name
  description = "Registry repo for managing doc-automation build artifacts"
  format = "DOCKER"
}

##########################################
# 4. Manage bucket for logs
module "bucket" {
  source  = "terraform-google-modules/cloud-storage/google//modules/simple_bucket"
  version = "~> 1.3"

  name       = var.gcs-bucket
  project_id = var.project-id
  location   = var.gcp-region
  force_destroy = true
}
##########################################
# 5. IAM Policies
module "project-iam-bindings" {
  source   = "terraform-google-modules/iam/google//modules/projects_iam"
  projects = [var.project-id]
  mode     = "additive"
  bindings = {
    "roles/run.developer" = [
      "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com",
    ]
  }
}
