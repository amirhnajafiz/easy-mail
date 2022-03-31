package cmd

import (
	"github.com/amirhnajafiz/easy-mail/internal/http/handler"
	"github.com/amirhnajafiz/easy-mail/internal/mail"
	"github.com/gin-gonic/gin"
)

func GetServer(p mail.Mail) *gin.Engine {
	router := gin.Default()

	h := handler.Handler{
		Postman: p,
	}

	router.POST("/send/main", h.SendMail)

	return router
}
