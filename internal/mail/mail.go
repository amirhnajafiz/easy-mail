package mail

import "github.com/mailgun/mailgun-go"

type Mail struct {
	Cfg    Config
	Client *mailgun.MailgunImpl
}

func (m *Mail) Init() {
	m.Client = mailgun.NewMailgun(m.Cfg.Domain, m.Cfg.APIKEY)
}

func (m *Mail) SendSimpleMessage(e Envelope) (string, error) {
	ms := m.Client.NewMessage(
		e.From,
		e.Subject,
		e.Text,
		e.To,
	)
	_, id, err := m.Client.Send(ms)

	return id, err
}
