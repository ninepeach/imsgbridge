package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ninepeach/imsgbridge/internal/applemsg"
)

func main() {

	db, err := applemsg.OpenDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	reader := applemsg.NewReader(db)

	lastID, err := reader.LastID()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		"iMsgBridge watching...",
	)

	fmt.Println(
		"Start ROWID:",
		lastID,
	)

	ticker := time.NewTicker(
		time.Second,
	)

	defer ticker.Stop()

	for range ticker.C {

		raws, err := reader.ReadAfter(lastID)

		if err != nil {
			log.Println(err)
			continue
		}

		for _, raw := range raws {

			msg := applemsg.Parse(raw)

			fmt.Printf(
				"%d [%s] %s %s %q\n",
				msg.ID,
				msg.Direction,
				msg.Service,
				msg.Sender.Handle,
				msg.Text,
			)

			if msg.ID > lastID {
				lastID = msg.ID
			}
		}
	}
}
