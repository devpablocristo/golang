package authconn

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	sdkhclnt "github.com/devpablocristo/golang/sdk/pkg/rest/net-http/client"
	sdkhclntports "github.com/devpablocristo/golang/sdk/pkg/rest/net-http/client/ports"

	ports "github.com/devpablocristo/golang/sdk/ciudadanos/auth/internal/core/ports"
)

type HttpClient struct {
	httpClient  sdkhclntports.Client
	config      sdkhclntports.Config
	token       string
	tokenExpiry time.Time
	mu          sync.Mutex
}

// NewHttpClient crea una nueva instancia de HttpClient con la configuración proporcionada.
func NewHttpClient() (ports.HttpClient, error) { //config AuthConfig) (ports.HttpClient, error) {
	r, c, err := sdkhclnt.Bootstrap()
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}

	return &HttpClient{
		httpClient: r,
		config:     c,
	}, nil
}

// GetAccessToken obtiene un token de acceso válido desde el servidor de autenticación.
func (c *HttpClient) GetAccessToken(ctx context.Context) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.token != "" && time.Now().Before(c.tokenExpiry) {
		return c.token, nil
	}

	endpoint := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", c.config.GetAuthServerURL(), c.config.GetRealm())

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.config.GetClientID())
	data.Set("client_secret", c.config.GetClientSecret())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("non-200 response (%d): %s", resp.StatusCode, string(body))
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if tokenResponse.AccessToken == "" {
		return "", fmt.Errorf("access token is empty")
	}

	// Almacenar el token y su expiración
	expiryBuffer := 60 // segundos para restar y asegurar que el token no expire durante una solicitud
	c.token = tokenResponse.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(tokenResponse.ExpiresIn-expiryBuffer) * time.Second)

	return c.token, nil
}

// ValidateCUIT valida un CUIT con la AFIP utilizando el token de acceso.
func (c *HttpClient) ValidateCUIT(ctx context.Context, cuit string) error {
	if !isValidCUIT(cuit) {
		return fmt.Errorf("invalid CUIT format")
	}

	token, err := c.GetAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	endpoint := fmt.Sprintf("%s/api/validate-cuit?cuit=%s", c.config.GetAuthServerURL(), cuit)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("validation failed (%d): %s", resp.StatusCode, string(body))
	}

	// Procesar la respuesta si es necesario

	return nil
}

// isValidCUIT verifica si el CUIT tiene el formato correcto.
func isValidCUIT(cuit string) bool {
	if len(cuit) != 11 {
		return false
	}
	var (
		multipliers = []int{5, 4, 3, 2, 7, 6, 5, 4, 3, 2}
		sum         int
	)
	for i := 0; i < 10; i++ {
		digit := int(cuit[i] - '0')
		sum += digit * multipliers[i]
	}
	remainder := sum % 11
	checkDigit := 11 - remainder
	if checkDigit == 11 {
		checkDigit = 0
	} else if checkDigit == 10 {
		checkDigit = 9
	}
	return int(cuit[10]-'0') == checkDigit
}
