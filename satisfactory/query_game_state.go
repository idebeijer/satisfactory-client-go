package satisfactory

import (
	"context"
	"net/http"
)

// QueryServerStateRequest represents the request body for QueryServerState.
type QueryServerStateRequest struct {
	Function string `json:"function"`
}

// ServerGameState contains the state of the server.
type ServerGameState struct {
	ActiveSessionName   string  `json:"activeSessionName"`
	NumConnectedPlayers int     `json:"numConnectedPlayers"`
	PlayerLimit         int     `json:"playerLimit"`
	TechTier            int     `json:"techTier"`
	ActiveSchematic     string  `json:"activeSchematic"`
	GamePhase           string  `json:"gamePhase"`
	IsGameRunning       bool    `json:"isGameRunning"`
	TotalGameDuration   int     `json:"totalGameDuration"`
	IsGamePaused        bool    `json:"isGamePaused"`
	AverageTickRate     float64 `json:"averageTickRate"`
	AutoLoadSessionName string  `json:"autoLoadSessionName"`
}

// QueryServerState retrieves the current state of the Dedicated Server.
func (c *Client) QueryServerState(ctx context.Context) (*ServerGameState, *Response, error) {
	reqData := &QueryServerStateRequest{
		Function: "QueryServerState",
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=QueryServerState", reqData)
	if err != nil {
		return nil, nil, err
	}

	var data struct {
		ServerGameState ServerGameState `json:"serverGameState"`
	}

	resp, err := c.Do(req, &data)
	if err != nil {
		return nil, resp, err
	}

	return &data.ServerGameState, resp, nil
}
