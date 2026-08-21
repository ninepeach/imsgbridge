package applemsg

import "time"

type Message struct {
	ID        int64
	Handle    string
	Service   string
	Text      string
	FromMe    bool
	Timestamp time.Time
}
