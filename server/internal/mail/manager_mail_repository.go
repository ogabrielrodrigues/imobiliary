package mail

import (
	"context"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/types"
	"path/filepath"

	"gopkg.in/mail.v2"
)

type MailManagerRepository struct {
	client *Mail
}

func NewMailManagerRepository(client *Mail) MailManagerRepository {
	return MailManagerRepository{client}
}

func (mr *MailManagerRepository) ResetPasswordMessage(email types.Email) *httperr.HttpError {
	message := mail.NewMessage()
	message.SetHeader("From", "not-reply@imobiliary.com")
	message.SetHeader("To", email.Value())
	message.SetHeader("Subject", "Redefina sua senha")

	html, err := mr.client.ParseHTML(filepath.Join("internal", "mail", "templates", "reset-password.html"))
	if err != nil {
		return httperr.NewInternalServerError(context.Background(), "error on sending the reset password email")
	}

	message.SetBody("text/html", html)

	return mr.client.SendMail(message)
}
