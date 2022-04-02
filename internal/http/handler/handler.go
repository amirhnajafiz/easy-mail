package handler

import (
	"net/http"

	"github.com/amirhnajafiz/easy-mail/internal/mail"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	Postman mail.Mail
	Logger  *zap.Logger
}

func (h Handler) SendMail(c *gin.Context) {
	validation := true
	isHTML := false
	from, _ := c.Get("from")
	subject, _ := c.Get("subject")
	text, _ := c.Get("text")
	to, _ := c.Get("to")

	env := mail.Envelope{
		From:    from.(string),
		Subject: subject.(string),
		Text:    text.(string),
		To:      to.(string),
	}

	if key, ok := c.Get("validation"); ok {
		if key == "no" {
			validation = false
		}
	}

	if key, ok := c.Get("is_html"); ok {
		if key == "yes" {
			isHTML = true
		}
	}

	id, err := h.Postman.Send(env, validation, isHTML)
	if err != nil {
		h.Logger.Fatal("failed to send message", zap.Error(err))

		_ = c.Error(err)
	}

	c.String(http.StatusOK, id)
}
