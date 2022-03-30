export GCP_PROJECT=kunal-scratch
export GCP_REGION=us-central1
export REPO_NAME=doc-automation
export FILE_HANDLER_SERVICE_NAME=upload-doc
export BUILD_TEST_BUCKET_NAME=${GCP_PROJECT}-build-test
export PUBSUB_TOPIC=document-uploaded-topic
export INFRA=infra
export TF_PLAN=${GCP_PROJECT}-infra.tfplan

UPLOAD_DOC=upload-doc

init: 
	cd ${INFRA}; make init

infra-build:
	cd ${INFRA}; make build

infra-clean: 
	cd ${INFRA}; make clean

app-deploy:
	cd ${UPLOAD_DOC}; make cloud-deploy