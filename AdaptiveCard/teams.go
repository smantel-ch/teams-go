package teams

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	Webhook string
}

type Webhook struct {
	Url string
}

func NewWebhook(url string) *Webhook {
	return &Webhook{
		Url: url,
	}
}

func (w *Webhook) Send(msg *Message) error {
	json, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, w.Url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	} else {
		log.Printf("%+v", res)
	}

	return nil

}
