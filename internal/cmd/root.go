package cmd

import (
	"github.com/amirhnajafiz/easy-mail/internal/config"
	"github.com/amirhnajafiz/easy-mail/internal/mail"
)

func Exec() {
	cfg := config.Load()

	postman := mail.Mail{
		Cfg: cfg.MailGun,
	}

	postman.Init()

	// Begin testing
	s := GetServer(postman)
	_ = s.Run(":5000")
}
