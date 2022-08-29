package handler

import (
	"net/http"

	"github.com/amirhnajafiz/easy-mail/internal/mail"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Mailer mail.Mail
}

type request struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	To      string `json:"to"`

	Validation string `json:"validation"`
	IsHTML     string `json:"is_html"`
}

func (h Handler) SendMail(c *gin.Context) {
	var (
		req        request
		validation = false
		isHTML     = false
	)

	if err := c.BindJSON(&req); err != nil {
		_ = c.Error(err)
	}

	env := mail.Envelope{
		From:    req.From,
		Subject: req.Subject,
		Text:    req.Text,
		To:      req.To,
	}

	if key, ok := c.Get("validation"); ok {
		if key == "yes" {
			validation = true
		}
	}

	if key, ok := c.Get("is_html"); ok {
		if key == "yes" {
			isHTML = true
		}
	}

	id, err := h.Mailer.Send(env, validation, isHTML)
	if err != nil {
		_ = c.Error(err)
	}

	c.String(http.StatusOK, id)
}
