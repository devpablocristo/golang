package sdkhclnt

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/rest/net-http/client/ports"
)

type config struct {
	authServerURL string
	realm         string
	clientID      string
	clientSecret  string
}

func newConfig(authServerURL, realm, clientID, clientSecret string) ports.Config {
	return &config{
		authServerURL: authServerURL,
		realm:         realm,
		clientID:      clientID,
		clientSecret:  clientSecret,
	}
}

func (c *config) GetAuthServerURL() string {
	return c.authServerURL
}

func (c *config) SetAuthServerURL(url string) {
	c.authServerURL = url
}

func (c *config) GetRealm() string {
	return c.realm
}

func (c *config) SetRealm(realm string) {
	c.realm = realm
}

func (c *config) GetClientID() string {
	return c.clientID
}

func (c *config) SetClientID(id string) {
	c.clientID = id
}

func (c *config) GetClientSecret() string {
	return c.clientSecret
}

func (c *config) SetClientSecret(secret string) {
	c.clientSecret = secret
}

func (c *config) Validate() error {
	if c.authServerURL == "" {
		return fmt.Errorf("auth server URL is not configured")
	}
	if c.realm == "" {
		return fmt.Errorf("realm is not configured")
	}
	if c.clientID == "" {
		return fmt.Errorf("client ID is not configured")
	}
	if c.clientSecret == "" {
		return fmt.Errorf("client secret is not configured")
	}
	return nil
}
