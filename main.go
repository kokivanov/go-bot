package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	dbl "github.com/kokivanov/go-bot/DiscordBotLib"
)

func OnMessage(c *dbl.Client, m dbl.Message) {
	fmt.Printf("Got message from %s: %s\n", m.Author.Username, m.Content)
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
