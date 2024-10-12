package satisfactory

import "context"

type RunCommandRequest struct {
	Function string `json:"function"`
	Data     struct {
		Command string `json:"command"`
	}
}

// RunCommand runs a command on the server.
func (c *Client) RunCommand(ctx context.Context, command string) (*Response, error) {
	reqData := &RunCommandRequest{
		Function: "RunCommand",
	}
	reqData.Data.Command = command

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
