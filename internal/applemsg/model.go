package applemsg

import "time"

type Direction int

const (
	Incoming Direction = iota
	Outgoing
)

type Handle struct {
	Handle string

	Type string

	Service string
}

type RawMessage struct {
	ID int64

	GUID string

	Text string

	AttributedBody []byte

	Handle string

	HandleType string

	Service string

	IsFromMe bool

	Date int64
}

type Message struct {
	ID int64

	Text string

	Sender Handle

	Time time.Time

	Direction Direction
}
