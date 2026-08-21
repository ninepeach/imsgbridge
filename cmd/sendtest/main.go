package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/ninepeach/imsgbridge/internal/sender"
)

func main() {

	s := sender.New()

	err := s.Send(
		"email",
		"iMsgBridge test",
	)

	if err != nil {

		if errors.Is(
			err,
			sender.ErrNotIMessage,
		) {

			fmt.Println(
				"not iMessage target",
			)

			return
		}

		log.Fatal(err)

	}

	fmt.Println(
		"sent",
	)

}
