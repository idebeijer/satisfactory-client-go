package satisfactory

import "context"

type SaveGameRequest struct {
	Function string `json:"function"`
	Data     struct {
		SaveName string `json:"saveName"`
	}
}

// SaveGame saves the game with the specified save name.
func (c *Client) SaveGame(ctx context.Context, saveName string) (*Response, error) {
	reqData := &SaveGameRequest{
		Function: "SaveGame",
	}
	reqData.Data.SaveName = saveName

	req, err := c.NewRequest(ctx, "POST", "/api/v1/", reqData)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
