package satisfactory

import (
	"context"
	"net/http"
)

// ShutdownRequest represents the request body for Shutdown.
type ShutdownRequest struct {
	Function string `json:"function"`
}

// Shutdown shuts down the Dedicated Server.
func (c *Client) Shutdown(ctx context.Context) (*Response, error) {
	reqData := &ShutdownRequest{
		Function: "Shutdown",
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=Shutdown", reqData)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
