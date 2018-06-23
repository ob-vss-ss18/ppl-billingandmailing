package mailing

import (
	"github.com/scorredoira/email"
	"fmt"
	"net/smtp"
	"strconv"
	"net/mail"
	"os"
)

type EmailConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

func mailing(receivers []string, subject string, message string, attachments []string) {
	// authentication configuration
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, error := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpPass := os.Getenv("SMTP_PASS")
	smtpUser := os.Getenv("SMTP_USER")

	if error != nil {
		// handle error
		fmt.Println(error)
		os.Exit(2)
	}

	emailConf := &EmailConfig{smtpUser, smtpPass, smtpHost, smtpPort}

	emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)

	sender := os.Getenv("SMTP_USER")

	emailContent := email.NewMessage(subject, message)

	emailContent.From = mail.Address{Name: "From", Address: sender}
	emailContent.To = receivers

	for _, filename := range attachments {
		err := emailContent.Attach(filename)

		if err != nil {
			fmt.Println(err)
		}
	}

	// send out the email
	err := email.Send(smtpHost+":"+strconv.Itoa(emailConf.Port),
		emailauth,
		emailContent)

	if err != nil {
		fmt.Println(err)
	}
}