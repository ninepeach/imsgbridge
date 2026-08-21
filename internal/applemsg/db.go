package applemsg

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type DB struct {
	conn *sql.DB
}

func OpenDB() (*DB, error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	path := filepath.Join(
		home,
		"Library",
		"Messages",
		"chat.db",
	)

	db, err := sql.Open(
		"sqlite",
		"file:"+path+"?mode=ro",
	)

	if err != nil {
		return nil, ErrDatabaseUnavailable
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, ErrPermissionDenied
	}

	return &DB{
		conn: db,
	}, nil
}

func (d *DB) Close() error {

	if d.conn != nil {
		return d.conn.Close()
	}

	return nil
}
