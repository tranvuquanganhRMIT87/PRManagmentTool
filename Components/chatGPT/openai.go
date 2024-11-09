package chatGPT

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

type openAI struct {
	Token  string
	Model  string
	Client *openai.Client
}

type CodeReviewAssistant interface {
	Connect(ctx context.Context, httpClient *http.Client) error
	GetOpenAiSuggestions(ctx context.Context, patch string) (string, error)
}

func NewOpenAI(token string, model string) *openAI {
	return &openAI{
		Token: token,
		Model: model,
	}
}
func (o *openAI) Connect(ctx context.Context, httpClient *http.Client) error {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}

	if o.Token == "" {
		return fmt.Errorf("OpenAI API token is empty")
	}

	client := openai.NewClient(o.Token)
	o.Client = client

	_, err := o.Client.ListModels(ctx)
	if err != nil {
		log.Printf("Error connecting to OpenAI API: %v", err)
		return fmt.Errorf("failed to connect to OpenAI API: %v", err)
	}

	logrus.Info("Connected to GitHub API")

	return nil
}

func (o *openAI) GetOpenAiSuggestions(ctx context.Context, patch string) (string, error) {
	if patch == "" {
		return "", fmt.Errorf("missing patch for OpenAI suggestion")
	}

	if o.Client == nil {
		return "", fmt.Errorf("OpenAI client is not initialized")
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "Analyze the following code patch and provide suggestions for improvement.",
		},
		{
			Role:    "user",
			Content: patch,
		},
	}

	req := openai.ChatCompletionRequest{
		Model:    o.Model,
		Messages: messages,
	}

	// Send the request to OpenAI
	resp, err := o.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		logrus.Errorf("Error getting suggestion from OpenAI: %v", err)
		return "", fmt.Errorf("failed to get suggestion from OpenAI: %v", err)
	}

	// Check if there's a response
	if len(resp.Choices) > 0 && resp.Choices[0].Message.Content != "" {
		return resp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no suggestion received from OpenAI")

}
