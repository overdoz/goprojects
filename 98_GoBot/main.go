package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type Definitions struct {
	Elements []Definition `json:"list"`
}

type Definition struct {
	Info string `json:"definition"`
}

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

	b.Handle("/dict", func(m *tb.Message) {
		temp := strings.Split(m.Payload, " ")
		payload := strings.Join(temp, "+")
		req, err := http.NewRequest("GET", "https://mashape-community-urban-dictionary.p.rapidapi.com/define?term="+payload, strings.NewReader(""))
		if err != nil {
			return
		}
		req.Header.Set("X-RapidAPI-Host", "mashape-community-urban-dictionary.p.rapidapi.com")
		req.Header.Set("X-RapidAPI-Key",  "c10834fe35mshdc5083d2e82abf8p124944jsnd50094792e4c")

		client := &http.Client{}
		res, err := client.Do(req)


		body, err := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))

		var d Definitions

		_ = json.Unmarshal(body, &d)
		fmt.Println(d.Elements[0].Info)
		b.Send(m.Chat, d.Elements[0].Info)
		b.Send(m.Chat, d.Elements[1].Info)
		b.Send(m.Chat, d.Elements[2].Info)
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
		fmt.Println(ko.Text)
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