package teams

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//
type Webhook struct {
	// Webhook URL where the message is sent to
	Url string
	// http client used to send the request including whatever configuration required
	client *http.Client
}

func NewWebhook(url string) (*Webhook, error) {
	webhook := &Webhook{}

	if ok := isValidUri(url); ok != true {
		return nil, errors.New("url is not valid")
	} else {
		webhook.Url = url
	}

	client := &http.Client{}

	webhook.client = client

	return webhook, nil
}

func (w *Webhook) Send(card Card) error {
	var msg interface{}
	switch card.(type) {
	case *AdaptiveCard:
		msg = &Message{
			Type: "message",
			Attachments: []Attachment{
				{
					ContentType: "application/vnd.microsoft.card.adaptive",
					ContentUrl:  "",
					Content:     card,
				},
			},
		}
	default:
		msg = card
	}

	json, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, w.Url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := w.client.Do(req)
	if err != nil {
		return err
	} else {
		if res.StatusCode != http.StatusOK {
			return errors.New(fmt.Sprintf("Expected StatusCode:%d, but got %d", http.StatusOK, res.StatusCode))
		}
	}

	return nil

}
