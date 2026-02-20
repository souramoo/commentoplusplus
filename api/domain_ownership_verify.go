package main

import ()

func domainOwnershipVerify(ownerHex string, domain string) (bool, error) {
	if ownerHex == "" || domain == "" {
		return false, errorMissingField
	}

	statement := `
		SELECT EXISTS (
			SELECT 1
			FROM domainOwners
			INNER JOIN domains ON domainOwners.domain = domains.domain
			WHERE domainOwners.ownerHex=$1 AND $2 LIKE domains.domain
		);
	`
	row := db.QueryRow(statement, ownerHex, domain)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		logger.Errorf("cannot query if domain owner: %v", err)
		return false, errorInternal
	}

	return exists, nil
}
