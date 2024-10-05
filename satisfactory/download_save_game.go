package satisfactory

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DownloadSaveGameRequest represents the request body for DownloadSaveGame.
type DownloadSaveGameRequest struct {
	Function string                      `json:"function"`
	Data     DownloadSaveGameRequestData `json:"data"`
}

// DownloadSaveGameRequestData contains the data for the DownloadSaveGame request.
type DownloadSaveGameRequestData struct {
	SaveName string `json:"saveName"`
}

// DownloadSaveGame retrieves a save game file by name from the Dedicated Server.
// It returns the file data as a byte slice or an error if the operation fails.
func (c *Client) DownloadSaveGame(ctx context.Context, saveName string) ([]byte, *Response, error) {
	reqData := &DownloadSaveGameRequest{
		Function: "DownloadSaveGame",
		Data: DownloadSaveGameRequestData{
			SaveName: saveName,
		},
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=DownloadSaveGame", reqData)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, response, err
	}
	response.Body = bodyBytes

	// Handle error responses (assuming error responses are JSON)
	if resp.Header.Get("Content-Type") == "application/json" {
		var apiErr ErrorResponse
		if err := json.Unmarshal(bodyBytes, &apiErr); err == nil && apiErr.ErrorCode != "" {
			apiErr.Response = response
			return nil, response, &apiErr
		}
	}

	if resp.StatusCode == 404 {
		return nil, response, fmt.Errorf("save game '%s' not found", saveName)
	}

	if resp.StatusCode >= 400 {
		return nil, response, fmt.Errorf("HTTP error: status code %d", resp.StatusCode)
	}

	return bodyBytes, response, nil
}
