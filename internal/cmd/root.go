package cmd

import (
	"github.com/amirhnajafiz/easy-mail/internal/config"
	"github.com/amirhnajafiz/easy-mail/internal/http/handler"
	"github.com/amirhnajafiz/easy-mail/internal/mail"
	"github.com/gin-gonic/gin"
)

func Execute() {
	cfg := config.Load()

	router := gin.Default()

	postman := mail.Mail{
		Cfg: cfg.MailGun,
	}
	postman.Init()

	h := handler.Handler{
		Postman: postman,
	}

	router.POST("/mail/send", h.SendMail)

	_ = router.Run(cfg.Server)
}
