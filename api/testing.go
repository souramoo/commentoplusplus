package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/op/go-logging"
)

func failTestOnError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("failed test: %v", err)
	}
}

func getPublicTables() ([]string, error) {
	statement := `
		SELECT tablename
		FROM pg_tables
		WHERE schemaname='public';
	`
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot query public tables: %v", err)
		return []string{}, err
	}

	defer rows.Close()

	tables := []string{}
	for rows.Next() {
		var table string
		if err = rows.Scan(&table); err != nil {
			fmt.Fprintf(os.Stderr, "cannot scan table name: %v", err)
			return []string{}, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}

func clearDatabase() error {
	tables, err := getPublicTables()
	if err != nil {
		return err
	}

	// Drop all tables so we always work with the same data.
	// That includes the `migrations` table` so that all migrations re-run cleanly.
	for _, table := range tables {
		_, err = db.Exec(fmt.Sprintf("DROP TABLE %s CASCADE;", table))
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot drop %s: %v", table, err)
			return err
		}
	}

	// Drop any custom enum types so migrations can recreate them
	statement := `
		SELECT t.typname
		FROM pg_type t
		JOIN pg_namespace n ON t.typnamespace = n.oid
		WHERE n.nspname = 'public' AND t.typtype = 'e';
	`
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot query custom types: %v", err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var typeName string
		if err = rows.Scan(&typeName); err != nil {
			fmt.Fprintf(os.Stderr, "cannot scan type name: %v", err)
			return err
		}

		_, err = db.Exec(fmt.Sprintf("DROP TYPE %s CASCADE;", typeName))
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot drop type %s: %v", typeName, err)
			return err
		}
	}

	return nil
}

func setupTestDatabase() error {
	if os.Getenv("COMMENTO_POSTGRES") != "" {
		// set it manually because we need to use commento_test, not commento, by mistake
		os.Setenv("POSTGRES", os.Getenv("COMMENTO_POSTGRES"))
	} else {
		os.Setenv("POSTGRES", "postgres://postgres:postgres@localhost/commento_test?sslmode=disable")
	}

	if err := dbConnect(0); err != nil {
		return err
	}

	if err := clearDatabase(); err != nil {
		return err
	}

	// Recreate the migrations table since clearDatabase removes it
	dbCreateMigrationsTable()

	if err := migrateFromDir("../db/"); err != nil {
		return err
	}

	return nil
}

func clearTables() error {
	tables, err := getPublicTables()
	if err != nil {
		return err
	}

	for _, table := range tables {
		_, err = db.Exec(fmt.Sprintf("DELETE FROM %s;", table))
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot clear %s: %v", table, err)
			return err
		}
	}

	return nil
}

func createTestOwner() error {
	statement := `
		INSERT INTO owners (ownerHex, email, name, passwordHash, confirmedEmail, joinDate)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := db.Exec(statement, "temp-owner-hex", "test@test.com", "Test Owner", "dummy-hash", true, time.Now().UTC())
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create test owner: %v", err)
		return err
	}

	return nil
}

var setupComplete bool

func setupTestEnv() error {
	if !setupComplete {
		setupComplete = true

		if err := loggerCreate(); err != nil {
			return err
		}

		// Print messages to console only if verbose. Sounds like a good idea to
		// keep the console clean on `go test`.
		if !testing.Verbose() {
			logging.SetLevel(logging.CRITICAL, "")
		}

		if err := setupTestDatabase(); err != nil {
			return err
		}

		if err := markdownRendererCreate(); err != nil {
			return err
		}
	}

	if err := clearTables(); err != nil {
		return err
	}

	if err := createTestOwner(); err != nil {
		return err
	}

	hub = newHub()
	go hub.run()

	return nil
}
