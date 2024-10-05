package satisfactory

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// EnumerateSessionsRequest represents the request body for EnumerateSessions.
type EnumerateSessionsRequest struct {
	Function string `json:"function"`
}

// EnumerateSessionsResponse represents the response from EnumerateSessions.
type EnumerateSessionsResponse struct {
	Data EnumerateSessionsResponseData `json:"data"`
}

// EnumerateSessionsResponseData contains the data from the EnumerateSessions response.
type EnumerateSessionsResponseData struct {
	Sessions            []SessionSaveStruct `json:"sessions"`
	CurrentSessionIndex int                 `json:"currentSessionIndex"`
}

// SessionSaveStruct represents a session and its associated save game files.
type SessionSaveStruct struct {
	SessionName string       `json:"sessionName"`
	SaveHeaders []SaveHeader `json:"saveHeaders"`
}

// SaveHeader represents the metadata for a save game file.
type SaveHeader struct {
	SaveName              string     `json:"saveName"`
	SaveDateTime          CustomTime `json:"saveDateTime"`
	SaveVersion           int        `json:"saveVersion"`
	BuildVersion          int        `json:"buildVersion"`
	SaveLocationInfo      string     `json:"saveLocationInfo"`
	MapName               string     `json:"mapName"`
	MapOptions            string     `json:"mapOptions"`
	SessionName           string     `json:"sessionName"`
	PlayDurationSeconds   int        `json:"playDurationSeconds"`
	IsModdedSave          bool       `json:"isModdedSave"`
	IsEditedSave          bool       `json:"isEditedSave"`
	IsCreativeModeEnabled bool       `json:"isCreativeModeEnabled"`
}

// EnumerateSessions retrieves the list of sessions and their associated save game files.
func (c *Client) EnumerateSessions(ctx context.Context) (*EnumerateSessionsResponseData, *Response, error) {
	reqData := &EnumerateSessionsRequest{
		Function: "EnumerateSessions",
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/", reqData)
	if err != nil {
		return nil, nil, err
	}

	var respData EnumerateSessionsResponse
	resp, err := c.Do(req, &respData)
	if err != nil {
		return nil, resp, err
	}

	// Unmarshal the response body with strict mode
	decoder := json.NewDecoder(bytes.NewReader(resp.Body))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&respData); err != nil {
		return nil, resp, errors.New("failed to unmarshal JSON response: " + err.Error())
	}

	return &respData.Data, resp, nil
}
