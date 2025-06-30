# SAM Makefile for Echo TODO

.PHONY: build deploy clean logs test-local

# Build SAM application
build:
	@echo "🔨 Building SAM application..."
	sam build

# Deploy SAM application  
deploy: build
	@echo "🚀 Deploying SAM application..."
	sam deploy

# Deploy with guided setup (first time)
deploy-guided: build
	@echo "🚀 Deploying SAM application (guided)..."
	sam deploy --guided

# Test locally
local: build
	@echo "🧪 Starting local API..."
	sam local start-api --host 0.0.0.0 --port 3000

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf .aws-sam/

# View logs
logs:
	@echo "📜 Viewing Lambda logs..."
	sam logs --stack-name echo-todo --tail

# Get Function URL
get-url:
	@echo "📋 Getting Function URL..."
	@aws cloudformation describe-stacks \
		--stack-name echo-todo \
		--query 'Stacks[0].Outputs[?OutputKey==`EchoTodoApi`].OutputValue' \
		--output text --region ap-northeast-1

# Test deployed API
test-api:
	$(eval FUNCTION_URL := $(shell aws cloudformation describe-stacks --stack-name echo-todo --query 'Stacks[0].Outputs[?OutputKey==`EchoTodoApi`].OutputValue' --output text --region ap-northeast-1))
	@echo "🧪 Testing deployed API..."
	@echo "Function URL: $(FUNCTION_URL)"
	curl $(FUNCTION_URL)health
	@echo ""
	curl -X POST $(FUNCTION_URL)api/v1/todos \
		-H "Content-Type: application/json" \
		-d '{"title":"SAM Test TODO","description":"Created via SAM"}'

# Delete stack
delete:
	@echo "🗑️ Deleting SAM stack..."
	sam delete --stack-name echo-todo --region ap-northeast-1

# Help
help:
	@echo "🔧 SAM Commands:"
	@echo "  build          - Build SAM application"
	@echo "  deploy         - Deploy to AWS"
	@echo "  deploy-guided  - Deploy with guided setup (first time)"
	@echo "  local          - Start local API server"
	@echo "  clean          - Clean build artifacts"
	@echo "  logs           - View Lambda logs"
	@echo "  get-url        - Get Function URL"
	@echo "  test-api       - Test deployed API"
	@echo "  delete         - Delete stack"