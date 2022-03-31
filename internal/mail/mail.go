package mail

import (
	"github.com/mailgun/mailgun-go"
)

type Mail struct {
	Cfg    Config
	Client *mailgun.MailgunImpl
}

func (m *Mail) Init() {
	m.Client = mailgun.NewMailgun(m.Cfg.Domain, m.Cfg.APIKEY)
}

func (m *Mail) Send(e Envelope) (string, error) {
	v := mailgun.NewEmailValidator(m.Cfg.APIKEY)

	senderEmail, err := v.ValidateEmail(e.From, false)
	if err != nil {
		return "", err
	}

	getterEmail, err := v.ValidateEmail(e.To, false)
	if err != nil {
		return "", err
	}

	ms := m.Client.NewMessage(
		senderEmail.Address,
		e.Subject,
		e.Text,
		getterEmail.Address,
	)
	_, id, err := m.Client.Send(ms)

	return id, err
}
