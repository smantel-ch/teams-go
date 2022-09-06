package teams

const (
	TypeMessage Type = "Message"
)

type Card interface {
	IsCard() bool
}

func (a *AdaptiveCard) IsCard() bool {
	return true
}

type Message struct {
	Type        Type         `json:"type"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	ContentType string `json:"contentType"`
	ContentUrl  string `json:"contentUrl"`
	Content     Card   `json:"content"`
}
