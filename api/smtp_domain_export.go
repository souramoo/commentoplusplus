package main

import (
	"bytes"
	"os"
)

type domainExportPlugs struct {
	Origin    string
	Domain    string
	ExportHex string
}

func smtpDomainExport(to string, toName string, domain string, exportHex string) error {
	var body bytes.Buffer
	templates["domain-export"].Execute(&body, &domainExportPlugs{Origin: os.Getenv("ORIGIN"), ExportHex: exportHex})

	err := smtpSendMail(to, toName, "", "Commento Data Export", body.String())
	if err != nil {
		logger.Errorf("cannot send data export email: %v", err)
		return errorCannotSendEmail
	}

	return nil
}
