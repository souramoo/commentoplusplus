package main

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
)

func smtpSendMail(toAddress string, toName string, contentType string, subject string, body string) error {
	from := mail.Address{"Commento", os.Getenv("SMTP_FROM_ADDRESS")}
	to := mail.Address{toName, toAddress}

	// Setup headers
	headers := make(map[string]string)
	headers["To"] = to.String()
	headers["Subject"] = subject
	if contentType == "" {
		headers["Content-Type"] = "text/plain; charset=UTF-8"
	} else {
		headers["Content-Type"] = contentType
	}

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")

	return smtp.SendMail(servername, smtpAuth, from.String(), []string{to.String()}, []byte(message))
}
