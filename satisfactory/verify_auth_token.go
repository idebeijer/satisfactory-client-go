package satisfactory

import (
	"context"
	"net/http"
)

// VerifyAuthenticationTokenRequest represents the request body for VerifyAuthenticationToken.
type VerifyAuthenticationTokenRequest struct {
	Function string `json:"function"`
}

// VerifyAuthenticationToken verifies the authentication token provided to the API.
func (c *Client) VerifyAuthenticationToken(ctx context.Context) (*Response, error) {
	reqData := &VerifyAuthenticationTokenRequest{
		Function: "VerifyAuthenticationToken",
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=VerifyAuthenticationToken", reqData)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
