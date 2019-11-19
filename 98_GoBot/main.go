package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
	"unicode/utf8"

	tb "gopkg.in/tucnak/telebot.v2"
)

type Definitions struct {
	Elements []Definition `json:"list"`
}

type Definition struct {
	Info string `json:"definition"`
}



func main() {

	var todos []string


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


	//################## insert quote ##############################################################################

	b.Handle("/quote", func(m *tb.Message) {
		maxCharsPerLine := 35
		stringArray := strings.Split(m.Payload, " ")
		longestString := 0

		var in []int
		q := make(map[string]string)

		for i, s := range stringArray {
			if strings.Contains(s, ":") {
				in = append(in, i)
				if utf8.RuneCountInString(s) > longestString {
					longestString = utf8.RuneCountInString(s)
				}
			}
		}


		for i := 0; i < len(in); i++ {
			if i < len(in)-1 {
				currentName := in[i]
				nextName := in[i+1]
				q[stringArray[currentName]] = strings.Join(stringArray[currentName+1:nextName], " ")

			} else {
				currentName := in[i]
				q[stringArray[currentName]] = strings.Join(stringArray[currentName+1:], " ")
			}
		}

		outputString := ""



		for i, s := range q {
			outputString = outputString + i + "\n" + formatString(s, maxCharsPerLine, longestString) + "\n\n"
		}

		fmt.Println(outputString)

		// printLKT(outputString)

	})



	//################## add to shop ######################################################################################

	b.Handle("/shop", func(m *tb.Message) {
		temp := strings.Split(m.Payload, " ")
		for _, v := range temp {
			todos = append(todos, "- " + v)
		}

	})


	//################## print shopping list ##############################################################################

	b.Handle("/print", func(m *tb.Message) {
		l := "Bitte bring folgendes mit:\n_________________\n\n" + strings.Join(todos[:], "\n")
		note := "\n\n\n.~~~~.\ni====i_\n|cccc|_)\n|cccc|\n`-==-'\n\nund bei Bedarf etwas Bier"
		printLKT(l + note)


	})


	//################## look up urban dict ################################################################################

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


func printLKT(t string) {
	printText := []byte(t)
	err := ioutil.WriteFile("./test.txt", printText, 0644)
	if err != nil {
		log.Fatal(err)
	}

	sh := `lp test.txt -d LKT`
	// sh := `echo "Hello World"`
	args := strings.Split(sh, " ")

	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s \n", b)
}

func formatString(text string, maxLen, longest int) string {
	output := ""
	col := 0
	max := maxLen - longest
	tempText := strings.Split(text, " ")
	// fmt.Println(tempText)

	for _, v := range tempText {
		if col == 0 {
			output = output + strings.Repeat(" ", longest)
			col += longest
		} else if (col + utf8.RuneCountInString(v)) < max {
			output = output + v + " "
			col += utf8.RuneCountInString(v) + 1
		} else {
			output = output + "\n" + strings.Repeat(" ", longest) + v + " "
			col = longest + utf8.RuneCountInString(v) + 1
		}
	}
	// fmt.Println(output)

	return output


	//output = output + strings.Repeat(" ", longest)
}