#!/bin/bash

# Echo TODO Lambda Deployment Script

set -e

# Variables
BINARY_NAME="bootstrap"
ZIP_FILE="lambda-deployment.zip"
FUNCTION_NAME="echo-todo-api"
REGION="ap-northeast-1"
TABLE_NAME="todos"
IAM_ROLE="echo_todo_dynamodb_role"

# Get AWS Account ID
ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)

echo "🚀 Echo TODO Lambda Deployment"
echo "==============================="
echo "Function Name: $FUNCTION_NAME"
echo "Region: $REGION"
echo "Table Name: $TABLE_NAME"
echo "IAM Role: $IAM_ROLE"
echo "Account ID: $ACCOUNT_ID"
echo ""

# Step 1: Build for Lambda
echo "🔨 Building for Lambda (Linux)..."
GOOS=linux GOARCH=amd64 go build -o $BINARY_NAME cmd/lambda/main.go
echo "✅ Lambda build complete: $BINARY_NAME"

# Step 2: Create ZIP
echo "📦 Creating deployment ZIP..."
rm -f $ZIP_FILE
zip $ZIP_FILE $BINARY_NAME
echo "✅ ZIP created: $ZIP_FILE"

# Step 3: Create Lambda function
echo "🚀 Creating Lambda function..."
aws lambda create-function \
    --function-name $FUNCTION_NAME \
    --runtime provided.al2 \
    --role arn:aws:iam::$ACCOUNT_ID:role/$IAM_ROLE \
    --handler $BINARY_NAME \
    --zip-file fileb://$ZIP_FILE \
    --timeout 30 \
    --memory-size 512 \
    --environment "Variables={DYNAMODB_TABLE_NAME=$TABLE_NAME,PORT=8080}" \
    --region $REGION

echo "✅ Lambda function created: $FUNCTION_NAME"

# Step 4: Create Function URL
echo "🌐 Creating Function URL..."
aws lambda create-function-url-config \
    --function-name $FUNCTION_NAME \
    --auth-type NONE \
    --region $REGION

# Step 5: Add public access permission
echo "🔓 Adding public access permission..."
aws lambda add-permission \
    --function-name $FUNCTION_NAME \
    --statement-id AllowPublicInvoke \
    --action lambda:InvokeFunctionUrl \
    --principal "*" \
    --function-url-auth-type NONE \
    --region $REGION

# Step 6: Get Function URL
echo "📋 Function URL:"
FUNCTION_URL=$(aws lambda get-function-url-config \
    --function-name $FUNCTION_NAME \
    --region $REGION \
    --query 'FunctionUrl' --output text)

echo $FUNCTION_URL

echo ""
echo "🎉 Deployment complete!"
echo "==============================="
echo "Function URL: $FUNCTION_URL"
echo ""
echo "🧪 Test your API:"
echo "curl $FUNCTION_URL"
echo "curl -X POST $FUNCTION_URL/api/v1/todos -H 'Content-Type: application/json' -d '{\"title\":\"Test TODO\"}'"
echo ""
echo "🧹 Clean up:"
echo "rm -f $BINARY_NAME $ZIP_FILE"