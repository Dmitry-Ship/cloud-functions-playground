package main

import (
	"context"
	"log"
	"os"

	functions "example.com/module"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	ctx := context.Background()

	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/getSentimentAnalysis", functions.GetSentimentAnalysis); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	// Use PORT environment variable, or default to 8081.
	port := "8081"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
