package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ninepeach/imsgbridge/internal/applemsg"
	"github.com/ninepeach/imsgbridge/internal/state"
)

func main() {

	db, err := applemsg.OpenDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	reader := applemsg.NewReader(db)

	service := applemsg.NewService(
		reader,
	)

	store, err := state.NewStore()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		"iMsgBridge watching...",
	)

	ctx := context.Background()

	err = service.Run(
		ctx,
		store,

		func(msg applemsg.Message) {

			fmt.Printf(
				"%d [%s] %s %q\n",
				msg.ID,
				msg.Direction,
				msg.Sender.Handle,
				msg.Text,
			)

		},
	)

	if err != nil {

		log.Fatal(err)

	}

}
