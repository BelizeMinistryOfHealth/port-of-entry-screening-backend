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

	ctx := context.Background()

	if err := funcframework.RegisterCloudEventFunctionContext(ctx, "/arrivalCreated", poebackend.CloudEventHandler); err != nil {
		panic(fmt.Sprintf("funcframework.RegisterEventFunctionContext failed: %v", err))
	}

	if err := funcframework.RegisterEventFunctionContext(ctx, "/arrivalCreat", poebackend.ArrivalCreated); err != nil {
		panic(fmt.Sprintf("funcframework.RegisterEventFunctionContext failed: %v", err))
	}

	if err := funcframework.Start(port); err != nil {
		panic(fmt.Sprintf("funcframework.Start failed: %v", err))
	}
}
