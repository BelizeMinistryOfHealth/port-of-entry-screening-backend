package main

import (
	"bz.moh.epi/poebackend"
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"os"
)

func main() {
	port := "8080"
	// Use PORT env variable, or default to 8080
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	poebackend.GetServer().BackendBaseURL = fmt.Sprintf("http://localhost:%s", port)
	ctx := context.Background()

	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/hello", poebackend.HandlerEcho); err != nil {
		panic(fmt.Sprintf("funcframework.RegisterHTTPFunctionContext failed: %v", err))
	}

	if err := funcframework.Start(port); err != nil {
		panic(fmt.Sprintf("funcframework.Start: %v", err))
	}
}
