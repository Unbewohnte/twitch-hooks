package discordhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Webhook struct {
	WebhookUrl string
}

const contentTypeJson string = "application/json"

type Message struct {
	Message   string `json:"content"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// Post Message struct to given webhook
func Post(webhookUrl string, message Message) error {
	json, err := json.Marshal(&message)
	if err != nil {
		return fmt.Errorf("could not marshal given JsonMessage: %s", err)
	}

	resp, err := http.Post(webhookUrl, contentTypeJson, bytes.NewBuffer(json))
	if err != nil {
		return fmt.Errorf("could not POST to given url: %s", err)
	}
	defer resp.Body.Close()

	return nil
}
