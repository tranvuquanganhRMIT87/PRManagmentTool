package cmd

import (
	"log"
	"net/http"
)

func Execute() {
	http.HandleFunc("/webhook", githubWebhookHandler)
	log.Println("Server is listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
