package satisfactory

import (
	"context"
	"net/http"
)

// PasswordLoginRequest represents the request body for PasswordLogin.
type PasswordLoginRequest struct {
	Function string                   `json:"function"`
	Data     PasswordLoginRequestData `json:"data"`
}

// PasswordLoginRequestData contains the data for the PasswordLogin request.
type PasswordLoginRequestData struct {
	MinimumPrivilegeLevel string `json:"minimumPrivilegeLevel"`
	Password              string `json:"password"`
}

// PasswordLogin performs a login to the server without a password.
// The `minimumPrivilegeLevel` parameter specifies the required privilege level for the login.
// Possible values for `minimumPrivilegeLevel` include:
//   - `NotAuthenticated`
//   - `Client`
//   - `Administrator`
//   - `InitialAdmin`
//   - `APIToken`
//
// If the login is successful, the authentication token is set for the client.
func (c *Client) PasswordLogin(ctx context.Context, minimumPrivilegeLevel, password string) (*Response, error) {
	reqData := &PasswordLoginRequest{
		Function: "PasswordLogin",
		Data: PasswordLoginRequestData{
			MinimumPrivilegeLevel: minimumPrivilegeLevel,
			Password:              password,
		},
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "/api/v1/?function=PasswordLogin", reqData)
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

	c.SetAuthToken(data.AuthenticationToken)

	return resp, nil
}
