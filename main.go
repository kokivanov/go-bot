package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	dbl "github.com/kokivanov/go-bot/DiscordBotLib"
)

func OnMessage(c *dbl.Client, m dbl.Message) {
	fmt.Print("\n\n\nGOT MESSAGE!!!!!!\n\n\n")
}

func OnReady(c *dbl.Client) {
	return
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

	client.SetIntent(512)

	client.AddHandler(OnMessage)
	client.AddHandler(OnReady)
	fmt.Println(client.GetAvialableHandlers())

	client.Run(fmt.Sprintf("%s", conf["token"]))

}
