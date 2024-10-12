package satisfactory

import (
	"context"
	"net/http"
)

// PasswordlessLoginRequest represents the request body for PasswordlessLogin.
type PasswordlessLoginRequest struct {
	Function string                       `json:"function"`
	Data     PasswordlessLoginRequestData `json:"data"`
}

// PasswordlessLoginRequestData contains the data for the PasswordlessLogin request.
type PasswordlessLoginRequestData struct {
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel"`
}

// PasswordlessLogin performs a login to the server without a password.
// The `minimumPrivilegeLevel` parameter specifies the required privilege level for the login.
// Possible values for `minimumPrivilegeLevel` include:
//   - `NotAuthenticated`
//   - `Client`
//   - `Administrator`
//   - `InitialAdmin`
//   - `APIToken`
//
// If the login is successful, the authentication token is set for the client.
func (c *Client) PasswordlessLogin(ctx context.Context, minimumPrivilegeLevel string) (*Response, error) {
	reqData := &PasswordlessLoginRequest{
		Function: "PasswordlessLogin",
		Data: PasswordlessLoginRequestData{
			MinimumPrivilegeLevel: minimumPrivilegeLevel,
		},
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=PasswordlessLogin", reqData)
	if err != nil {
		return nil, err
	}

	var data struct {
		AuthenticationToken string `json:"authenticationToken"`
	}

	resp, err := c.Do(req, &data)
	if err != nil {
		return resp, err
	}

	// Set the authentication token for future requests.
	c.SetAuthToken(data.AuthenticationToken)

	return resp, nil
}
