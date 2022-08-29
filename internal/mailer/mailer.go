package mailer

import (
	"github.com/mailgun/mailgun-go"
)

type Mailer struct {
	APIKEY string
	Client *mailgun.MailgunImpl
}

type Mail struct {
	From    string
	Subject string
	Text    string
	To      string
}

func New(cfg Config) Mailer {
	return Mailer{
		APIKEY: cfg.APIKEY,
		Client: mailgun.NewMailgun(cfg.Domain, cfg.APIKEY),
	}
}

func (m *Mailer) Send(e Mail, validation bool, isHTML bool) (string, error) {
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

	if isHTML {
		ms.SetHtml(e.Text)
	}

	_, id, err := m.Client.Send(ms)

	return id, err
}

func (m *Mailer) validate(sender string, receiver string) error {
	v := mailgun.NewEmailValidator(m.APIKEY)

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
