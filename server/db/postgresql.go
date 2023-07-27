package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:secret@localhost:5433/go-chat")
	if err != nil {
		return nil, err
	}
	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() error {
	if err := d.db.Close(); err != nil {
		return err
	}
	return nil
}

func (d *Database) GetDB() *sql.DB {
	return d.db 
}
