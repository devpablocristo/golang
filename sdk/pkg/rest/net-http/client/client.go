package sdkhclnt

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/devpablocristo/golang/sdk/pkg/rest/net-http/client/ports"
)

var (
	instance  ports.Client
	once      sync.Once
	initError error
)

type client struct {
	config     ports.Config
	httpClient *http.Client
}

func newClient(config ports.Config) (ports.Client, error) {
	once.Do(func() {
		err := config.Validate()
		if err != nil {
			initError = err
			return
		}

		instance = &client{
			config:     config,
			httpClient: http.DefaultClient,
		}
	})
	return instance, initError
}

func (c *client) GetAccessToken(ctx context.Context) (string, error) {
	endpoint := c.config.GetAuthServerURL() + "/realms/" + c.config.GetRealm() + "/protocol/openid-connect/token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.config.GetClientID())
	data.Set("client_secret", c.config.GetClientSecret())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get access token")
	}

	var tokenRes struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenRes); err != nil {
		return "", err
	}

	if tokenRes.AccessToken == "" {
		return "", errors.New("access token not found in response")
	}

	return tokenRes.AccessToken, nil
}

func (c *client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}
