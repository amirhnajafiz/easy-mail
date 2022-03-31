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

func (m *Mail) Send(e Envelope, validation bool) (string, error) {
	if validation {
		if err := m.validate(e.From, e.To); err != nil {
			return "", err
		}
	}

	ms := m.Client.NewMessage(
		e.From,
		e.Subject,
		e.Text,
		e.To,
	)
	_, id, err := m.Client.Send(ms)

	return id, err
}

func (m *Mail) validate(sender string, receiver string) error {
	v := mailgun.NewEmailValidator(m.Cfg.APIKEY)

	_, err := v.ValidateEmail(sender, false)
	if err != nil {
		return err
	}

	_, err = v.ValidateEmail(receiver, false)
	if err != nil {
		return err
	}

	return nil
}
