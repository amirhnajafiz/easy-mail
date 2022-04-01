package cmd

import (
	"github.com/amirhnajafiz/easy-mail/internal/cmd/server"
	"github.com/amirhnajafiz/easy-mail/internal/config"
)

func Exec() {
	cfg := config.Load()

	// Begin testing
	s := server.GetServer(cfg)
	_ = s.Run(cfg.Server.Address)
}
