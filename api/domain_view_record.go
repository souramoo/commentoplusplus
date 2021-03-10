package main

import (
	"os"
	"time"
)

func domainViewRecord(domain string, commenterHex string) {
	if os.Getenv("ENABLE_LOGGING") != "false" && os.Getenv("ENABLE_LOGGING") != "" {
		statement := `
			INSERT INTO
			views  (domain, commenterHex, viewDate)
			VALUES ($1,     $2,           $3      );
		`
		_, err := db.Exec(statement, domain, commenterHex, time.Now().UTC())
		if err != nil {
			logger.Warningf("cannot insert views: %v", err)
		}
	}
}
