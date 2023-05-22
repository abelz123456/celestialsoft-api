package helpers

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type MailHelper struct {
	host string
	port int
	auth smtp.Auth
	from string
	log  log.Log
}

func NewMailHelper(cfg config.Config) *MailHelper {
	logger := log.NewLog()
	if cfg.SmtpHost == "" || cfg.SmtpPort == 0 || cfg.SmtpUser == "" || cfg.SmtpPass == "" {
		logger.Panic(errors.New("please provide SMTP_HOST, SMTP_PORT, SMTP_USER and SMTP_PASS to use MailHelper"), "", map[string]interface{}{
			"SMTP_HOST": cfg.SmtpHost,
			"SMTP_PORT": cfg.SmtpPort,
			"SMTP_USER": strings.ToLower(cfg.SmtpUser),
			"SMTP_PASS": cfg.SmtpPass,
		})
	}

	return &MailHelper{
		host: cfg.SmtpHost,
		port: cfg.SmtpPort,
		auth: smtp.PlainAuth("", cfg.SmtpUser, cfg.SmtpPass, cfg.SmtpHost),
		from: fmt.Sprintf("%s <%s>", cases.Title(language.English).String(cfg.MailFromName), strings.ToLower(cfg.MailFromAddr)),
		log:  logger,
	}
}

func (m *MailHelper) Send(to []string, subject string, body string) error {
	content := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", m.from, strings.Join(to, ","), subject, body)
	return smtp.SendMail(fmt.Sprintf("%s:%d", m.host, m.port), m.auth, m.from, to, []byte(content))
}
