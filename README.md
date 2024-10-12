# Satisfactory Client Go

This Go client library provides an interface to interact with the Satisfactory dedicated server HTTP API.

## Installation

To install the library, use go get:
```bash
go get github.com/idebeijer/satisfactory-client-go/satisfactory
```

## Usage

Create the client and login with the password. The password is the same as the one you use to login to the server.
```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/idebeijer/satisfactory-client-go/satisfactory"
)

func main() {
	ctx := context.Background()
	client := satisfactory.NewClient("https://localhost:7777", nil, true)

	password := os.Getenv("SF_PASSWD") // Replace with your password
	if _, err := client.PasswordLogin(ctx, "Administrator", password); err != nil {
		fmt.Println(err)
		return
	}
}
```

## Examples
### Run a health check
```go
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
```