# echo-todo Makefile
# Lambda ZIP creation and deployment

# Variables
BINARY_NAME = bootstrap
ZIP_FILE = lambda-deployment.zip
FUNCTION_NAME = echo-todo-api
REGION = ap-northeast-1
TABLE_NAME = todos
IAM_ROLE = echo-todo-lambda-role

# Get AWS Account ID
ACCOUNT_ID := $(shell aws sts get-caller-identity --query Account --output text)

# Default target
.PHONY: all
all: zip

# Build for Lambda (Linux)
.PHONY: build-lambda
build-lambda:
	@echo "🔨 Building for Lambda (Linux)..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) cmd/server/main.go
	@echo "✅ Lambda build complete: $(BINARY_NAME)"

# Create ZIP file for Lambda deployment
.PHONY: zip
zip: build-lambda
	@echo "📦 Creating deployment ZIP..."
	@rm -f $(ZIP_FILE)
	zip $(ZIP_FILE) $(BINARY_NAME)
	@echo "✅ ZIP created: $(ZIP_FILE)"

# Create Lambda function
.PHONY: create-lambda
create-lambda: zip
	@echo "🚀 Creating Lambda function..."
	aws lambda create-function \
		--function-name $(FUNCTION_NAME) \
		--runtime provided.al2 \
		--role arn:aws:iam::$(ACCOUNT_ID):role/$(IAM_ROLE) \
		--handler $(BINARY_NAME) \
		--zip-file fileb://$(ZIP_FILE) \
		--timeout 30 \
		--memory-size 512 \
		--environment Variables='{"DYNAMODB_TABLE_NAME":"$(TABLE_NAME)","AWS_REGION":"$(REGION)","PORT":"8080"}' \
		--region $(REGION)
	@echo "✅ Lambda function created: $(FUNCTION_NAME)"

# Create Function URL
.PHONY: create-function-url
create-function-url:
	@echo "🌐 Creating Function URL..."
	aws lambda create-function-url-config \
		--function-name $(FUNCTION_NAME) \
		--auth-type NONE \
		--region $(REGION)
	@echo "📋 Function URL:"
	@aws lambda get-function-url-config \
		--function-name $(FUNCTION_NAME) \
		--region $(REGION) \
		--query 'FunctionUrl' --output text

# Deploy: ZIP + Lambda + Function URL
.PHONY: deploy
deploy: zip create-lambda create-function-url
	@echo "🎉 Deployment complete!"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -f $(ZIP_FILE)
	@echo "✅ Clean complete"