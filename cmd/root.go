package cmd

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"ngrok-go-quickstart/Components/chatGPT"
	githubs "ngrok-go-quickstart/Components/github"
	components "ngrok-go-quickstart/Components/logrus"
	internal "ngrok-go-quickstart/Components/telegram"
	"ngrok-go-quickstart/middleware"
	"ngrok-go-quickstart/modules/telegram_bot"
	"ngrok-go-quickstart/share"
)

func Execute() {
	components.InitLogger()
	// Initialize config
	config := share.NewEnvConfig()
	config.InitConfig()

	// Initialize Telegram bot
	token := config.GetTelegramToken()

	bot := internal.NewTelegram(token)
	if err := bot.Connect(); err != nil {
		log.Fatalf("Failed to connect to Telegram: %v", err)
	}

	// Connect to GitHub
	ctx := context.Background()
	githubAPI := githubs.NewGithub(config.GetGithubToken())
	err := githubAPI.Connect(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to connect to GitHub: %v", err)
	}

	openAI := chatGPT.NewOpenAI(config.GetOpenAIToken(), config.GetModels())
	err = openAI.Connect(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to connect to OpenAI: %v", err)
	}

	// Initialize server mux
	mux := http.NewServeMux()

	// Initialize service context
	serviceContext := share.NewServiceContext(bot, config, githubAPI, openAI)

	// Set up routes
	telegram_bot.SetupTelegramBotService(serviceContext, mux)

	// Add ping route
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			panic(err)
		}
	})

	// Recover from panic
	logrus.Info("Starting server on port ", config.GetHttpAddr())
	logrus.Error(http.ListenAndServe(fmt.Sprintf(":%v", config.GetHttpAddr()), middleware.RecoverMiddleware(mux)))
}
