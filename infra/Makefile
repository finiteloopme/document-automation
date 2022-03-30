dummy:

enable-compute:
	gcloud services enable compute.googleapis.com

clean: enable-compute
	terraform destroy -var="project-id=${GCP_PROJECT}" -auto-approve
	rm -fr .terraform*; rm *.tfplan; rm terraform.tfstate* 

init: enable-compute
	terraform init
	terraform plan -out=${TF_PLAN} -var="project-id=${GCP_PROJECT}" -var="pubsub-topic=${PUBSUB_TOPIC}" -var="repo-name=${REPO_NAME}" -var="gcs-bucket=${BUILD_TEST_BUCKET_NAME}"

build:
	terraform apply ${TF_PLAN}
