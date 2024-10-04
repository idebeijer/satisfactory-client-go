package satisfactory

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	logger     *log.Logger
	token      string
}

func NewClient(baseURL string, logger *log.Logger, insecureSkipVerify bool) *Client {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			},
		},
	}
	if logger == nil {
		logger = log.Default()
	}
	return &Client{
		baseURL:    parsedURL,
		httpClient: httpClient,
		logger:     logger,
	}
}

// Response represents an API response.
type Response struct {
	*http.Response
	Body []byte
}

// SetAuthToken sets the authentication token for the client.
func (c *Client) SetAuthToken(token string) {
	c.token = token
}

// NewRequest creates an HTTP request with the appropriate headers and authentication.
func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	u := c.baseURL.ResolveReference(rel)

	var buf io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	return req, nil
}

// CommonResponse represents the common structure of API responses, which can contain either data or an error.
type CommonResponse struct {
	Data         json.RawMessage `json:"data"`
	ErrorCode    string          `json:"errorCode"`
	ErrorMessage string          `json:"errorMessage"`
	ErrorData    interface{}     `json:"errorData"`
}

// Do sends an HTTP request and decodes the response into the provided output interface.
// It returns the response and any error encountered.
func (c *Client) Do(req *http.Request, out interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response.Body = bodyBytes

	// Handle empty response body
	if len(bodyBytes) == 0 {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			// Success status code but empty body
			return response, nil
		} else {
			// Error status code and empty body
			return response, fmt.Errorf("empty response with status code %d", resp.StatusCode)
		}
	}

	// Unmarshal into CommonResponse to check for API-level errors
	var commonResp CommonResponse
	if err := json.Unmarshal(bodyBytes, &commonResp); err != nil {
		return response, err
	}

	// Check if the response contains an error
	if commonResp.ErrorCode != "" {
		apiErr := &ErrorResponse{
			Response:     response,
			ErrorCode:    commonResp.ErrorCode,
			ErrorMessage: commonResp.ErrorMessage,
			ErrorData:    commonResp.ErrorData,
		}
		return response, apiErr
	}

	// If out is nil or response is No Content, return early
	if out == nil || resp.StatusCode == http.StatusNoContent {
		return response, nil
	}

	// Unmarshal the data into out
	if err := json.Unmarshal(commonResp.Data, out); err != nil {
		return response, err
	}

	return response, nil
}
