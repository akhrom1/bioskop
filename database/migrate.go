package database

import (
	"database/sql"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

func RunMigration(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Printf("Applied %d migrations\n", n)
}
