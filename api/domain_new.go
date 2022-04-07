package main

import (
	"net/http"
	"os"
	"strings"
	"time"
)

func domainNew(ownerHex string, name string, domain string) error {
	if ownerHex == "" || name == "" || domain == "" {
		return errorMissingField
	}

	if strings.Contains(domain, "/") {
		return errorInvalidDomain
	}

	// if asked to disable wildcards, then don't allow them...
	if os.Getenv("ENABLE_WILDCARDS") == "false" {
		if strings.Contains(domain, "%") || strings.Contains(domain, "_") {
			return errorInvalidDomain
		}
	}

	// test if domain already exists
	statement := `
		SELECT COUNT(*) FROM
		domains WHERE
		canon(regexp_replace($1, '[%]', '')) LIKE canon(domain) OR canon(domain) LIKE canon($1);
	`
	row := db.QueryRow(statement, domain)
	var err error
	var count int

	if err = row.Scan(&count); err != nil {
		return errorInvalidDomain
	}

	if count > 0 {
		return errorDomainAlreadyExists
	}

	statement = `
		INSERT INTO
		domains (ownerHex, name, domain, creationDate)
		VALUES  ($1,       $2,   $3,     $4          );
	`
	_, err = db.Exec(statement, ownerHex, name, domain, time.Now().UTC())
	if err != nil {
		// TODO: This should not happen given the above check, so this is likely not the error. Be more informative?
		return errorDomainAlreadyExists
	}

	return nil
}

func domainNewHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		OwnerToken *string `json:"ownerToken"`
		Name       *string `json:"name"`
		Domain     *string `json:"domain"`
	}

	var x request
	if err := bodyUnmarshal(r, &x); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	o, err := ownerGetByOwnerToken(*x.OwnerToken)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	domain := domainStrip(*x.Domain)

	if err = domainNew(o.OwnerHex, *x.Name, domain); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	if err = domainModeratorNew(domain, o.Email); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	bodyMarshal(w, response{"success": true, "domain": domain})
}
