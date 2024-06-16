package emails

import (
	"fmt"
	"github.com/wisle25/media-stock-be/applications/emails"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/domains/entity"
	"net/smtp"
)

type SmtpEmailService struct /* implements EmailService */ {
	smtpHost  string
	smtpPort  string
	auth      smtp.Auth
	emailFrom string
}

func NewStmpEmailService(config *commons.Config) emails.EmailService {
	auth := smtp.PlainAuth("", config.SmtpUsername, config.SmtpPassword, config.SmtpHost)

	return &SmtpEmailService{
		smtpHost:  config.SmtpHost,
		smtpPort:  config.SmtpPort,
		auth:      auth,
		emailFrom: "noreply@anonCreation.com",
	}
}

func (s *SmtpEmailService) SendEmail(email entity.Email) {
	// Arrange email
	msg := []byte("From: " + s.emailFrom + "\n" +
		"To: " + email.To + "\n" +
		"Subject: " + email.Subject + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"<html><body>" + email.Body + "</body></html>",
	)

	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)

	// Send
	err := smtp.SendMail(addr, s.auth, s.emailFrom, []string{email.To}, msg)
	if err != nil {
		panic(fmt.Errorf("send_email_err: sending email: %v", err))
	}
}
