package satisfactory

import (
	"context"
	"net/http"
)

// ClaimServerRequest represents the request body for ClaimServer.
type ClaimServerRequest struct {
	Function string                 `json:"function"`
	Data     ClaimServerRequestData `json:"data"`
}

// ClaimServerRequestData contains the data for the ClaimServer request.
type ClaimServerRequestData struct {
	// ServerName is the new name given to the server.
	ServerName string `json:"serverName"`
	// ServerPassword is the password to set for the server.
	AdminPassword string `json:"adminPassword"`
}

// ClaimServer attempts to claim a server by giving it a name and setting the admin password.
func (c *Client) ClaimServer(ctx context.Context, serverName, adminPassword string) (*Response, error) {
	reqData := &ClaimServerRequest{
		Function: "ClaimServer",
		Data: ClaimServerRequestData{
			ServerName:    serverName,
			AdminPassword: adminPassword,
		},
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=ClaimServer", reqData)
	if err != nil {
		return nil, err
	}

	return c.Do(req, nil)
}
