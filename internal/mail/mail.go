package mail

import "github.com/mailgun/mailgun-go"

type Mail struct {
	Cfg Config
}

func (m Mail) SendSimpleMessage() (string, error) {
	mg := mailgun.NewMailgun(m.Cfg.Domain, m.Cfg.APIKEY)
	ms := mg.NewMessage(
		"Excited User <elliot@tutorialedge.net>",
		"Hello",
		"Testing some Mailgun!",
		"elliot@tutorialedge.net",
	)
	_, id, err := mg.Send(ms)

	return id, err
}
