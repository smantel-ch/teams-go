package main

import (
	"encoding/json"
	"log"
	"os"

	teams "github.com/smantel-ch/teams-go/AdaptiveCard"
)

func main() {
	card := teams.NewAdaptiveCard()

	tb1 := teams.NewTextBlock("TextBlock 1")
	card.Body = append(card.Body, tb1)

	json, err := json.Marshal(card)
	if err != nil {
	}
	log.Println(string(json))

	webhookUrl := os.Getenv("TEAMS_WEBHOOK")

	webhook := teams.NewWebhook(webhookUrl)
	if err := webhook.Send(&teams.Message{
		Type: "message",
		Attachments: []teams.Attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				ContentUrl:  "",
				Content:     card,
			},
		},
	}); err != nil {
		log.Println(err)
	}

}
