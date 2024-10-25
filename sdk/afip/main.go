package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config holds the configuration for AFIP authentication
type Config struct {
	AuthServerURL string
	Realm         string
	ClientID      string
	ClientSecret  string
}

// TokenResponse represents the structure of the token response from AFIP
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// AFIPClient handles communication with AFIP
type AFIPClient struct {
	httpClient  *http.Client
	config      Config
	token       string
	tokenExpiry time.Time
	mu          sync.Mutex
}

// NewAFIPClient creates a new instance of AFIPClient with the provided configuration
func NewAFIPClient(config Config) (*AFIPClient, error) {
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Configure the HTTP client with an appropriate timeout
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &AFIPClient{
		httpClient: httpClient,
		config:     config,
	}, nil
}

// validateConfig checks if all required configuration fields are present
func validateConfig(config Config) error {
	if config.AuthServerURL == "" {
		return fmt.Errorf("AuthServerURL is required")
	}
	if config.Realm == "" {
		return fmt.Errorf("Realm is required")
	}
	if config.ClientID == "" {
		return fmt.Errorf("ClientID is required")
	}
	if config.ClientSecret == "" {
		return fmt.Errorf("ClientSecret is required")
	}
	return nil
}

// GetAccessToken retrieves a valid access token, refreshing it if necessary
func (a *AFIPClient) GetAccessToken(ctx context.Context) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Reuse the token if it's still valid
	if a.token != "" && time.Now().Before(a.tokenExpiry) {
		return a.token, nil
	}

	// Build the token endpoint URL
	endpoint := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", a.config.AuthServerURL, a.config.Realm)

	// Prepare the form data
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", a.config.ClientID)
	data.Set("client_secret", a.config.ClientSecret)

	// Create the POST request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Decode the response
	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("access token is empty")
	}

	// Store the token and its expiry time with a buffer
	expiryBuffer := 60 // seconds
	a.token = tokenResp.AccessToken
	a.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn-expiryBuffer) * time.Second)

	return a.token, nil
}

// ValidateCUIT validates a CUIL with AFIP using the access token
func (a *AFIPClient) ValidateCUIT(ctx context.Context, cuil string) error {
	if !isValidCUIT(cuil) {
		return fmt.Errorf("invalid CUIL format")
	}

	// Obtain the access token
	token, err := a.GetAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	// Build the validation endpoint URL
	endpoint := fmt.Sprintf("https://tst.autenticar.gob.ar/api/validate-cuil?cuil=%s", cuil)

	// Create the GET request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to create validation request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("validation request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("validation failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Optionally, process the response body
	/*
		var validationResp ValidationResponse
		if err := json.NewDecoder(resp.Body).Decode(&validationResp); err != nil {
			return fmt.Errorf("failed to decode validation response: %w", err)
		}
		// Further processing based on validationResp
	*/

	return nil
}

// isValidCUIT verifies if the CUIL has the correct format and checksum
func isValidCUIT(cuil string) bool {
	if len(cuil) != 11 {
		return false
	}
	for _, char := range cuil {
		if char < '0' || char > '9' {
			return false
		}
	}

	multipliers := []int{5, 4, 3, 2, 7, 6, 5, 4, 3, 2}
	sum := 0
	for i := 0; i < 10; i++ {
		digit := int(cuil[i] - '0')
		sum += digit * multipliers[i]
	}
	remainder := sum % 11
	checkDigit := 11 - remainder
	if checkDigit == 11 {
		checkDigit = 0
	} else if checkDigit == 10 {
		checkDigit = 9
	}
	return int(cuil[10]-'0') == checkDigit
}

// loadConfig reads configuration from environment variables or config files
func loadConfig() (Config, error) {
	viper.AutomaticEnv()

	// Load variables from .env file if it exists
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found; continue with environment variables
	}

	// Set default values (adjust as needed)
	viper.SetDefault("AUTH_SERVER_URL", "https://tst.autenticar.gob.ar/auth")
	viper.SetDefault("REALM", "tesi-afip")
	viper.SetDefault("CLIENT_ID", "tesi")
	viper.SetDefault("CLIENT_SECRET", "ce5abdb2-9b00-431c-a213-8c815cb97226") // Ensure this is handled securely

	config := Config{
		AuthServerURL: viper.GetString("AUTH_SERVER_URL"),
		Realm:         viper.GetString("REALM"),
		ClientID:      viper.GetString("CLIENT_ID"),
		ClientSecret:  viper.GetString("CLIENT_SECRET"),
	}

	return config, nil
}

func main() {
	// Load the .env file if it exists
	if err := godotenv.Load("env.local"); err != nil {
		log.Println("No .env file found")
	}

	// Load the configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create an instance of AFIPClient
	afipClient, err := NewAFIPClient(config)
	if err != nil {
		log.Fatalf("Error creating AFIP client: %v", err)
	}

	// Set up the Gin router
	router := gin.Default()

	// Define the validation route
	router.POST("/validate", func(c *gin.Context) {
		// Extract the CUIL from the form
		cuil := c.PostForm("cuil")
		if cuil == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "CUIL is required"})
			return
		}

		// Create a context with timeout for the operation
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Validate the CUIL with AFIP
		err := afipClient.ValidateCUIT(ctx, cuil)
		if err != nil {
			log.Printf("Validation error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Validation failed"})
			return
		}

		// Respond successfully
		c.JSON(http.StatusOK, gin.H{"message": "CUIL validated successfully"})
	})

	// Start the server on port 8080
	log.Println("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
