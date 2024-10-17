package ports

import (
	"context"
	"net/http"
)

type Client interface {
	GetAccessToken(ctx context.Context) (string, error)
	Do(req *http.Request) (*http.Response, error)
}

type Config interface {
	GetAuthServerURL() string
	SetAuthServerURL(string)
	GetRealm() string
	SetRealm(string)
	GetClientID() string
	SetClientID(string)
	GetClientSecret() string
	SetClientSecret(string)
	Validate() error
}
