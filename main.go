package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"

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

var kanaConv = unicode.SpecialCase{
	unicode.CaseRange{
		Lo: 0x3041, // ぁ
		Hi: 0x3093, // ん
		Delta: [unicode.MaxCase]rune{
			0x30a1 - 0x3041, // strings.ToUpperCase でカタカナに変換されるマッピング
			0,               // strings.ToLowerCase に対応 (今回は使わないことにして 0 を書いておく)
			0,               // strings.ToTitleCase に対応 (今回は使わないことにして 0 を書いておく)
		},
	},
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
					p, errs := get.Select(strings.ToUpperSpecial(kanaConv, message.Text))
					if errs != nil {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("%v", errs))).Do(); err != nil {
							log.Print(err)
						}
					} else {
						pokemon := p.No + p.Name
						pokemondata := p.H + "-" + p.A + "-" + p.B + "-" + p.C + "-" + p.D + "-" + p.S + "-" + p.Sum
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("%s\n%s", pokemon, pokemondata))).Do(); err != nil {
							log.Print(err)
						}
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
