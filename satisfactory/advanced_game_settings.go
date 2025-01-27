package satisfactory

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// GetAdvancedGameSettingsRequest represents the request body for GetAdvancedGameSettings.
type GetAdvancedGameSettingsRequest struct {
	Function string `json:"function"`
}

// GetAdvancedGameSettingsResponseData contains the data from the GetAdvancedGameSettings response.
type GetAdvancedGameSettingsResponseData struct {
	CreativeModeEnabled  bool                 `json:"creativeModeEnabled"`
	AdvancedGameSettings AdvancedGameSettings `json:"advancedGameSettings"`
}

// AdvancedGameSettings represents the advanced game settings with specific types.
type AdvancedGameSettings struct {
	NoPower                         string `json:"FG.GameRules.NoPower,omitempty"`
	StartingTier                    string `json:"FG.GameRules.StartingTier,omitempty"`
	DisableArachnidCreatures        string `json:"FG.GameRules.DisableArachnidCreatures,omitempty"`
	NoUnlockCost                    string `json:"FG.GameRules.NoUnlockCost,omitempty"`
	SetGamePhase                    string `json:"FG.GameRules.SetGamePhase,omitempty"`
	UnlockAllResearchSchematics     string `json:"FG.GameRules.UnlockAllResearchSchematics,omitempty"`
	UnlockInstantAltRecipes         string `json:"FG.GameRules.UnlockInstantAltRecipes,omitempty"`
	UnlockAllResourceSinkSchematics string `json:"FG.GameRules.UnlockAllResourceSinkSchematics,omitempty"`

	NoBuildCost string `json:"FG.PlayerRules.NoBuildCost,omitempty"`
	GodMode     string `json:"FG.PlayerRules.GodMode,omitempty"`
	FlightMode  string `json:"FG.PlayerRules.FlightMode,omitempty"`
}

// GetAdvancedGameSettings retrieves the advanced game settings of the Dedicated Server.
func (c *Client) GetAdvancedGameSettings(ctx context.Context) (*GetAdvancedGameSettingsResponseData, *Response, error) {
	reqData := &GetAdvancedGameSettingsRequest{
		Function: "GetAdvancedGameSettings",
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
		Data GetAdvancedGameSettingsResponseData `json:"data"`
	}
	decoder := json.NewDecoder(bytes.NewReader(resp.Body))
	if err := decoder.Decode(&respData); err != nil {
		return nil, resp, errors.New("failed to unmarshal JSON response: " + err.Error())
	}

	return &respData.Data, resp, nil
}

// ApplyAdvancedGameSettingsRequest represents the request body for ApplyAdvancedGameSettings.
type ApplyAdvancedGameSettingsRequest struct {
	Function string                               `json:"function"`
	Data     ApplyAdvancedGameSettingsRequestData `json:"data"`
}

// ApplyAdvancedGameSettingsRequestData contains the data for the ApplyAdvancedGameSettings request.
type ApplyAdvancedGameSettingsRequestData struct {
	//UpdatedAdvancedGameSettings AdvancedGameSettings `json:"advancedGameSettings"`
	UpdatedAdvancedGameSettings AdvancedGameSettings `json:"appliedAdvancedGameSettings"`
}

// ApplyAdvancedGameSettings applies new advanced game settings to the Dedicated Server.
func (c *Client) ApplyAdvancedGameSettings(ctx context.Context, options ApplyAdvancedGameSettingsRequestData) (*Response, error) {
	reqData := &ApplyAdvancedGameSettingsRequest{
		Function: "ApplyAdvancedGameSettings",
		Data:     options,
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
