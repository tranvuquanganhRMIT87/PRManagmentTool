package main

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
