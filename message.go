package imsgbridge

import "time"

type Message struct {
	ID int64

	From string

	Text string

	Time time.Time

	IsFromMe bool
}
