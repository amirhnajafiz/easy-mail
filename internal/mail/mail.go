package mail

import "github.com/mailgun/mailgun-go"

func SendSimpleMessage(domain string, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		"Excited User <elliot@tutorialedge.net>",
		"Hello",
		"Testing some Mailgun!",
		"elliot@tutorialedge.net",
	)
	_, id, err := mg.Send(m)
	return id, err
}
