package mail

import (
	"bytes"
	"context"
	"fmt"
	"imobiliary/config"
	"imobiliary/internal/application/httperr"
	"strconv"
	"text/template"

	mail "gopkg.in/mail.v2"
)

type Mail struct {
	addr     string
	host     string
	port     int
	user     string
	password string
}

func NewMailClient() (*Mail, error) {
	config := config.GetConfig()

	mailPort, err := strconv.Atoi(config.MailPort)
	if err != nil {
		return nil, fmt.Errorf("mail port invalid")
	}

	return &Mail{
		addr:     config.MailAddr,
		host:     config.MailHost,
		port:     mailPort,
		user:     config.MailUser,
		password: config.MailPassword,
	}, nil
}

func CreateMessage(to, subject, body string) *mail.Message {
	config := config.GetConfig()

	message := mail.NewMessage()
	message.SetHeader("From", config.MailAddr)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	message.SetBody("text/html", body)

	return message
}

func (m *Mail) SendMail(message *mail.Message) *httperr.HttpError {
	dialer := mail.NewDialer(m.host, m.port, m.user, m.password)

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println(err)
		return httperr.NewInternalServerError(context.Background(), "fail to send mail")
	}

	return nil
}

func (m *Mail) ParseHTML(file string, data any) (string, error) {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
