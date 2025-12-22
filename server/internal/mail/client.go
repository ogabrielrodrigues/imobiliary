package mail

import (
	"bytes"
	"context"
	"imobiliary/config"
	"imobiliary/internal/application/httperr"
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
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	return &Mail{
		addr:     config.GetMailAddr(),
		host:     config.GetMailHost(),
		port:     config.GetMailPort(),
		user:     config.GetMailUser(),
		password: config.GetMailPassword(),
	}, nil
}

func (m *Mail) SendMail(message *mail.Message) *httperr.HttpError {
	dialer := mail.NewDialer(m.host, m.port, m.user, m.password)

	if err := dialer.DialAndSend(message); err != nil {
		return httperr.NewInternalServerError(context.Background(), "fail to send mail")
	}

	return nil
}

func (m *Mail) ParseHTML(file string) (string, error) {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = tmpl.Execute(buf, nil); err != nil {
		return "", err
	}

	return buf.String(), nil
}
