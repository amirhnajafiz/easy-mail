package server

import (
	"github.com/amirhnajafiz/easy-mail/internal/config"
	"github.com/amirhnajafiz/easy-mail/internal/http/handler"
	"github.com/amirhnajafiz/easy-mail/internal/logger"
	"github.com/amirhnajafiz/easy-mail/internal/mail"
	"github.com/gin-gonic/gin"
)

func GetServer(cfg config.Config) *gin.Engine {
	router := gin.Default()
	log := logger.New(cfg.Logger)

	postman := mail.Mail{
		Cfg:    cfg.MailGun,
		Logger: log.Named("mail"),
	}
	postman.Init()

	h := handler.Handler{
		Postman: postman,
		Logger:  log.Named("handler"),
	}

	router.POST("/mail/send", h.SendMail)

	return router
}
