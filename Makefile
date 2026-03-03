generate:
	@echo "Generating..."
	@go generate ./...
deploy:
	@echo "Deploying..."
	@go mod tidy
	@gcloud run deploy --source . petstore --region us-south1
