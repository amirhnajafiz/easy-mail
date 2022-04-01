package mail

import (
	"github.com/mailgun/mailgun-go"
	"go.uber.org/zap"
)

type Mail struct {
	Cfg    Config
	Logger *zap.Logger
	Client *mailgun.MailgunImpl
}

func (m *Mail) Init() {
	m.Client = mailgun.NewMailgun(m.Cfg.Domain, m.Cfg.APIKEY)
}

func (m *Mail) Send(e Envelope, validation bool, isHTML bool) (string, error) {
	if validation {
		if err := m.validate(e.From, e.To); err != nil {
			m.Logger.Error("failed in validation", zap.Error(err))

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
