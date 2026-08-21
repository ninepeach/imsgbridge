package sender

type Sender interface {
	Send(
		target string,
		text string,
	) error
}
