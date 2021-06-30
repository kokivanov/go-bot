package DiscordBotLib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

// ================================ Utility section =====================================

// ------------------------------------------
//                    Utility
// ------------------------------------------
// Contains all utility functions that are will be used during event handling, parsing, making requests, etc.
//

func getEventHandler(h interface{}) EventHandler { // TODO: Complete
	switch v := h.(type) {
	case func(*Client):
		return OnReady(v)
	case func(*Client, Message):
		return OnMessage(v)
	case func(*Client, interface{}):
		return interfaceEventHandler(v)
	default:
		return nil
	}
}

func (c *Client) getEventPayload(p Payload) interface{} { // TODO: Complete
	switch p.Type {
	case "MESSAGE_CREATE":
		m := Message{
			ClientPTR:   c,
			Author:      User{ClientPTR: c},
			GuildMember: &GuildMember{ClientPTR: c},
		}
		json.Unmarshal(p.RawData, &m)
		return m
	default:
		return nil
	}
}

// ============================= Client section =========================================

// ----------------------------------------
//   				Client
// ----------------------------------------
// Client represents user-friendly struct that allows yuo to work with discord api as a bot.
// Here you can fined all functions that are related to Discord Bot.
//
//
//
//

func (c *Client) setupCloseHandler() {
	cc := make(chan os.Signal, 3)
	signal.Notify(cc, os.Interrupt, syscall.SIGTERM)
	<-cc
	fmt.Println("\r- Ctrl+C pressed in Terminal")
	c.Stop()
}

func (c *Client) Run(token string) error {
	c.interrupt = make(chan int, 2)

	c.wG.Add(2)
	//go c.setupCloseHandler()
	if err := c.init(token); err != nil {
		return err
	}
	go c.heartbeat(c.heartbeatInterval, c.interrupt)
	go c.listen(c.interrupt)
	c.wG.Wait()
	return nil
}

// NewClient() returns pointer to new Client struct with all needed initialized fields
// It's recommended to use this function to get client struct with default parameters
func NewClient() *Client {
	c := &Client{
		httpClient: &http.Client{},
		authHeader: &http.Header{},
	}
	return c
}

// Adds function that will be called on specific gateway event to the functions stack.
// Enables itents depending on function type (see docs or types.go)
func (c *Client) AddHandler(handler interface{}) {
	if c.handlers == nil {
		c.handlers = make(map[string]*EventHandler)
	}

	he := getEventHandler(handler)
	if he != nil {
		c.handlers[he.Type()] = &he
	}
}

func (c *Client) GetAvialableHandlers() []string {
	res := make([]string, 0)
	for k := range c.handlers {
		res = append(res, k)
	}

	return res
}

func (c *Client) SetIntent(intent int) error {
	c.intent = c.intent | intent
	return nil
}

func (c *Client) GetIntent() int {
	return c.intent
}

// GetGateway() returns a valid discord Websocker API gateway
func (c *Client) getGateway() error {

	req, err := http.NewRequest("GET", fmt.Sprintf(APIURL+"/v%d"+GetGatewayEndpoint, APIVersion), nil)
	if err != nil {
		log.Print(err)
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Print(err)
		return err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return err
	}

	var body map[string]interface{}
	err = json.Unmarshal(b, &body)
	if err != nil {
		log.Print(err)
		return err
	}

	if tmp, ok := body["url"]; ok {
		c.wsGateway = fmt.Sprintf("%v", tmp)
	} else {
		return errors.New("no url provided")
	}

	return nil
}

// connect() initializes field Client.wsConn where websocket ("github.com/gorilla/websocket") connection object will be located
func (c *Client) connect() error {
	header := http.Header{}
	header.Add("accept-encoding", "zlib")

	var err error
	c.wsConn, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s?v=%d&encoding=%s", c.wsGateway, GatewayVersion, GatewayEncoding), header)
	if err != nil {
		log.Fatal("Error occured while connecting to websocket!\n\n", err)
	}

	mt, m, err := c.wsConn.ReadMessage()
	if err != nil {
		log.Fatal("Error occured while reading first message!\n\n", err)
	}

	if mt == websocket.BinaryMessage {
		fmt.Println("It's a binary message!")
		return errors.New("can't decode binary message") // TODO: Make it!
	}

	var reader io.Reader = bytes.NewBuffer(m)

	var e *Payload
	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&e); err != nil {
		fmt.Printf("error decoding websocket message, %s \n", err)
	}

	if c.LogLevel >= 8 {
		log.Printf("Got payload:\n%+v\n", e)
	}

	if e.Operation != GatewayOpHello {
		log.Fatalf("Expected Opcode 10 Hello, but got %d", e.Operation)
	}

	var tmp map[string]interface{}
	err = json.Unmarshal(e.RawData, &tmp)
	if err != nil {
		return err
	}

	if c.LogLevel >= LogAll {
		log.Printf("Got raw data: %v\n", tmp)
	}

	if heartbeatInterval, ok := tmp["heartbeat_interval"]; ok {
		switch tp := heartbeatInterval.(type) {
		case float64:
			if tmp, ok := heartbeatInterval.(float64); ok {
				c.heartbeatInterval = int(tmp)
			} else {
				return fmt.Errorf("can not get heatrbeat interval (not float64 or int), got %v", tp)
			}
		default:
			if c.heartbeatInterval, ok = heartbeatInterval.(int); !ok {
				return fmt.Errorf("can not get heatrbeat interval (not float64 or int), got %v", tp)
			}
		}
	} else {
		return errors.New("not a hello payload")
	}

	return nil
}

// Returns urrent state
func (c *Client) GetState() uint8 {
	return c.state
}

// As discribed at discord developers portal
// https://discord.com/developers/docs/topics/gateway#heartbeat
func (c *Client) heartbeat(interval int, ls <-chan int) { // TODO: Make zombied connection check!

	tk := time.NewTicker(time.Duration(interval) * time.Millisecond)
	for {

		select {

		case <-tk.C:

			if c.LogLevel >= LogAll {
				fmt.Println("Still heartbeating!")
			}

			if c.wsConn == nil {
				c.wG.Done()
				return
			}

			if c.LogLevel >= LogAll {
				log.Printf("Sending hertbeat sequnce %d", c.lastSequence)
			}

			if c.state != 1 {
				c.wG.Done()
				return
			}

			c.wsMutex.Lock()
			err := c.wsConn.WriteJSON(Heartbeat{Op: 1, D: &c.lastSequence})
			c.wsMutex.Unlock()

			if err != nil {
				fmt.Printf("error while sending hertbaet: %v", err.Error())
				return
			}

		case <-ls:
			log.Println("Called interrupt, heartbeating terminated. Number of goroutines running: ", runtime.NumGoroutine())
			c.wG.Done()
			return
		}
	}
}

func (c *Client) listen(ls <-chan int) {
	for {

		if c.LogLevel >= LogAll {
			fmt.Println("Still listening!")
		}

		if c.wsConn == nil {
			log.Println("Websocketconnection termanated")
			c.wG.Done()
			return
		}

		if c.state != 1 {
			log.Println("Client isn't in working state")
			c.wG.Done()
			return
		}

		mt, m, err := c.wsConn.ReadMessage()
		if err != nil {
			log.Printf("Error occured while listening to message!\n %v", err.Error())
			c.wG.Done()
			return
		}
		if mt == websocket.BinaryMessage {
			fmt.Println("I don't know how to handle binary message!")
			return
		}

		event := Payload{}
		c.lastSequence = event.Sequence
		json.Unmarshal(m, &event)
		if c.LogLevel >= LogAll {
			log.Printf("Got %v event %v with sequence %v: %s", event.Operation, event.Type, event.Sequence, string(event.RawData))
		}

		select {
		default:
			c.handleEvent(event)
		case <-ls:
			log.Println("Called interrupt, listening terminated. Number of goroutines running: ", runtime.NumGoroutine())
			c.wG.Done()
			return
		}
	}
}

// Unexported function that used to sent identify peayload to gateway
func (c *Client) identify() error {
	// TODO: Check if resuming is needed
	identifyProp := IdentifyProperties{
		Os:      runtime.GOOS,
		Browser: LibName,
		Device:  LibName,
	}

	identifyPayload := Identify{
		Token:      c.token,
		Properties: identifyProp,
		Intents:    c.intent,
	}

	var ident = struct {
		Operation int      `json:"op"`
		Data      Identify `json:"d"`
	}{
		Operation: 2,
		Data:      identifyPayload,
	}

	c.wsMutex.Lock()
	log.Printf("Identifying with struct: \n%#v\n", ident)
	err := c.wsConn.WriteJSON(ident)
	c.wsMutex.Unlock()

	if err != nil {
		c.state = 2
		return err
	}

	return nil
}

// Will be called when listiner get
func (c *Client) handleEvent(payload Payload) {
	switch payload.Operation {
	case GatewayOpHeartbeatACK:
		c.Lock()
		c.lastHeartbeatACK = uint64(time.Now().Unix())
		c.Unlock()
	case GatewayOpDispatch:
		if ev, ok := c.handlers[payload.Type]; !ok && c.LogLevel >= LogWarnings {
			log.Printf("Can't handle event %s.\n", payload.Type)
		} else {
			if c.LogLevel >= LogMessages {
				log.Printf("Handling event %s.\n", payload.Type)
			}
			(*ev).Handle(c, c.getEventPayload(payload))
		}
	}

}

func (c *Client) resume() {

}

func (c *Client) Reconnect() {

}

func (c *Client) updatePresence() {

}

func (c *Client) Stop() error {

	if c.LogLevel >= 4 {
		log.Println("Commiting suicide!")
	}

	c.interrupt <- 4
	c.interrupt <- 4

	c.state = 2

	err := c.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, ""))
	if err != nil {
		return err
	}

	if err := c.wsConn.Close(); err != nil {
		return err
	}

	c.wsConn = nil

	if c.LogLevel >= 4 {
		log.Println("Stop function ended!")
	}
	return nil
}

func (c *Client) init(token string) error {
	if err := c.getGateway(); err != nil {
		return errors.New("error getting gateway")
	}

	go c.setupCloseHandler()
	c.token = token

	c.Lock()
	c.state = 1

	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	if c.authHeader == nil {
		c.authHeader = &http.Header{}
	}

	if c.handlers == nil {
		c.handlers = make(map[string]*EventHandler)
	}

	c.authHeader.Add("Authorization", ("Bot " + token))
	c.authHeader.Add("User-Agent", "DiscordBot (\"https://github.com/kokivanov/go-bot/DiscordBotLib\", DoscordBotLib)")
	err := c.connect()
	if err != nil {
		return err
	}
	c.identify()
	c.Unlock()

	return nil
}
