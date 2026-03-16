package sql

import (
	"database/sql"
	_ "embed"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schemaSQL string

func InitDB(dbpath string) (*sql.DB, error) {
	dir := filepath.Dir(dbpath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", dbpath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(schemaSQL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
