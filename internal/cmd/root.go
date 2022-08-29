package cmd

import (
	"github.com/amirhnajafiz/easy-mail/internal/config"
	"github.com/amirhnajafiz/easy-mail/internal/http/handler"
	"github.com/amirhnajafiz/easy-mail/internal/mailer"
	"github.com/gin-gonic/gin"
)

func Execute() {
	cfg := config.Load()

	router := gin.Default()

	m := mailer.New(cfg.MailGun)

	h := handler.Handler{
		Mailer: m,
	}

	router.POST("/mail/send", h.SendMail)

	_ = router.Run(cfg.Server)
}
