package services

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	gomail "gopkg.in/mail.v2"
)

type EmailService interface {
	SendConfirmEmail(email, name, pinCode string) error
}

type emailService struct {
	smtpHost     string
	smtpPort     int
	username     string
	password     string
	from         string
	templatePath string
}

func NewEmailService(smtpHost string, smtpPort int, username, password, from, templatePath string) *emailService {
	return &emailService{
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		username:     username,
		password:     password,
		from:         from,
		templatePath: templatePath,
	}
}

func (es *emailService) SendConfirmEmail(email, name, pinCode string) error {
	templateFile := filepath.Join(es.templatePath, "confirm_account.html")
	body, err := es.renderTemplate(templateFile, map[string]interface{}{
		"Name":    name,
		"PinCode": pinCode,
	})
	if err != nil {
		return err
	}

	subject := "Bem-vindo à nossa plataforma!"

	return es.sendEmail(email, subject, body)
}

func (es *emailService) renderTemplate(templateFile string, data map[string]interface{}) (string, error) {
	content, err := os.ReadFile(templateFile)
	if err != nil {
		return "", fmt.Errorf("FAILED_TO_READ_TEMPLATE_FILE: %v", err)
	}

	tmpl, err := template.New("email").Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("FAILED_TO_PARSE_TEMPLATE: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("FAILED_TO_EXECUTE_TEMPLATE: %v", err)
	}

	return buf.String(), nil
}

func (es *emailService) sendEmail(to, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", es.from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(es.smtpHost, es.smtpPort, es.username, es.password)

	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("FAILED_TO_SEND_EMAIL: %v", err)
	}

	return nil
}
