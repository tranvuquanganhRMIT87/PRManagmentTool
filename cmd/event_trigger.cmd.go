package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/pkg/ioutils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"ngrok-go-quickstart/Model"
	internal "ngrok-go-quickstart/internal/telegram"
	"os"
	"strconv"
)

func githubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	eventType := r.Header.Get("X-GitHub-Event")
	fmt.Println("Received a webhook event", eventType)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//var rs string

	if err := json.NewDecoder(r.Body).Decode(&Model.Payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	//if err := json.NewDecoder(r.Body).Decode(&rs); err != nil {
	//	http.Error(w, "Invalid payload", http.StatusBadRequest)
	//	return
	//}
	//
	//fmt.Println("rs", rs) //var message string
	if Model.Payload.Action == "opened" {
		message := fmt.Sprintf(
			"🆕: New Pull Request in repository %s:\nTitle: %s\nBy: %s\nPR URL: %s\nRepository URL: %s",
			Model.Payload.Repository.FullName,
			Model.Payload.PullRequest.Title,
			Model.Payload.PullRequest.User.Login,
			Model.Payload.PullRequest.URL,
			Model.Payload.Repository.HTMLURL,
		)

		fmt.Println("Sending message:", message)

		bot := internal.NewTelegram(os.Getenv("TELE_BOT_TOKEN"))
		fmt.Println("Connected to telegram", bot)
		fmt.Println("Key", os.Getenv("TELE_BOT_TOKEN"))
		if err := bot.Connect(); err != nil {
			log.Fatal(err)
		}

		threadID, err := strconv.Atoi(os.Getenv("THREAD_ID"))
		if err != nil {
			log.Fatal(err)
		}

		// After
		payloadBytes, err := json.Marshal(&Model.Payload)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutils.AtomicWriteFile("thread_id.txt", payloadBytes, 0644)
		if err != nil {
			return
		}

		chatId, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

		formatedMessage := internal.NewBotMessage(chatId, &threadID, message)

		// Send the message to Telegram
		if err := bot.SendMessage(formatedMessage); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Sent message to telegram")
	}

}
