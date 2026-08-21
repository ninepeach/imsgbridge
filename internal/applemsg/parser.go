package applemsg

import "time"


func Parse(
	raw RawMessage,
) Message {


	text := NormalizeText(
		raw.Text,
		raw.AttributedBody,
	)


	direction := Incoming

	if raw.IsFromMe {

		direction = Outgoing

	}


	return Message{

		ID: raw.ID,

		Text: text,


		Sender: Handle{

			Handle: raw.Handle,

			Type: raw.HandleType,

			Service: raw.Service,

		},


		Time: convertAppleTime(
			raw.Date,
		),


		Direction: direction,

	}

}





// NormalizeText returns the user visible message text.
//
// Apple Messages stores normal text in message.text.
// attributedBody is a private NSAttributedString archive.
// Do not parse it here until a proper decoder exists.
func NormalizeText(
	text string,
	attributedBody []byte,
) string {


	if text != "" {

		return text

	}


	// Ignore attributedBody for now.
	// Avoid returning Apple internal metadata such as:
	// __kIMMessagePartAttributeName

	return ""

}





func convertAppleTime(
	value int64,
) time.Time {


	// Apple timestamp:
	// seconds since 2001-01-01 00:00:00 UTC

	return time.Unix(
		value+978307200,
		0,
	)

}
