package apis

import (
	"encoding/json"
	"net/http"
)

// Define an interface for sending emails
type EmailSender interface {
	Send(recipient, subject, body string) error
}

// Define your API structure
type EmailAPI struct {
	Sender EmailSender
}

func (api *EmailAPI) SendEmailApi(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RecipientEmail string `json:"RecipientEmail"`
		Subject        string `json:"Subject"`
		Body           string `json:"Body"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Use the interface to send the email
	err := api.Sender.Send(req.RecipientEmail, req.Subject, req.Body)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}