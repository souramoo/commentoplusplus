package main

import ()

var domainsRowColumns = `
	domains.domain,
	domains.name,
	domains.creationDate,
	domains.state,
	domains.importedComments,
	domains.autoSpamFilter,
	domains.requireModeration,
	domains.requireIdentification,
	domains.moderateAllAnonymous,
	domains.emailNotificationPolicy,
	domains.commentoProvider,
	domains.googleProvider,
	domains.twitterProvider,
	domains.githubProvider,
	domains.gitlabProvider,
	domains.ssoProvider,
	domains.ssoSecret,
	domains.ssoUrl,
	domains.defaultSortPolicy
`

func domainsRowScan(s sqlScanner, d *domain) error {
	return s.Scan(
		&d.Domain,
		&d.Name,
		&d.CreationDate,
		&d.State,
		&d.ImportedComments,
		&d.AutoSpamFilter,
		&d.RequireModeration,
		&d.RequireIdentification,
		&d.ModerateAllAnonymous,
		&d.EmailNotificationPolicy,
		&d.CommentoProvider,
		&d.GoogleProvider,
		&d.TwitterProvider,
		&d.GithubProvider,
		&d.GitlabProvider,
		&d.SsoProvider,
		&d.SsoSecret,
		&d.SsoUrl,
		&d.DefaultSortPolicy,
	)
}

func domainOwnerList(dmn string) ([]domainOwner, error) {
	if dmn == "" {
		return []domainOwner{}, errorMissingField
	}

	statement := `
		SELECT domainOwners.ownerHex, owners.email, owners.name, domainOwners.addDate
		FROM domainOwners
		INNER JOIN owners ON domainOwners.ownerHex = owners.ownerHex
		WHERE domainOwners.domain = $1
		ORDER BY domainOwners.addDate ASC;
	`
	rows, err := db.Query(statement, dmn)
	if err != nil {
		logger.Errorf("cannot query domain owners: %v", err)
		return nil, errorInternal
	}
	defer rows.Close()

	owners := []domainOwner{}
	for rows.Next() {
		var o domainOwner
		if err = rows.Scan(&o.OwnerHex, &o.Email, &o.Name, &o.AddDate); err != nil {
			logger.Errorf("cannot scan domain owner: %v", err)
			return nil, errorInternal
		}
		owners = append(owners, o)
	}

	return owners, rows.Err()
}

func domainGet(dmn string) (domain, error) {
	if dmn == "" {
		return domain{}, errorMissingField
	}

	statement := `
		SELECT ` + domainsRowColumns + `
		FROM domains
		WHERE canon($1) LIKE canon(domain);
	`
	row := db.QueryRow(statement, dmn)

	var err error
	d := domain{}
	if err = domainsRowScan(row, &d); err != nil {
		return d, errorNoSuchDomain
	}

	d.Moderators, err = domainModeratorList(d.Domain)
	if err != nil {
		return domain{}, err
	}

	d.Owners, err = domainOwnerList(d.Domain)
	if err != nil {
		return domain{}, err
	}

	return d, nil
}
