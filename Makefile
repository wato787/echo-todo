# echo-todo Makefile
# Lambda ZIP creation

# Variables
BINARY_NAME = bootstrap
ZIP_FILE = lambda-deployment.zip

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

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -f $(ZIP_FILE)
	@echo "✅ Clean complete"