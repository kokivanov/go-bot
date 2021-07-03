package DiscordBotLib

/*
----------------------------------------
  				Client
----------------------------------------
Client represents user-friendly struct that allows yuo to work with discord api as a bot.
Here you can fined all functions that are related to Discord Bot.
*/

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
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

type Identify struct {
	Token      string             `json:"token"`
	Properties IdentifyProperties `json:"properties"`
	Intents    int                `json:"intents"`

	Compress       bool            `json:"compress,omitempty"`
	LargeThreshold int             `json:"large_threshold,omitempty"`
	Shard          *[2]int         `json:"shard,omitempty"`
	Presence       *UpdatePresence `json:"presence,omitempty"`
}

type IdentifyProperties struct {
	Os      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

type UpdatePresence struct {
	Since      *int              `json:"since"`
	Activities *[]ActivityObject `json:"activities"`
	Status     string            `json:"string"`
	AFK        bool              `json:"afk"`
}

type Client struct { // TODO: omitempty
	// Mutex
	sync.RWMutex

	wsMutex   sync.Mutex
	HTTPMutex sync.Mutex
	wG        sync.WaitGroup

	// Struncts that are used to work with api
	wsGateway  string
	wsConn     *websocket.Conn
	httpClient *http.Client
	authHeader *http.Header

	// Discord related fields
	intent            int
	token             string
	heartbeatInterval int
	lastSequence      int64
	lastHeartbeatACK  uint64
	SessionID         string

	handlers map[string]*EventHandler

	// information about application and it's owner
	Owner User
	MeApp Application
	Me    User

	LogLevel             int
	state                uint8
	interrupt            chan int
	ReconnectMaxAttempts uint16
	currentAttempt       uint16

	// TODO: Task Queqe
}

/* Just a function that will terminate bot on keyboard interrup signal (^C) */
func (c *Client) setupCloseHandler() {
	cc := make(chan os.Signal, 3)
	signal.Notify(cc, os.Interrupt, syscall.SIGTERM)
	<-cc
	c.Stop(3)
	if c.state == 2 {
		c.wG.Done()
		c.wG.Done()
	}
	fmt.Println("\r- Ctrl+C pressed in Terminal")
}

/* Runs your bot with provided token
Initializes all required veriables, spawns goroutines listen and heratbeat and prevents app from early termination*/
func (c *Client) Run(token string, wg *sync.WaitGroup) error {

	if c.state == 1 || c.wsConn != nil {
		return errors.New("already running")
	}

	err := c.Resume()

	if err != nil && c.LogLevel >= LogWarnings {
		log.Printf("Can't resume: %s", err.Error())

		c.interrupt = make(chan int)

		c.wG.Add(2)
		//go c.setupCloseHandler()
		if err := c.init(token); err != nil {
			return err
		}
		go c.heartbeat(c.heartbeatInterval, c.interrupt)
		go c.listen(c.interrupt)
		c.wG.Wait()

		if wg != nil {
			wg.Done()
		}
	}

	return nil
}

/* NewClient() returns pointer to new Client struct with all needed initialized fields
It's recommended to use this function to get client struct with default parameters */
func NewClient() *Client {
	c := &Client{
		httpClient: &http.Client{},
		authHeader: &http.Header{},
	}
	return c
}

/* Adds function that will be called on specific gateway event to the functions stack.
Enables itents depending on function type (see docs or eventHandlers.go) */
func (c *Client) AddHandler(handler interface{}, Type string) {
	if c.handlers == nil {
		c.handlers = make(map[string]*EventHandler)
	}

	he := getEventHandler(handler, Type) // TODO: Add intents depending on handler type
	if he != nil {
		c.SetIntent(getEventIntent(he.Type()))
		c.handlers[he.Type()] = &he
	}
}

/*Returns all successfully registeret handlers*/
func (c *Client) GetAvialableHandlers() []string { // TODO: Handler overwrite check
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

	// Creating HTTP GET request to Discord API in order to get working gateway url
	req, err := http.NewRequest("GET", fmt.Sprintf(APIURL+"/v%d"+GetGatewayEndpoint, APIVersion), nil)
	if err != nil {
		log.Print(err)
		return err
	}

	// Making request that was createt above
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Print(err)
		return err
	}

	// Reading request body (it must be json)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return err
	}

	// Unmarshaling
	var body map[string]interface{}
	err = json.Unmarshal(b, &body)
	if err != nil {
		log.Print(err)
		return err
	}

	// Getting URL
	if tmp, ok := body["url"]; ok {
		c.wsGateway = fmt.Sprintf("%v", tmp)
	} else {
		return errors.New("no url provided")
	}

	return nil
}

/* connect() initializes field Client.wsConn where websocket ("https://github.com/gorilla/websocket") connection object will be located*/
func (c *Client) connect() error {
	header := http.Header{}
	header.Add("accept-encoding", "zlib")

	/* Oppening gateway and connecting to it */
	var err error
	if c.LogLevel >= LogAll {
		log.Printf("Creating dialer")
	}
	c.wsConn, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("%s?v=%d&encoding=%s", c.wsGateway, GatewayVersion, GatewayEncoding), header)
	if err != nil {
		return err
	}

	c.wsConn.SetCloseHandler(func(code int, text string) error {
		log.Println("Closing gateway")
		return nil
	})

	/* Reading first message, it must be Opcode 10 Hello, see https://discord.com/developers/docs/topics/gateway#connecting-to-the-gateway for more info*/
	if c.LogLevel >= LogAll {
		log.Printf("reading opcode hello")
	}

	mt, m, err := c.wsConn.ReadMessage()
	if err != nil {
		return err
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
		return err
	}

	if c.LogLevel >= 8 {
		log.Printf("Got payload:\n%+v\n", e)
	}

	if *e.Operation != GatewayOpHello {
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

/*As discribed at discord developers portal
https://discord.com/developers/docs/topics/gateway#heartbeat
*/
func (c *Client) heartbeat(interval int, ls <-chan int) { // TODO: Make zombied connection check!

	tk := time.NewTicker(time.Duration(interval) * time.Millisecond)
	for {

		select {

		case <-tk.C:

			if c.LogLevel >= LogAll {
				fmt.Println("Still heartbeating!")
			}

			if c.state != 1 || c.wsConn == nil {
				if c.state == 3 {
					if c.state == 3 {
						c.wG.Done()
					}
					if c.LogLevel >= LogWarnings {
						log.Printf("Heartbeat: State: %v, calling wg.Done()", c.state)
					}
					c.wG.Done()
				}
				return
			}

			if c.LogLevel >= LogAll {
				log.Printf("Sending hertbeat sequnce %d", c.lastSequence)
			}

			c.wsMutex.Lock()
			err := c.wsConn.WriteJSON(Heartbeat{Op: 1, D: c.lastSequence})
			c.wsMutex.Unlock()

			if err != nil {
				fmt.Printf("error while sending hertbaet: %v", err.Error())
				return
			}

		case <-ls:
			log.Println("Called interrupt, heartbeating terminated.")
			if c.state == 3 {
				if c.LogLevel >= LogWarnings {
					log.Printf("Heartbeat: State: %v, calling wg.Done()", c.state)
				}
				c.wG.Done()
			}
			return
		}
	}
}

/* Always waits for message to recieve
On message spawns goroutine handleEvent()*/
func (c *Client) listen(ls <-chan int) {
	for {

		if c.LogLevel >= LogAll {
			fmt.Println("Still listening!")
		}

		if c.wsConn == nil {
			log.Println("Websocket connection terminated")
			if c.state == 3 {
				c.wG.Done()
				return
			}

			if c.state == 2 {
				return
			}
			return
		}

		if c.state != 1 {
			log.Println("Client isn't in working state")
			if c.state == 3 {
				if c.LogLevel >= LogWarnings {
					log.Printf("Listen: State: %v, calling wg.Done()", c.state)
				}
				c.wG.Done()
				return
			}

			if c.state == 2 {
				return
			}
			return
		}

		mt, m, err := c.wsConn.ReadMessage()
		if err != nil {

			log.Printf("App in %v state", c.state)

			if c.state == 1 {
				c.Resume()
				return
			}

			log.Printf("Error occured while listening to message!\n %v", err.Error())
			if c.state == 3 {
				if c.LogLevel >= LogWarnings {
					log.Printf("Listen: State: %v, calling wg.Done()", c.state)
				}
				c.wG.Done()
				return
			}

			if c.state == 2 {
				return
			}

			if strings.Contains(err.Error(), "1000") {
				log.Println("Listening to closed websocket connection")
				return
			}
		}
		if mt == websocket.BinaryMessage {
			fmt.Println("I don't know how to handle binary message!")
			return
		}

		event := Payload{}

		json.Unmarshal(m, &event)
		if c.LogLevel >= LogAll {
			log.Printf("Got %v event %v with sequence %v: %s", *event.Operation, event.Type, event.Sequence, string(event.RawData))
		}

		c.Lock()
		c.lastSequence = event.Sequence
		c.Unlock()

		select {
		case <-ls:
			log.Println("Called interrupt, listening terminated. Number of goroutines running: ", runtime.NumGoroutine())
			if c.state == 3 {
				if c.LogLevel >= LogWarnings {
					log.Printf("Listen: State: %v, calling wg.Done()", c.state)
				}
				c.wG.Done()
				return
			}

			if c.state == 2 {
				return
			}
			return

		default:
			go c.handleEvent(event)
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
		if c.state != 3 {
			c.state = 2
		}
		return err
	}

	return nil
}

// Will be called when listiner gets event
func (c *Client) handleEvent(payload Payload) {
	if payload.Operation != nil {

		switch *payload.Operation {
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

}

type resume struct {
	Op   int `json:"op"`
	Data struct {
		Token     string `json:"token"`
		SessionID string `json:"session_id"`
		Sequence  int64  `json:"seq"`
	} `json:"d"`
}

func (c *Client) Resume() error {

	c.Stop(2)

	c.currentAttempt = 0

	if c.state != 3 {
		c.state = 2
	}

	if c.lastSequence == 0 {
		log.Printf("can't resume, last sequence %v", c.lastSequence)
		err := fmt.Errorf("can't resume, last sequence %v", c.lastSequence)
		return err
	}

	for {

		err := c.getGateway()
		if err != nil {
			log.Printf("Error occured while retrying: %v", err)
		}

		err = c.connect()
		if err != nil {
			log.Printf("Error occured while retrying: %v", err)
		} else {
			c.Lock()
			res := resume{}
			res.Op = GatewayOpResume
			res.Data.Sequence = int64(c.lastSequence)
			res.Data.SessionID = c.SessionID
			res.Data.Token = c.token

			c.wsMutex.Lock()
			log.Printf("Resuming with struct: \n%#v\n", res)
			err = c.wsConn.WriteJSON(res)
			c.wsMutex.Unlock()

			if err != nil {
				log.Printf("Error occured while resumin, attempt %v", c.currentAttempt)
			}

			if c.ReconnectMaxAttempts <= c.currentAttempt {
				c.Unlock()
				c.wG.Done()
				c.wG.Done()
				return fmt.Errorf("hit max reconnect attempts")
			}

			c.currentAttempt++
			_, m, err := c.wsConn.ReadMessage()

			if len(m) <= 10 {
				log.Fatalf("It didn't work")
			}

			if err != nil {
				log.Printf("error while reading message: %v", err)
			}

			event := Payload{}
			json.Unmarshal(m, &event)

			log.Printf("Got payload: %+v", event)

			if event.Operation != nil {

				if *event.Operation == 9 {
					c.Unlock()
					res, _ := strconv.ParseBool(string(event.RawData))
					log.Printf("Can't resume, payload \"d\": %v", res)

					err := c.identify()
					if err == nil {
						c.state = 1

						go c.listen(c.interrupt)
						go c.heartbeat(c.heartbeatInterval, c.interrupt)

						go c.handleEvent(event)

						return nil
					}
				}

				if *event.Operation == 0 {
					c.Unlock()
					log.Println("Successfully resumed.")

					c.state = 1

					go c.listen(c.interrupt)
					go c.heartbeat(c.heartbeatInterval, c.interrupt)

					go c.handleEvent(event)

					return nil
				}

			}

			c.Unlock()

		}

		wait := time.Duration(time.Second * 5)
		log.Println("Retryig in ", wait)
		time.Sleep(wait)
		wait *= 2
		if wait > 300 {
			wait = 300
		}
	}

}

func (c *Client) UpdatePresence() {

}

func (c *Client) Stop(code int) error {

	if c.state == 0 || c.state == 2 || c.state == 3 {
		return errors.New("application already stopped")
	}

	if c.LogLevel >= 4 {
		log.Println("Commiting suicide!")
	}

	if code == 3 {
		c.state = 3
	} else {
		c.state = 2
	}

	c.interrupt <- 4

	err := c.wsConn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(1000, ""), time.Now().Add(time.Second*10))
	if err != nil {
		if c.LogLevel >= 4 {
			c.wsConn.Close()
			log.Println("Error while closing connection", err)
		}
		return err
	}

	_, m, _ := c.wsConn.ReadMessage()
	if c.LogLevel >= LogWarnings {
		log.Println(string(m))
	}

	if err := c.wsConn.Close(); err != nil {
		if c.LogLevel >= 4 {
			log.Println("Error while closing connection", err)
		}
		return err
	}

	c.wsConn = nil

	if c.LogLevel >= 4 {
		log.Printf("Stop function ended!\nWSConn: %v\nState: %v", c.wsConn, c.state)
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

	if c.ReconnectMaxAttempts == 0 {
		c.ReconnectMaxAttempts = 10
	}

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

// TODO: func register slash commnad
// TODO: func get guild
// TODO: func get guilds
// TODO: func get user
// TODO: func get DMChannel
// TODO: func IsMe(u *User) bool
