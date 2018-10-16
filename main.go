package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ryusei1026/Pokemonbot/get"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		"7e5bb6e7b67a07dc929db47a65c3d970",
		"nY0wki+j6Kz9KnTb/MWlbE1rhpVvmn/Ywtd1xem2LYjb8MV4x/fDJBk4Rj5OXrBdT5X2XTi+pWpgmVqt5a25P4yZvYGx1V+PAdkjOYgNIWG21/oqvhBtA4YY0V6QnOKXDLsuSFGJUN9KbM0Rrnqn6QdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Println("111")
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					p, errs := get.Select(message.Text)
					if errs == nil {
						pokemon := (p.No + p.Name + p.H + p.A + p.B + p.C + p.D + p.S + p.Sum)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(pokemon)).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("%v", errs))).Do(); err != nil {
							log.Println("222")
							log.Print(err)
							log.Println("aaa")
						}
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
