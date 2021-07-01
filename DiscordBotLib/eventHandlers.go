package DiscordBotLib

type OnMessage func(*Client, Message)
type OnReady func(*Client)

func (om OnMessage) Handle(c *Client, i interface{}) {
	om(c, i.(Message))
}
func (om OnMessage) Type() string {
	return EventMessageCreate
}

func (or OnReady) Handle(c *Client, i interface{}) { // TODO: Change interface{} to ReadyPayload or Application and pass it to function
	or(c)
}

func (or OnReady) Type() string {
	return "READY"
}
