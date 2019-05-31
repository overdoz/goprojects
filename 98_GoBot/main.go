package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "854869015:AAGS9yhBdhkzKBsAswLZEi38H949yvOvx5I",
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		URL: "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Chat, "WHAAAATUP BRAA")
	})

	b.Handle("/fuckyou", func(m *tb.Message) {
		b.Send(m.Chat, "FUCK YOU TOOOOO 3===o")
	})

	b.Handle("/whack", func(m *tb.Message) {
		b.Send(m.Chat, "find dich auch whack")
	})

	we := tb.InlineButton{
		Unique: "Gras",
		Text: "Gras",
	}
	ko := tb.InlineButton{
		Unique: "Koks",
		Text: "Koks",
	}
	bl := tb.InlineButton{
		Unique: "Blowjob",
		Text: "Blowjob",
	}
	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{we},
		[]tb.InlineButton{ko},
		[]tb.InlineButton{bl},
	}


	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		b.Send(m.Chat, "Brauchste wat?!", &tb.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
	})

	b.Handle(&we, func(c *tb.Callback) {
		// on inline button pressed (callback!)

		// always respond!
		b.Respond(c, &tb.CallbackResponse{Text: "Bin unterwegs!"})
	})
	b.Handle(&ko, func(c *tb.Callback) {
		// on inline button pressed (callback!)

		// always respond!
		b.Respond(c, &tb.CallbackResponse{Text: "Bin unterwegs!"})
	})
	b.Handle(&bl, func(c *tb.Callback) {
		// on inline button pressed (callback!)

		// always respond!
		b.Respond(c, &tb.CallbackResponse{Text: "Bin unterwegs!"})
	})



	b.Handle("/ferdistinkt", func(m *tb.Message) {
		b.Send(m.Chat, "find ich auch")
		b.Send(m.Sender, "Bruder, findest du nicht auch, dass Ferdi stinkt?!")
	})

	b.Start()
}