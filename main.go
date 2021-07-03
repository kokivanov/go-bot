package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	dbl "github.com/kokivanov/go-bot/DiscordBotLib"
)

const prefix = "a-"

func OnMessage(c *dbl.Client, m dbl.Message) {
	fmt.Printf("%v: Got message from %s: %s\n", m.Timestamp.ToTimeObject().UTC(), m.Author.Username, m.Content)

	s := 2
	g := "PONG! Have a button!"
	if m.Content == prefix+"ping" {
		m.Reply(&g, false, nil, nil, nil, &[]dbl.Component{{
			Type: 1,
			Components: &[]dbl.Component{{
				Type:     2,
				Label:    "CLICK ME!!!!",
				Style:    &s,
				CustomID: "kokibutton",
			}},
		}})
	}
}

func OnReady(c *dbl.Client, r dbl.Ready) {

}

func main() {

	jsonFile, err := os.Open("settings.json")
	if err != nil {
		println(err)
	}

	conf := make(map[string]interface{})

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &conf)

	var client = dbl.NewClient()

	client.LogLevel = 1024

	client.AddHandler(OnMessage, dbl.EventMessageCreate)
	client.AddHandler(OnReady, dbl.EventReady)
	fmt.Println(client.GetAvialableHandlers())

	client.Run(fmt.Sprintf("%s", conf["token"]), nil)
}
