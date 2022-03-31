package config

import (
	"github.com/amirhnajafiz/easy-mail/internal/mail"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"log"
)

type Config struct {
	MailGun mail.Config `koanf:"mail_gun"`
}

func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config.yaml file: %v\n", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %v\n", err)
	}

	return instance
}
