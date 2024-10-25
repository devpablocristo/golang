package authconn

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	sdkhclnt "github.com/devpablocristo/golang/sdk/pkg/rest/net-http/client"
	sdkhclntports "github.com/devpablocristo/golang/sdk/pkg/rest/net-http/client/ports"

	ports "github.com/devpablocristo/golang/sdk/sg/auth/internal/core/ports"
)

type HttpClient struct {
	httpClient  sdkhclntports.Client
	config      sdkhclntports.Config
	token       string
	tokenExpiry time.Time
	mu          sync.Mutex
}

// NewHttpClient crea una nueva instancia de HttpClient con la configuración proporcionada.
func NewHttpClient() (ports.HttpClient, error) {
	r, c, err := sdkhclnt.Bootstrap("AFIP_TOKEN_ENDPOINT", "x", "AFIP_CLIENT_SECRET", "y")
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}

	return &HttpClient{
		httpClient: r,
		config:     c,
	}, nil
}

func (c *HttpClient) GetAccessToken(ctx context.Context) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.token != "" && time.Now().Before(c.tokenExpiry) {
		return c.token, nil
	}

	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", c.config.GetClientID())
	params.Set("client_secret", c.config.GetClientSecret())
	params.Set("scope", c.config.GetAdditionalParams().Get("scope"))

	tokenResponse, err := c.httpClient.GetAccessToken(ctx, c.config.GetTokenEndpoint(), params)
	if err != nil {
		return "", fmt.Errorf("error al obtener el token de acceso: %w", err)
	}

	c.token = tokenResponse.GetAccessToken()

	// Asumiendo que el token expira en 1 hora (ajusta según tu caso)
	c.tokenExpiry = time.Now().Add(55 * time.Minute)

	return c.token, nil
}

// ValidateCUIT valida un CUIL con la AFIP utilizando el token de acceso.
func (c *HttpClient) ValidateCUIT(ctx context.Context, cuil string) error {
	if !isValidCUIT(cuil) {
		return fmt.Errorf("formato de CUIL inválido")
	}

	token, err := c.GetAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("error al obtener el token de acceso: %w", err)
	}

	// Aquí iría la lógica para validar el CUIL usando el token
	// Por ejemplo:
	endpoint := fmt.Sprintf("%s/api/validate-cuil?cuil=%s", c.config.GetTokenEndpoint(), cuil)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error al crear la solicitud: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error en la solicitud: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("la validación falló: código de estado %d", resp.StatusCode)
	}

	return nil
}

// isValidCUIT verifica si el CUIL tiene el formato correcto.
func isValidCUIT(cuil string) bool {
	if len(cuil) != 11 {
		return false
	}
	var (
		multipliers = []int{5, 4, 3, 2, 7, 6, 5, 4, 3, 2}
		sum         int
	)
	for i := 0; i < 10; i++ {
		digit := int(cuil[i] - '0')
		sum += digit * multipliers[i]
	}
	remainder := sum % 11
	checkDigit := 11 - remainder
	switch checkDigit {
	case 11:
		checkDigit = 0
	case 10:
		checkDigit = 9
	}
	return int(cuil[10]-'0') == checkDigit
}
