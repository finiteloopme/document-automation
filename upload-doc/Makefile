dummy:

init:
	go mod tidy
	go mod vendor

test: init
	go test

build:
	go build -o bin/document-automation

# run: 
# 	go run .

clean:
	rm -fr bin

cloud-deploy:
	gcloud builds submit --config=cloudbuild.yaml --substitutions=_GCP_PROJECT=${GCP_PROJECT},_SERVICE_NAME=${FILE_HANDLER_SERVICE_NAME},_REPO_NAME=${REPO_NAME},_GCP_REGION=${GCP_REGION},_BUCKET_NAME=${BUILD_TEST_BUCKET_NAME},_PUBSUB_TOPIC=${PUBSUB_TOPIC}