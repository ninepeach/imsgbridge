package sender

import (
	"fmt"
	"os/exec"
)

type macOSSender struct{}

func New() Sender {

	return &macOSSender{}

}

func (s *macOSSender) Send(
	target string,
	text string,
) error {

	script := fmt.Sprintf(`
tell application "Messages"

	set targetService to 1st service whose service type is iMessage


	try

		set targetBuddy to buddy "%s" of targetService


		send "%s" to targetBuddy


	on error errMsg

		error "NOT_IMESSAGE"

	end try


end tell
`,
		escape(target),
		escape(text),
	)

	cmd := exec.Command(
		"osascript",
		"-e",
		script,
	)

	err := cmd.Run()

	if err != nil {

		if contains(
			err.Error(),
			"NOT_IMESSAGE",
		) {

			return ErrNotIMessage

		}

		return err
	}

	return nil
}

func escape(
	s string,
) string {

	result := ""

	for _, c := range s {

		switch c {

		case '\\':

			result += "\\\\"

		case '"':

			result += "\\\""

		case '\n':

			result += "\\n"

		default:

			result += string(c)

		}

	}

	return result
}

func contains(
	s string,
	sub string,
) bool {

	return len(s) >= len(sub) &&
		find(s, sub)

}

func find(
	s string,
	sub string,
) bool {

	for i := 0; i <= len(s)-len(sub); i++ {

		if s[i:i+len(sub)] == sub {

			return true

		}

	}

	return false
}
