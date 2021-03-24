package main

import (
	"bytes"
	"fmt"
	ht "html/template"
	"os"
)

type emailNotificationPlugs struct {
	Origin               string
	Kind                 string
	UnsubscribeSecretHex string
	Domain               string
	Path                 string
	CommentHex           string
	CommenterName        string
	Title                string
	Html                 ht.HTML
}

func smtpEmailNotification(to string, toName string, kind string, domain string, path string, commentHex string, commenterName string, title string, html string, unsubscribeSecretHex string) error {
	t, err := ht.ParseFiles(fmt.Sprintf("%s/templates/email-notification.txt", os.Getenv("STATIC")))
	if err != nil {
		logger.Errorf("cannot parse %s/templates/email-notification.txt: %v", os.Getenv("STATIC"), err)
		return errorMalformedTemplate
	}

	var body bytes.Buffer
	err = t.Execute(&body, &emailNotificationPlugs{
		Origin:               os.Getenv("ORIGIN"),
		Kind:                 kind,
		Domain:               domain,
		Path:                 path,
		CommentHex:           commentHex,
		CommenterName:        commenterName,
		Title:                title,
		Html:                 ht.HTML(html),
		UnsubscribeSecretHex: unsubscribeSecretHex,
	})
	if err != nil {
		logger.Errorf("error generating templated HTML for email notification: %v", err)
		return err
	}

	err = smtpSendMail(to, toName, "text/html; charset=UTF-8", "[Commento] "+title, body.String())
	if err != nil {
		logger.Errorf("cannot send email notification email: %v", err)
		return errorCannotSendEmail
	}

	return nil
}
