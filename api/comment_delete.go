package main

import (
	"net/http"
	"time"
)

func commentDelete(commentHex string, deleterHex string, domain string, path string) error {
	if commentHex == "" || deleterHex == "" {
		return errorMissingField
	}

	statement := `
		UPDATE comments
		SET
			deleted = true,
			markdown = '[deleted]',
			html = '[deleted]',
			commenterHex = 'anonymous',
			deleterHex = $2,
			deletionDate = $3
		WHERE commentHex = $1;
	`
	_, err := db.Exec(statement, commentHex, deleterHex, time.Now().UTC())

	if err != nil {
		// TODO: make sure this is the error is actually non-existant commentHex
		return errorNoSuchComment
	}

	// Since we're no longer actually deleting comments, we are no longer running the trigger function!
	statement = `
		UPDATE pages
		SET commentCount = commentCount - 1
		WHERE canon($1) LIKE canon(domain) AND path = $2;
	`
	_, err = db.Exec(statement, domain, path)

	if err != nil {
		return errorNoSuchComment
	}

	hub.broadcast <- []byte(domain + path)

	return nil
}

func commentDeleteHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		CommenterToken *string `json:"commenterToken"`
		CommentHex     *string `json:"commentHex"`
	}

	var x request
	if err := bodyUnmarshal(r, &x); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	c, err := commenterGetByCommenterToken(*x.CommenterToken)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	cm, err := commentGetByCommentHex(*x.CommentHex)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	domain, path, err := commentDomainPathGet(*x.CommentHex)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	isModerator, err := isDomainModerator(domain, c.Email)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	if !isModerator && cm.CommenterHex != c.CommenterHex {
		bodyMarshal(w, response{"success": false, "message": errorNotModerator.Error()})
		return
	}

	if err = commentDelete(*x.CommentHex, *x.CommenterToken, domain, path); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	bodyMarshal(w, response{"success": true})
}

func commentOwnerDeleteHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		OwnerToken *string `json:"ownerToken"`
		CommentHex *string `json:"commentHex"`
	}

	var x request
	if err := bodyUnmarshal(r, &x); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	domain, path, err := commentDomainPathGet(*x.CommentHex)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	o, err := ownerGetByOwnerToken(*x.OwnerToken)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	isOwner, err := domainOwnershipVerify(o.OwnerHex, domain)
	if err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	if !isOwner {
		bodyMarshal(w, response{"success": false, "message": errorNotAuthorised.Error()})
		return
	}

	if err = commentDelete(*x.CommentHex, *x.OwnerToken, domain, path); err != nil {
		bodyMarshal(w, response{"success": false, "message": err.Error()})
		return
	}

	bodyMarshal(w, response{"success": true})
}
