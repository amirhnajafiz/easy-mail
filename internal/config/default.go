package config

import (
	"github.com/amirhnajafiz/easy-mail/internal/mailer"
)

func Default() Config {
	return Config{
		Server: ":5000",
		MailGun: mailer.Config{
			Domain: "",
			APIKEY: "",
		},
	}
}
