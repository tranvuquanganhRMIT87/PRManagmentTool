package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ngrok-go-quickstart/Shared"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodPost {
	//	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	//	return
	//}
	//
	//var pr Model.PullRequest
	//if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
	//	http.Error(w, "Error parsing JSON", http.StatusBadRequest)
	//	return
	//}
	//
	//if pr.Action == "opened" {
	//	message := fmt.Sprintf("\U0001F680 New Pull Request in *%s* by *%s*\n\nTitle: %s\n[View Pull Request](%s)", pr.Repository.Name, pr.Sender.Login, pr.PullRequest.Title, pr.PullRequest.HTMLURL)
	//	sendTelegramMessage(message)
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte(`{"status": "received"}`))

	fmt.Printf("Received Payload: %+v\n", r.Body)
	fmt.Println("Received a webhook event", r.Method)
	fmt.Println("Received a ResponseWriter", w)
	var payload struct {
		Action      string `json:"action"`
		PullRequest struct {
			URL   string `json:"html_url"`
			Title string `json:"title"`
			User  struct {
				Login string `json:"login"`
			} `json:"user"`
		} `json:"pull_request"`
		Repository struct {
			FullName string `json:"full_name"`
			HTMLURL  string `json:"html_url"`
		} `json:"repository"`
	}

	// Decode the JSON payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Only handle "opened" pull requests
	if payload.Action == "opened" {
		message := fmt.Sprintf(
			"New Pull Request in repository %s:\nTitle: %s\nBy: %s\nPR URL: %s\nRepository URL: %s",
			payload.Repository.FullName,
			payload.PullRequest.Title,
			payload.PullRequest.User.Login,
			payload.PullRequest.URL,
			payload.Repository.HTMLURL,
		)

		// Send the message to Telegram
		if err := sendTelegramMessage(message); err != nil {
			fmt.Println("Error sending message:", err)
		}
	}

	var rawPayload map[string]interface{}
	json.NewDecoder(r.Body).Decode(&rawPayload)
	fmt.Printf("Received Payload: %+v\n", rawPayload)

	w.WriteHeader(http.StatusOK)
}

//func sendTelegramMessage(message string) {
//	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", Shared.TelegramBotToken)
//	payload := map[string]string{
//		"chat_id":    Shared.ChatID,
//		"text":       message,
//		"parse_mode": "Markdown",
//	}
//	payloadBytes, _ := json.Marshal(payload)
//	_, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
//	if err != nil {
//		log.Printf("Error sending message: %v", err)
//	}
//}

// Function to send a Telegram message
func sendTelegramMessage(message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", Shared.TelegramBotToken)
	payload := map[string]string{
		"chat_id": Shared.ChatID,
		"text":    message,
	}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func main() {

	//// new server
	//if err := server.Run(context.Background()); err != nil {
	//	log.Fatal(err)
	//}

	http.HandleFunc("/webhook", webhookHandler)
	log.Println("Server is listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
