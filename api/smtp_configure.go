package main

import (
	"os"
)

var smtpConfigured bool

func smtpConfigure() error {
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	if host == "" || port == "" {
		logger.Warningf("smtp not configured, no emails will be sent")
		smtpConfigured = false
		return nil
	}

	if os.Getenv("SMTP_FROM_ADDRESS") == "" {
		logger.Errorf("COMMENTO_SMTP_FROM_ADDRESS not set")
		smtpConfigured = false
		return errorMissingSmtpAddress
	}

	if username == "" || password == "" {
		logger.Warningf("no SMTP username/password set, Commento will assume they aren't required")
	}
	smtpConfigured = true
	return nil
}
