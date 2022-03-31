package mail

import "github.com/mailgun/mailgun-go"

type Mail struct {
	Cfg    Config
	Client *mailgun.MailgunImpl
}

func (m *Mail) Init() {
	m.Client = mailgun.NewMailgun(m.Cfg.Domain, m.Cfg.APIKEY)
}

func (m *Mail) SendSimpleMessage() (string, error) {
	ms := m.Client.NewMessage(
		"Excited User <elliot@tutorialedge.net>",
		"Hello",
		"Testing some Mailgun!",
		"elliot@tutorialedge.net",
	)
	_, id, err := m.Client.Send(ms)

	return id, err
}
