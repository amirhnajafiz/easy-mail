package config

import "github.com/amirhnajafiz/easy-mail/internal/mail"

func Default() Config {
	return Config{
		MailGun: mail.Config{
			Domain: "",
			APIKEY: "",
		},
	}
}
