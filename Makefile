generate:
	@echo "Generating..."
	@go generate ./...
GCLOUD_PROJECT ?= $(shell gcloud config get-value project)

deploy:
	@echo "Deploying..."
	@go mod tidy
	@gcloud run deploy --source . petstore --region us-south1 --project $(GCLOUD_PROJECT)

url:
	@echo "Getting URL..."
	@gcloud run services describe petstore --platform managed --format 'value(status.url)' --region us-south1 --project $(GCLOUD_PROJECT)