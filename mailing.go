package ppl_billingandmailing

import (
	"net/smtp"
	"log"
)

func mailing() {
	auth := smtp.PlainAuth(
		"",
		"user@example.com",
		"password",
		"mail.example.com",
	)

	err := smtp.SendMail(
		"mail.example.com:25",
		auth,
		"sender@example.org",
		[]string{"recipien@example.org"},
		[]byte("Hello World!"),
	)

	if err != nil {
		log.Fatal(err)
	}
}