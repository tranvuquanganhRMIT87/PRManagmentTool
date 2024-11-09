package openai

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
	Client *openai.Client
}

type CodeReviewAssistant interface {
	Connect(ctx context.Context, httpClient *http.Client) error
}

func NewOpenAI(token string) *openAI {
	return &openAI{
		Token: token,
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

func (o *openAI) CodeReviewAssistant(ctx context.Context, code string) (string, error) {
	return "", nil

}
