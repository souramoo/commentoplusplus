package main

import (
	"bytes"
	"os"
)

type domainExportErrorPlugs struct {
	Origin string
	Domain string
}

func smtpDomainExportError(to string, toName string, domain string) error {
	var body bytes.Buffer
	templates["data-export-error"].Execute(&body, &domainExportPlugs{Origin: os.Getenv("ORIGIN")})

	err := smtpSendMail(to, toName, "", "Commento Data Export", body.String())
	if err != nil {
		logger.Errorf("cannot send data export error email: %v", err)
		return errorCannotSendEmail
	}

	return nil
}
