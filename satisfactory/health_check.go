package satisfactory

import (
	"context"
	"net/http"
)

// HealthCheckRequest represents the request body for the HealthCheck function.
type HealthCheckRequest struct {
	Function string                 `json:"function"`
	Data     HealthCheckRequestData `json:"data"`
}

// HealthCheckRequestData contains the data for the HealthCheck request.
type HealthCheckRequestData struct {
	ClientCustomData string `json:"clientCustomData"`
}

// HealthCheckResponseData contains the data from the HealthCheck response.
type HealthCheckResponseData struct {
	Health           string `json:"health"`
	ServerCustomData string `json:"serverCustomData"`
}

// HealthCheck performs a health check on the Dedicated Server API.
func (c *Client) HealthCheck(ctx context.Context, clientCustomData string) (*HealthCheckResponseData, *Response, error) {
	reqData := &HealthCheckRequest{
		Function: "HealthCheck",
		Data: HealthCheckRequestData{
			ClientCustomData: clientCustomData,
		},
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=HealthCheck", reqData)
	if err != nil {
		return nil, nil, err
	}

	var respData HealthCheckResponseData
	resp, err := c.Do(req, &respData)
	if err != nil {
		return nil, resp, err
	}

	return &respData, resp, nil
}
