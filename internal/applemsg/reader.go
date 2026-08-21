package applemsg

import (
	"database/sql"
)

type Reader struct {
	db *DB
}

func NewReader(db *DB) *Reader {

	return &Reader{
		db: db,
	}
}

func (r *Reader) LastID() (int64, error) {

	var id sql.NullInt64

	err := r.db.conn.QueryRow(`
		SELECT COALESCE(MAX(ROWID),0)
		FROM message
	`).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id.Int64, nil
}

func (r *Reader) ReadAfter(
	lastID int64,
) ([]RawMessage, error) {

	rows, err := r.db.conn.Query(`
		SELECT
			m.ROWID,
			m.guid,
			COALESCE(m.text,''),
			m.attributedBody,
			COALESCE(h.id,''),
			COALESCE(h.service,''),
			m.is_from_me,
			m.date

		FROM message m

		LEFT JOIN handle h
			ON m.handle_id=h.ROWID

		WHERE m.ROWID > ?

		ORDER BY m.ROWID ASC

		LIMIT 100
	`,
		lastID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []RawMessage

	for rows.Next() {

		var (
			id      int64
			guid    string
			text    string
			attr    []byte
			handle  string
			service string
			fromMe  int
			date    int64
		)

		err := rows.Scan(
			&id,
			&guid,
			&text,
			&attr,
			&handle,
			&service,
			&fromMe,
			&date,
		)

		if err != nil {
			return nil, err
		}

		result = append(
			result,
			RawMessage{

				ID: id,

				GUID: guid,

				Text: text,

				AttributedBody: attr,

				Handle: handle,

				Service: service,

				HandleType: detectHandleType(handle),

				IsFromMe: fromMe == 1,

				Date: date,
			},
		)
	}

	return result, rows.Err()
}

func detectHandleType(
	handle string,
) string {

	if len(handle) > 0 && handle[0] == '+' {
		return "phone"
	}

	return "email"
}
