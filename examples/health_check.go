//go:build example

package main

import (
	"context"
	"fmt"

	"github.com/idebeijer/satisfactory-client-go/satisfactory"
)

func main() {
	ctx := context.Background()
	client := satisfactory.NewClient("https://localhost:7777", nil, true)

	healthcheck, _, err := client.HealthCheck(ctx, "Custom data from client")
	if err != nil {
		return
	}
	if healthcheck.Health != "healthy" {
		fmt.Println("Healthcheck failed")
		return
	}
}
