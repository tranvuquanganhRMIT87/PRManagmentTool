package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ngrok-go-quickstart/Model"
	"ngrok-go-quickstart/Shared"
	"ngrok-go-quickstart/internal/server"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var pr Model.PullRequest
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if pr.Action == "opened" {
		message := fmt.Sprintf("\U0001F680 New Pull Request in *%s* by *%s*\n\nTitle: %s\n[View Pull Request](%s)", pr.Repository.Name, pr.Sender.Login, pr.PullRequest.Title, pr.PullRequest.HTMLURL)
		sendTelegramMessage(message)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "received"}`))
}

func sendTelegramMessage(message string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", Shared.TelegramBotToken)
	payload := map[string]string{
		"chat_id":    Shared.ChatID,
		"text":       message,
		"parse_mode": "Markdown",
	}
	payloadBytes, _ := json.Marshal(payload)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func main() {

	// new server
	if err := server.Run(context.Background()); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", webhookHandler)
	log.Println("Server is listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
