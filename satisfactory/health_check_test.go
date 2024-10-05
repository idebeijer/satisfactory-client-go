package satisfactory

import (
	"context"
	"testing"
)

func TestClient_HealthCheck_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	c := NewClient("https://localhost:7777", nil, true)
	respData, _, err := c.HealthCheck(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}

	if respData.Health != "healthy" {
		t.Fatalf("expected health to be 'healthy', got %q", respData.Health)
	}

	t.Logf("HealthCheck response: %+v", respData)
}
