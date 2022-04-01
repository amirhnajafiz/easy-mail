package config

import (
	"github.com/amirhnajafiz/easy-mail/internal/cmd/server"
	"github.com/amirhnajafiz/easy-mail/internal/mail"
)

func Default() Config {
	return Config{
		Server: server.Config{
			Address: ":5000",
			Route:   "/send/mail",
		},
		MailGun: mail.Config{
			Domain: "",
			APIKEY: "",
		},
	}
}
