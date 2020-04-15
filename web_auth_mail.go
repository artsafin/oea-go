package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"net/url"
	"oea-go/common"
	"time"
)

const (
	AuthSubjectLine = "OEA Authentication"
)

type emailData struct {
	Subject   string
	From      string
	To        string
	Link      string
	IP        string
	Timestamp string
}

func sendMail(ip string, link url.URL, recipient string, cfg common.Config) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		cfg.SmtpUser,
		cfg.SmtpPass,
		cfg.SmtpHost,
	)

	emailData := emailData{
		Subject:   AuthSubjectLine,
		From:      cfg.SmtpUser,
		To:        recipient,
		IP:        ip,
		Timestamp: time.Now().Format(time.UnixDate),
		Link:      link.String(),
	}

	emailBuf := &bytes.Buffer{}
	emailTpl := template.Must(template.New("email").Parse(string(common.MustAsset("resources/email.go.html"))))
	if tplErr := emailTpl.Execute(emailBuf, emailData); tplErr != nil {
		panic(tplErr)
	}
	if emailBuf.Len() == 0 {
		panic("couldn't parse template")
	}

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	return smtp.SendMail(
		fmt.Sprintf("%s:%d", cfg.SmtpHost, cfg.SmtpPort),
		auth,
		cfg.SmtpUser,
		[]string{recipient},
		emailBuf.Bytes(),
	)
}
