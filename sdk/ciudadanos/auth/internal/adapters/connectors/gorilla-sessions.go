package authconn

import (
	"fmt"

	sdksession "github.com/devpablocristo/golang/sdk/pkg/sessions/gorilla"
	sdksessionports "github.com/devpablocristo/golang/sdk/pkg/sessions/gorilla/ports"

	ports "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core/ports"
)

type GorillaSessionManager struct {
	session sdksessionports.SessionManager
}

func NewGorillaSessionManager() (ports.SessionManager, error) {
	r, err := sdksession.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}

	return &GorillaSessionManager{
		session: r,
	}, nil
}


