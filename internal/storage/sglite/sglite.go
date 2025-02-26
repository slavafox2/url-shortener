package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	// "url-shortener/internal/storage"
	// _ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) Close() error {
    if s.db != nil {
        return s.db.Close()
    }
    return nil
}


func New(storagePath string) (*Storage, error) {

	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "url" (
			"id" INTEGER PRIMARY KEY,
			"alias" TEXT NOT NULL UNIQUE,
			"url" TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS "idx_alias" ON "url"("alias");
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Create index with proper quotes
	// _, err = db.Exec(`
    //     CREATE INDEX IF NOT EXISTS "idx_alias" ON "url"("alias");
    // `)

	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", op, err)
	// }

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
