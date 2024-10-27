generate:
	@echo "Generating..."
	@go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go generate ./...
deploy:
	@echo "Deploying..."
	@go mod tidy
	@gcloud run deploy --source . petstore --region us-south1
