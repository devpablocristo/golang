package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Message represents a message in the conversation.
type Message struct {
	Role    string `json:"role"`    // Can be "user" or "assistant".
	Content string `json:"content"` // The content of the message.
}

// OpenAIRequest structure for JSON request.
type OpenAIRequest struct {
	MaxTokens int       `json:"max_tokens"`
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
}

// Choice represents a choice in the OpenAI response.
type Choice struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

// OpenAIResponse represents the response from the OpenAI API.
type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

// GenerateJSONFromInput generates a JSON object from user input.
func GenerateJSONFromInput(userInput string) (string, error) {
	// The URL of the OpenAI API.
	url := "https://api.openai.com/v1/chat/completions"

	// Your OpenAI API key.
	apiKey := "" // Replace with your actual API key.

	// Create an instance of OpenAIRequest.
	requestBody := OpenAIRequest{
		MaxTokens: 300,
		Model:     "gpt-4", // or "text-davinci-004" for GPT-4.
		Messages:  []Message{{Role: "user", Content: userInput}},
	}

	// Convert the structure to JSON.
	jsonReq, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	// Create a new HTTP POST request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return "", err
	}

	// Add necessary headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Create an HTTP client and send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and return the response.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// ParseJSONResponse parses the JSON response and returns the formatted JSON string.
func ParseJSONResponse(responseJSON string) (string, error) {
	var response OpenAIResponse
	err := json.Unmarshal([]byte(responseJSON), &response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		content := response.Choices[0].Message.Content

		// Remove the "Content: " prefix and trim spaces
		jsonString := strings.TrimSpace(strings.TrimPrefix(content, "Content:"))

		// Deserialize the JSON into a Go object (e.g., into a map for flexibility)
		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonString), &result)
		if err != nil {
			return "", err
		}

		// Serialize and return the Go object as JSON
		formattedJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return "", err
		}
		return string(formattedJSON), nil
	}
	return "", fmt.Errorf("No content found in the response.")
}

func main() {
	promptContext := "Please generate a JSON object with the following structure: " +
		"{\n" +
		"  \"squad_id\": [integer],\n" +
		"  \"user_id\": [integer],\n" +
		"  \"squad_log_type_id\": [integer],\n" +
		"  \"skill_categories\": [array of strings],\n" +
		"  \"description\": [string],\n" +
		"  \"start_date\": [date in YYYY-MM-DD format],\n" +
		"  \"end_date\": [date in YYYY-MM-DD format]\n" +
		"}\n" +
		"Fill in the fields with appropriate sample values."

	userInput := promptContext + "\n create an accessible with description 'my description' and categories 'cat a, cat b, cat c'"

	responseJSON, err := GenerateJSONFromInput(userInput)
	if err != nil {
		fmt.Println("Error generating JSON from input:", err)
		return
	}

	formattedJSON, err := ParseJSONResponse(responseJSON)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	fmt.Println(formattedJSON)
}
