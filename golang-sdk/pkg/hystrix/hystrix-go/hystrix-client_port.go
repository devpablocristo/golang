package hystrixgo

import (
	"net/http"
)

type HystrixClientPort interface {
	Get(url string) (*http.Response, error)
}
