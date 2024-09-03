package main

import "net/http"

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPaylod mailMessage

	err := app.readJson(w, r, &requestPaylod)
	if err != nil {
		_ = app.errorJson(w, err)
		return
	}

	msg := Message{
		From:    requestPaylod.From,
		To:      requestPaylod.To,
		Subject: requestPaylod.Subject,
		Data:    requestPaylod.Message,
	}

	if err = app.Mailer.SendSMTPMessage(msg); err != nil {
		_ = app.errorJson(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "sent to " + requestPaylod.To,
	}

	_ = app.writeJson(w, http.StatusAccepted, payload)
}
