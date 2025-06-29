package main

import (
	"os"
)

type Config struct {
	TableName string
	Port      string
	AWSRegion string
}

func LoadConfig() Config {
	config := Config{
		TableName: getEnvWithDefault("DYNAMODB_TABLE_NAME", "todos"),
		Port:      getEnvWithDefault("PORT", "1323"),
		AWSRegion: getEnvWithDefault("AWS_REGION", "us-east-1"),
	}
	return config
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}