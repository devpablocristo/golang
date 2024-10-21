package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

type Step struct {
	Explanation string `json:"explanation"`
	Output      string `json:"output"`
}

type MathResponse struct {
	Steps       []Step `json:"steps"`
	FinalAnswer string `json:"final_answer"`
}

var mathResponseSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"steps": map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"explanation": map[string]interface{}{
						"type": "string",
					},
					"output": map[string]interface{}{
						"type": "string",
					},
				},
				"required":             []string{"explanation", "output"},
				"additionalProperties": false,
			},
		},
		"final_answer": map[string]interface{}{
			"type": "string",
		},
	},
	"required":             []string{"steps", "final_answer"},
	"additionalProperties": false,
}

type JSONSchemaMarshaler struct {
	Schema map[string]interface{}
}

func (j JSONSchemaMarshaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Schema)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY no está configurada")
		return
	}

	client := openai.NewClient(apiKey)

	schema := JSONSchemaMarshaler{Schema: mathResponseSchema}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o20240806,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful math tutor. Guide the user through the solution step by step.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "how can I solve 8x + 7 = -23",
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:   "MathResponse",
					Schema: schema,
					Strict: true,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(string(resp.Choices[0].Message.Content))
}
