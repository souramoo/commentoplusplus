package main

import (
	"fmt"
	"os"
	"text/template"
)

var templates map[string]*template.Template

func smtpTemplatesLoad() error {
	names := []string{
		"confirm-hex",
		"reset-hex",
		"domain-export",
		"domain-export-error",
	}

	templates = make(map[string]*template.Template)

	logger.Infof("loading templates: %v", names)
	for _, name := range names {
		var err error
		templates[name] = template.New(name)
		templates[name], err = template.ParseFiles(fmt.Sprintf("%s/templates/%s.txt", os.Getenv("STATIC"), name))
		if err != nil {
			logger.Errorf("cannot parse %s/templates/%s.txt: %v", os.Getenv("STATIC"), name, err)
			return errorMalformedTemplate
		}
	}

	return nil
}
