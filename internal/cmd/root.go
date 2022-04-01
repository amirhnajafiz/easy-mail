package cmd

import (
	"github.com/amirhnajafiz/easy-mail/internal/config"
)

func Exec() {
	cfg := config.Load()

	// Begin testing
	s := GetServer(cfg)
	_ = s.Run(":5000")
}
