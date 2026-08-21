package applemsg

import "time"

type Direction string

const (
	Incoming Direction = "incoming"
	Outgoing Direction = "outgoing"
)

type Identity struct {
	Handle string
	Type   string
}

type Message struct {
	ID int64

	GUID string

	Sender Identity

	Service string

	Text string

	Direction Direction

	Time time.Time
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
