package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"os"
)

func smtpSendMail(toAddress string, toName string, contentType string, subject string, body string) error {
	from := mail.Address{"", os.Getenv("SMTP_FROM_ADDRESS")}
	to := mail.Address{"", toAddress}

	// Setup headers
	headers := make(map[string]string)
	headers["Subject"] = subject
	headers["From"] = "Commento <" + from.String() + ">"
	if toName == "" {
		headers["To"] = to.String()
	} else {
		headers["To"] = toName + " <" + to.String() + ">"
	}
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

	host, _, _ := net.SplitHostPort(servername)

	// TLS config
	tlsconfig := &tls.Config{
		ServerName: host,
	}
	if os.Getenv("SMTP_SKIP_HOST_VERIFY") == "true" {
		tlsconfig.InsecureSkipVerify = true
	}

	c, err := smtp.Dial(servername)
	if err != nil {
		return err
	}

	err = c.StartTLS(tlsconfig)
	if err != nil {
		return err
	}

	if smtpAuth != nil {
		err = c.Auth(smtpAuth)
		if err != nil {
			return err
		}
	}

	err = c.Mail(from.Address)
	if err != nil {
		return err
	}

	err = c.Rcpt(to.Address)
	if err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}
