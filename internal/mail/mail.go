package mail

import "github.com/mailgun/mailgun-go"

type Mail struct {
	Cfg    Config
	Client *mailgun.MailgunImpl
}

func (m *Mail) Init() {
	m.Client = mailgun.NewMailgun(m.Cfg.Domain, m.Cfg.APIKEY)
}

func (m *Mail) SendSimpleMessage(p Packet) (string, error) {
	ms := m.Client.NewMessage(
		p.FFrom,
		p.SSubject,
		p.TText,
		p.TTo,
	)
	_, id, err := m.Client.Send(ms)

	return id, err
}
