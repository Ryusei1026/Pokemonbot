package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ryusei1026/Pokemonbot/get"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Post struct {
	No   string
	Name string
	H    string
	A    string
	B    string
	C    string
	D    string
	S    string
	Sum  string
}

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
					var p get.Post
					p, _ = get.Select(message.Text)
					pokemon := p.No + p.Name
					pokemondata := p.H + "-" + p.A + "-" + p.B + "-" + p.C + "-" + p.D + "-" + p.S + "-" + p.Sum
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("%s\n%s", pokemon, pokemondata))).Do(); err != nil {
						log.Print(p)
						log.Print(err)
					}
				}
			}
		}
	}) // This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
