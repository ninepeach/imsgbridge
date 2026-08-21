package main

import (
	"log"

	"github.com/ninepeach/imsgbridge"
)


func main() {


	client, err := imsgbridge.NewClient(
		imsgbridge.Config{},
	)


	if err != nil {

		log.Fatal(err)

	}



	client.OnMessage(
		func(msg imsgbridge.Message) {


			// ignore empty
			if msg.Text == "" {

				return

			}


			// ignore messages sent by assistant/self
			if msg.IsFromMe {

				return

			}



			log.Printf(
				"FROM=%s TEXT=%s",
				msg.From,
				msg.Text,
			)



			err := client.SendAssistant(
				msg.From,
				"Echo: "+msg.Text,
			)



			if err != nil {

				log.Println(
					"send failed:",
					err,
				)

			}

		},
	)



	log.Println(
		"iMsgBridge example running...",
	)



	if err := client.Start(); err != nil {

		log.Fatal(err)

	}

}
