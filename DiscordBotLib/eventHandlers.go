package DiscordBotLib

type interfaceEventHandler func(*Client, interface{})
type OnMessage func(*Client, Message)
type OnReady func(*Client)

func (om OnMessage) Handle(c *Client, i interface{}) {
	om(c, i.(Message))
}
func (om OnMessage) Type() string {
	return "MESSAGE_CREATE"
}

func (or OnReady) Handle(c *Client, i interface{}) {
	or(c)
}

func (or OnReady) Type() string {
	return "READY"
}

func (eh interfaceEventHandler) Handle(s *Client, i interface{}) {
	eh(s, i)
}

func (eh interfaceEventHandler) Type() string {
	return "__INTERFACE__"
}
