package applemsg

func Parse(raw RawMessage) Message {

	direction := Incoming

	if raw.IsFromMe {
		direction = Outgoing
	}

	return Message{
		ID:   raw.ID,
		GUID: raw.GUID,

		Sender: Identity{
			Handle: raw.Handle,
			Type:   raw.HandleType,
		},

		Service: raw.Service,

		Text: raw.Text,

		Direction: direction,

		Time: ParseAppleTime(raw.Date),
	}
}
