package main

//
//func webhookHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
//		return
//	}
//
//	var pr Model.PullRequest
//	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
//		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
//		return
//	}
//
//	if pr.Action == "opened" {
//		message := fmt.Sprintf("\U0001F680 New Pull Request in *%s* by *%s*\n\nTitle: %s\n[View Pull Request](%s)", pr.Repository.Name, pr.Sender.Login, pr.PullRequest.Title, pr.PullRequest.HTMLURL)
//		sendTelegramMessage(message)
//	}
//
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte(`{"status": "received"}`))
//}
