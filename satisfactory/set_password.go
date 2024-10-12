package satisfactory

import (
	"context"
	"net/http"
)

// SetPasswordRequest represents the request body for SetClientPassword and SetAdminPassword.
type SetPasswordRequest struct {
	Function string                 `json:"function"`
	Data     SetPasswordRequestData `json:"data"`
}

// SetPasswordRequestData contains the data for the SetPassword request.
type SetPasswordRequestData struct {
	Password string `json:"password"`
}

// SetClientPassword sets the password for the client.
func (c *Client) SetClientPassword(ctx context.Context, password string) (*Response, error) {
	reqData := &SetPasswordRequest{
		Function: "SetClientPassword",
		Data: SetPasswordRequestData{
			Password: password,
		},
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/", reqData)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SetAdminPassword updates the currently set Admin Password. This will invalidate all previously issued `Client` and `Admin` authentication tokens
func (c *Client) SetAdminPassword(ctx context.Context, password string) (*Response, error) {
	reqData := &SetPasswordRequest{
		Function: "SetAdminPassword",
		Data: SetPasswordRequestData{
			Password: password,
		},
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/", reqData)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
