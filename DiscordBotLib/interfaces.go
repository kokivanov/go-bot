package DiscordBotLib

type EventHandler interface {
	Handle(*Client, interface{})
	Type() string
}

type ClientMember interface {
	GetParent() *Client
}
