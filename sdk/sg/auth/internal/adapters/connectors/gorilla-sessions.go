package authconn

import (
	"fmt"

	"github.com/gin-gonic/gin"

	sdksession "github.com/devpablocristo/golang/sdk/pkg/sessions/gorilla"
	sdksessiondefs "github.com/devpablocristo/golang/sdk/pkg/sessions/gorilla/defs"

	ports "github.com/devpablocristo/golang/sdk/sg/auth/internal/core/ports"
)

type GorillaSessionManager struct {
	session sdksessiondefs.SessionManager
}

// NewGorillaSessionManager inicializa el manejador de sesiones Gorilla
func NewGorillaSessionManager() (ports.SessionManager, error) {
	r, err := sdksession.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}

	return &GorillaSessionManager{
		session: r,
	}, nil
}

// SaveJWTToSession stores a JWT in the user's session using Gin Context
func (gsm *GorillaSessionManager) JwtToSession(c *gin.Context, jwtToken, sessionName string) error {
	// Get the session (session name: "login")
	session, err := gsm.session.Get(c.Request, sessionName)
	if err != nil {
		return fmt.Errorf("failed to retrieve session: %w", err)
	}

	// Store the JWT in the session
	session.Values["jwt"] = jwtToken

	// Save the session to the client (cookie)
	if err = gsm.session.Save(c.Request, c.Writer, session); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}
