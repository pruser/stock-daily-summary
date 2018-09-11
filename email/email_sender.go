package email

type MessageSettings struct {
	subject   string
	sender    string
	recipient string
}

func NewMessageSettings(subject, sender, recipient string) *MessageSettings {
	return &MessageSettings{subject: subject, sender: sender, recipient: recipient}
}

type Sender interface {
	Send(settings MessageSettings, content string) error
}
