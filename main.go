package main

import (
	"fmt"
	"sync"
	"time"

	dbl "github.com/kokivanov/go-bot/DiscordBotLib"
)

type Test struct {
	Wg sync.WaitGroup
	CC chan int
}

func (t *Test) listen(ch <-chan int) {
	for {

		time.Sleep(time.Millisecond * 2000)
		fmt.Println("Listening")
		select {
		default:
		case <-ch:
			fmt.Println("Listening ended")
			t.Wg.Done()
			return
		}
	}
}
func (t *Test) Run() {
	t.CC = make(chan int, 2)

	t.Wg.Add(2)
	go t.heartbeat(t.CC)
	go t.listen(t.CC)
	t.Wg.Wait()
}

func (t *Test) Stop() {
	fmt.Println("Stopping")
	t.CC <- 1
	t.CC <- 1
}

func (t *Test) heartbeat(ch <-chan int) {

	tk := time.NewTicker(time.Millisecond * 2500)

	for {

		select {
		case <-tk.C:
			fmt.Println("Heartbeating")
		case <-ch:
			fmt.Println("Heartbeating ended")
			t.Wg.Done()
			return
		}
	}
}

func main() {

	// t := Test{}

	// go func() {
	// 	time.Sleep(time.Second * 20)
	// 	t.Stop()
	// }()

	// t.Run()

	var client = dbl.NewClient()

	client.LogLevel = 1024

	wg := sync.WaitGroup{}
	wg.Add(1)
	client.SetIntent(512)

	go func(c *dbl.Client) {
		time.Sleep(time.Second * 20)
		c.Stop()
		wg.Done()
	}(client)

	client.Run("NzAzOTYxNjE3NDM1NDU5NjE0.XqWNWA.kMLcsDyQ-FG9g71LXakp-GvGFto")

	wg.Wait()

	// fmt.Println("Hello bot!")

	// header := http.Header{}
	// header.Add("accept-encoding", "zlib")

	// conn, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?v=9&encoding=json", header)
	// if err != nil {
	// 	println("Error uccured!", err.Error())
	// }

	// mt, m, err := conn.ReadMessage()

	// var reader io.Reader = bytes.NewBuffer(m)

	// if mt == websocket.BinaryMessage {
	// 	fmt.Println("It's a binary message!")
	// }

	// var e *DiscordBotLib.Payload
	// decoder := json.NewDecoder(reader)
	// if err = decoder.Decode(&e); err != nil {
	// 	fmt.Printf("error decoding websocket message, %s \n", err)
	// }

	// fmt.Println(e.Operation)
}
