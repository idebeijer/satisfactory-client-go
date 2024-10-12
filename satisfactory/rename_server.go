package satisfactory

import (
	"context"
	"net/http"
)

type RenameServerRequest struct {
	Function string                  `json:"function"`
	Data     RenameServerRequestData `json:"data"`
}

type RenameServerRequestData struct {
	ServerName string `json:"serverName"`
}

// RenameServer renames the server.
func (c *Client) RenameServer(ctx context.Context, serverName string) (*Response, error) {
	reqData := &RenameServerRequest{
		Function: "RenameServer",
		Data: RenameServerRequestData{
			ServerName: serverName,
		},
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=RenameServer", reqData)
	if err != nil {
		return nil, err
	}

	return c.Do(req, nil)
}
