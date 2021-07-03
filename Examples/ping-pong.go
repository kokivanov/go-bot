package main

import (
	"fmt"
	"time"

	dbl "github.com/kokivanov/go-bot/DiscordBotLib"
)

const prefix = "a-"

func OnMessage(c *dbl.Client, m dbl.Message) {
	if c.IsMyID(m.Author.ID) {
		return
	}

	if m.Author.Bot != nil {
		if *m.Author.Bot {
			return
		}
	}

	fmt.Printf("%v: Got message from %s: %s\n", m.Timestamp.ToTimeObject().UTC(), m.Author.Username, m.Content)

	g := fmt.Sprintf("PONG! Ping is *%vms*! 🏓", (time.Now().UnixNano()-m.Timestamp.ToTimeObject().UnixNano())/1000000)
	if m.Content == prefix+"ping" {
		m.Reply(&g, false, nil, nil, nil, nil)
	}
}

func OnReady(c *dbl.Client, r dbl.Ready) {
	fmt.Printf("Successfully registered as %v#%v", r.User.Username, r.User.Discriminator)
}

func main() {

	var client = dbl.NewClient()

	client.AddHandler(OnMessage, dbl.EventMessageCreate)
	client.AddHandler(OnReady, dbl.EventReady)
	fmt.Println(client.GetAvialableHandlers())

	client.Run("YOUR_TOKEN_HERE", nil)
}
