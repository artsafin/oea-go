package auth

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"net/url"
	"oea-go/internal/common"
	"time"
)

const (
	SubjectLine = "OEA Authentication"
)

type emailData struct {
	Subject   string
	From      string
	To        string
	Link      string
	IP        string
	Timestamp string
}

func sendMail(inf common.AuthInfo, link url.URL, recipient string, cfg *common.Config) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		cfg.SmtpUser,
		cfg.SmtpPass,
		cfg.SmtpHost,
	)

	emailData := emailData{
		Subject:   SubjectLine,
		From:      cfg.SmtpUser,
		To:        recipient,
		IP:        inf.IP,
		Timestamp: inf.TS.Format(time.UnixDate),
		Link:      link.String(),
	}

	emailBuf := &bytes.Buffer{}
	emailTpl := template.Must(template.New("email").Parse(string(common.MustAsset("resources/email.go.html"))))
	if tplErr := emailTpl.Execute(emailBuf, emailData); tplErr != nil {
		return tplErr
	}
	if emailBuf.Len() == 0 {
		return errors.New("couldn't parse template")
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
