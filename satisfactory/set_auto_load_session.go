package satisfactory

import (
	"context"
	"net/http"
)

type SetAutoLoadSessionRequest struct {
	Function string                        `json:"function"`
	Data     SetAutoLoadSessionRequestData `json:"data"`
}

type SetAutoLoadSessionRequestData struct {
	SessionName bool `json:"sessionName"`
}

// SetAutoLoadSession sets the autoload session for the client.
func (c *Client) SetAutoLoadSession(ctx context.Context, sessionName bool) (*Response, error) {
	reqData := &SetAutoLoadSessionRequest{
		Function: "SetAutoLoadSession",
		Data: SetAutoLoadSessionRequestData{
			SessionName: sessionName,
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
