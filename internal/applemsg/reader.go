package applemsg

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type Reader struct {
	db *sql.DB
}

func NewReader() (*Reader, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dbPath := filepath.Join(
		home,
		"Library",
		"Messages",
		"chat.db",
	)

	db, err := sql.Open(
		"sqlite",
		"file:"+dbPath+"?mode=ro",
	)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &Reader{
		db: db,
	}, nil
}

func (r *Reader) Close() error {
	return r.db.Close()
}


// 返回当前数据库最大的 ROWID
func (r *Reader) LastID() (int64, error) {

	var id sql.NullInt64

	err := r.db.QueryRow(`
		SELECT MAX(ROWID)
		FROM message
	`).Scan(&id)

	if err != nil {
		return 0, err
	}

	if !id.Valid {
		return 0, nil
	}

	return id.Int64, nil
}


// 读取 ROWID 之后的新消息
func (r *Reader) ReadAfter(lastID int64) ([]Message, error) {

	rows, err := r.db.Query(`
		SELECT
			m.ROWID,
			m.is_from_me,
			COALESCE(h.id,''),
			COALESCE(h.service,''),
			COALESCE(m.text,'')
		FROM message m

		LEFT JOIN handle h
			ON m.handle_id = h.ROWID

		WHERE m.ROWID > ?

		ORDER BY m.ROWID ASC
	`,
		lastID,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"query messages: %w",
			err,
		)
	}

	defer rows.Close()


	var messages []Message


	for rows.Next() {

		var (
			id       int64
			fromMe   int
			handle   string
			service  string
			text     string
		)


		err := rows.Scan(
			&id,
			&fromMe,
			&handle,
			&service,
			&text,
		)

		if err != nil {
			return nil, err
		}


		messages = append(
			messages,
			Message{
				ID:      id,
				Handle:  handle,
				Service: service,
				Text:    text,
				FromMe:  fromMe == 1,
			},
		)
	}


	if err := rows.Err(); err != nil {
		return nil, err
	}


	return messages, nil
}
