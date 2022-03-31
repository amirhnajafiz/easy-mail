package config

import "github.com/amirhnajafiz/easy-mail/internal/mail"

type Config struct {
	MailGun mail.Config `koanf:"mail_gun"`
}

func Load() Config {
	return Default()
}
