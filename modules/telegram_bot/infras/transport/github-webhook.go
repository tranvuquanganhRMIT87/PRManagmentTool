package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	telegrambotmodel "ngrok-go-quickstart/modules/telegram_bot/model"
)

func (h *httpService) GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	eventType := r.Header.Get("X-GitHub-Event")
	fmt.Println("Received a webhook event:", eventType)

	var payload telegrambotmodel.Payload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		log.Printf("Error decoding payload (Event: %s): %v", eventType, err)
		return
	}
	defer r.Body.Close()

	ok, err := h.pushMessageHandler.Execute(r.Context(), payload)
	if err != nil {
		http.Error(w, "Failed to execute handler", http.StatusInternalServerError)
		log.Printf("Error executing push message handler (Event: %s): %v", eventType, err)
		return
	}

	if !ok {
		http.Error(w, "No action taken", http.StatusNoContent)
		log.Printf("No action taken for the webhook event (Event: %s)", eventType)
		return
	}

	w.WriteHeader(http.StatusOK)
}
