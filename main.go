package main

import (
	"log"
	"os"

	teams "github.com/smantel-ch/teams-go/AdaptiveCard"
)

func main() {
	testMsg := teams.AdaptiveCard{
		Type:    teams.TypeAdaptiveCard,
		Version: teams.Version13,
		Body: []teams.Element{
			&teams.Container{
				Type:      teams.TypeContainer,
				Separator: true,
				Spacing:   teams.SpacingExtraLarge,
				Items: []teams.Element{
					&teams.TextBlock{
						Type: teams.TypeTextBlock,
						Text: "**alertWatcher triggered a new Alert**",
					},
					&teams.TextBlock{
						Type:    teams.TypeTextBlock,
						Text:    "fruit.bananana",
						Spacing: teams.SpacingNone,
						Size:    teams.FontSizeSmall,
					},
				},
			},
			&teams.Container{
				Type:      teams.TypeContainer,
				Separator: true,
				Spacing:   teams.SpacingExtraLarge,
				Items: []teams.Element{
					&teams.FactSet{
						Type: teams.TypeFactSet,
						Facts: []teams.Fact{
							{
								Title: "Ruleset",
								Value: "bananana",
							},
							{
								Title: "Date/Time",
								Value: "Fri, 02 Sep 2022 08:11:50 CEST",
							},
							{
								Title: "Severity",
								Value: "EXTREME",
							},
							{
								Title: "Status",
								Value: "🔥 FIRING",
							},
						},
					},
				},
			},
			&teams.Container{
				Type:      teams.TypeContainer,
				Separator: true,
				Spacing:   teams.SpacingExtraLarge,
				Items: []teams.Element{
					&teams.ActionSet{
						Type: teams.TypeActionSet,
						Actions: []teams.Action{
							&teams.ActionOpenUrl{
								Type:  teams.TypeActionOpenUrl,
								Title: "Open in 4Dmetrics",
								Url:   "https://4dmetrics.4data.ch",
							},
						},
					},
				},
			},
		},
	}

	webhookUrl := os.Getenv("TEAMS_WEBHOOK")

	webhook, err := teams.NewWebhook(webhookUrl)
	if err != nil {
		log.Println(err)
	}
	if err := webhook.Send(&testMsg); err != nil {
		log.Println(err)
	}

}
