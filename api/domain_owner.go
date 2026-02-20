package main

import (
	"net/http"
	"time"
)

func domainOwnerAdd(domain string, ownerHex string, addDate time.Time) error {
	if domain == "" || ownerHex == "" {
		return errorMissingField
	}

	// Check if the owner exists
	_, err := ownerGetByOwnerHex(ownerHex)
	if err != nil {
		return err
	}

	// Check if already an owner
	statement := `
		SELECT EXISTS (
			SELECT 1 FROM domainOwners
			WHERE domain=$1 AND ownerHex=$2
		);
	`
	row := db.QueryRow(statement, domain, ownerHex)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		logger.Errorf("cannot check if already domain owner: %v", err)
		return errorInternal
	}

	if exists {
		return errorDomainOwnerAlreadyExists
	}

	statement = `
		INSERT INTO domainOwners (domain, ownerHex, addDate)
		VALUES ($1, $2, $3);
	`
	_, err = db.Exec(statement, domain, ownerHex, addDate)
	if err != nil {
		logger.Errorf("cannot add domain owner: %v", err)
		return errorInternal
	}

	return nil
}

func domainOwnerAddByEmail(domain string, email string) error {
	if domain == "" || email == "" {
		return errorMissingField
	}

	// Look up owner by email
	statement := `
		SELECT ownerHex FROM owners
		WHERE email=$1;
	`
	row := db.QueryRow(statement, email)
	var ownerHex string
	if err := row.Scan(&ownerHex); err != nil {
		return errorNoSuchOwner
	}

	return domainOwnerAdd(domain, ownerHex, time.Now().UTC())
}

func domainOwnerRemove(domain string, ownerHex string) error {
	if domain == "" || ownerHex == "" {
		return errorMissingField
	}

	// Check that there will be at least one owner left
	statement := `
		SELECT COUNT(*) FROM domainOwners
		WHERE domain=$1;
	`
	row := db.QueryRow(statement, domain)
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Errorf("cannot count domain owners: %v", err)
		return errorInternal
	}

	if count <= 1 {
		return errorCannotRemoveLastOwner
	}

	statement = `
		DELETE FROM domainOwners
		WHERE domain=$1 AND ownerHex=$2;
	`
	result, err := db.Exec(statement, domain, ownerHex)
	if err != nil {
		logger.Errorf("cannot remove domain owner: %v", err)
		return errorInternal
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Errorf("cannot get rows affected: %v", err)
		return errorInternal
	}

	if rowsAffected == 0 {
		return errorNoSuchOwner
	}

	return nil
}

func domainOwnerAddHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		OwnerToken *string `json:"ownerToken"`
		Domain     *string `json:"domain"`
		Email      *string `json:"email"`
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
	isOwner, err := domainOwnershipVerify(o.OwnerHex, domain)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	if !isOwner {
		bodyMarshal(w, response{"success": false, "message": errorNotAuthorised.Error()})
		return
	}

	if err = domainOwnerAddByEmail(domain, *x.Email); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	bodyMarshal(w, response{"success": true})
}

func domainOwnerRemoveHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		OwnerToken    *string `json:"ownerToken"`
		Domain        *string `json:"domain"`
		RemoveOwnerHex *string `json:"removeOwnerHex"`
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
	isOwner, err := domainOwnershipVerify(o.OwnerHex, domain)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	if !isOwner {
		bodyMarshal(w, response{"success": false, "message": errorNotAuthorised.Error()})
		return
	}

	if err = domainOwnerRemove(domain, *x.RemoveOwnerHex); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	bodyMarshal(w, response{"success": true})
}
