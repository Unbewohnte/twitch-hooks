package vkhooks

import (
	"fmt"

	vk "github.com/go-vk-api/vk"
)

type ApiKey struct {
	Key string
}

var client *vk.Client

type Message struct {
	Message   string `json:"message"`
	GroupChat bool   `json:"is_group_chat"`
	ID        uint   `json:"receiver_id"`
}

func Initialise(vkapikey string) {
	// create a client on init
	vkClient, err := vk.NewClientWithOptions(
		vk.WithToken(vkapikey),
	)
	if err != nil {
		panic(err)
	}

	client = vkClient
}

// Sends message to the given id
func Send(message Message) error {
	switch message.GroupChat {
	case true:
		err := client.CallMethod("messages.send", vk.RequestParams{
			"chat_id":   message.ID,
			"message":   message.Message,
			"random_id": 0,
		}, nil)
		if err != nil {
			return fmt.Errorf("could not send vk message: %s", err)
		}

	case false:
		err := client.CallMethod("messages.send", vk.RequestParams{
			"peer_id":   message.ID,
			"message":   message.Message,
			"random_id": 0,
		}, nil)
		if err != nil {
			return fmt.Errorf("could not send vk message: %s", err)
		}
	}

	return nil
}
