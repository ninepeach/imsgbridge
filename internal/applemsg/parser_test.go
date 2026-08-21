package applemsg

import (
	"testing"
)

func TestParseIncomingMessage(t *testing.T) {

	raw := RawMessage{
		ID: 100,

		GUID: "test-guid",

		Text: "hello",

		Handle: "+861234567890",

		HandleType: "phone",

		Service: "iMessage",

		IsFromMe: false,

		Date: 1750000000000000000,
	}

	msg := Parse(raw)

	if msg.ID != 100 {
		t.Fatalf(
			"unexpected id: %d",
			msg.ID,
		)
	}

	if msg.Text != "hello" {
		t.Fatalf(
			"unexpected text: %s",
			msg.Text,
		)
	}

	if msg.Direction != Incoming {
		t.Fatalf(
			"expected incoming",
		)
	}

	if msg.Sender.Handle != "+861234567890" {
		t.Fatalf(
			"unexpected handle",
		)
	}

}

func TestParseOutgoingMessage(t *testing.T) {

	raw := RawMessage{

		ID: 200,

		Text: "reply",

		Handle: "+81900000000",

		HandleType: "phone",

		Service: "iMessage",

		IsFromMe: true,
	}

	msg := Parse(raw)

	if msg.Direction != Outgoing {

		t.Fatalf(
			"expected outgoing",
		)
	}

}
