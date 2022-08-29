package config

import (
	"github.com/amirhnajafiz/easy-mail/internal/mail"
)

func Default() Config {
	return Config{
		Server: ":5000",
		MailGun: mail.Config{
			Domain: "",
			APIKEY: "",
		},
	}
}
