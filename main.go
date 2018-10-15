package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ryusei1026/Pokemonbot/get"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
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
					p, err := get.Select(message.Text)
					if err != nil {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("%v", err))).Do(); err != nil {
							log.Print(err)
						}
					} else {
						pokemon := p.No + p.Name + p.H + p.A + p.B + p.C + p.D + p.S + p.Sum
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(pokemon)).Do(); err != nil {
							log.Print(err)
						}
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}