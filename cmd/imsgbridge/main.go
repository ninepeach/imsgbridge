package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ninepeach/imsgbridge/internal/applemsg"
	"github.com/ninepeach/imsgbridge/internal/state"
)

func main() {


	reader, err := applemsg.NewReader()

	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()



	store, err := state.NewStore()

	if err != nil {
		log.Fatal(err)
	}



	s, err := store.Load()

	if err != nil {
		log.Fatal(err)
	}



	// 第一次启动
	if s.LastMessageID == 0 {

		id, err := reader.LastID()

		if err != nil {
			log.Fatal(err)
		}

		s.LastMessageID = id

		err = store.Save(s)

		if err != nil {
			log.Fatal(err)
		}
	}



	fmt.Println("iMsgBridge watching...")
	fmt.Println(
		"Start ROWID:",
		s.LastMessageID,
	)



	ticker := time.NewTicker(
		1 * time.Second,
	)

	defer ticker.Stop()



	for range ticker.C {


		messages, err := reader.ReadAfter(
			s.LastMessageID,
		)


		if err != nil {

			log.Println(
				"read error:",
				err,
			)

			continue
		}



		for _, msg := range messages {


			direction := "IN"

			if msg.FromMe {
				direction = "OUT"
			}



			fmt.Printf(
				"%d [%s] %s %s %q\n",
				msg.ID,
				direction,
				msg.Service,
				msg.Handle,
				msg.Text,
			)



			// 消息处理完成后保存
			if msg.ID > s.LastMessageID {

				s.LastMessageID = msg.ID

				err := store.Save(s)

				if err != nil {

					log.Println(
						"save state error:",
						err,
					)
				}
			}
		}
	}
}
