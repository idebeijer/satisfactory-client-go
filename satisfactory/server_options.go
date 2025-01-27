package satisfactory

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// GetServerOptionsRequest represents the request body for GetServerOptions.
type GetServerOptionsRequest struct {
	Function string `json:"function"`
}

// GetServerOptionsResponseData contains the data from the GetServerOptions response.
type GetServerOptionsResponseData struct {
	ServerOptions        ServerOptions `json:"serverOptions"`
	PendingServerOptions ServerOptions `json:"pendingServerOptions"`
}

// ServerOptions represents the server options with specific types.
type ServerOptions struct {
	AutoPause             string `json:"FG.DSAutoPause"`
	AutoSaveOnDisconnect  string `json:"FG.DSAutoSaveOnDisconnect"`
	AutosaveInterval      string `json:"FG.AutosaveInterval"`
	DisableSeasonalEvents string `json:"FG.DisableSeasonalEvents"`
	ServerRestartTimeSlot string `json:"FG.ServerRestartTimeSlot"`
	SendGameplayData      string `json:"FG.SendGameplayData"`
	NetworkQuality        string `json:"FG.NetworkQuality"`
}

// GetServerOptions retrieves the current server options from the Dedicated Server.
func (c *Client) GetServerOptions(ctx context.Context) (*GetServerOptionsResponseData, *Response, error) {
	reqData := &GetServerOptionsRequest{
		Function: "GetServerOptions",
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/", reqData)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}

	// Unmarshal the response body with strict mode
	var respData struct {
		Data GetServerOptionsResponseData `json:"data"`
	}
	decoder := json.NewDecoder(bytes.NewReader(resp.Body))
	if err := decoder.Decode(&respData); err != nil {
		return nil, resp, errors.New("failed to unmarshal JSON response: " + err.Error())
	}

	return &respData.Data, resp, nil
}

// ApplyServerOptionsRequest represents the request body for ApplyServerOptions.
type ApplyServerOptionsRequest struct {
	Function string                        `json:"function"`
	Data     ApplyServerOptionsRequestData `json:"data"`
}

// ApplyServerOptionsRequestData contains the data for the ApplyServerOptions request.
type ApplyServerOptionsRequestData struct {
	UpdatedServerOptions ServerOptions `json:"updatedServerOptions"`
}

// ApplyServerOptions applies new server options to the Dedicated Server.
func (c *Client) ApplyServerOptions(ctx context.Context, options ServerOptions) (*Response, error) {
	reqData := &ApplyServerOptionsRequest{
		Function: "ApplyServerOptions",
		Data: ApplyServerOptionsRequestData{
			UpdatedServerOptions: options,
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
