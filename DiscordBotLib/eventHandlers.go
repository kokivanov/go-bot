package DiscordBotLib

type OnMessage func(*Client, Message)

func (om OnMessage) Handle(c *Client, i interface{}) {
	om(c, i.(Message))
}
func (om OnMessage) Type() string {
	return EventMessageCreate
}

type OnReady func(*Client, Ready)

func (or OnReady) Handle(c *Client, i interface{}) { // TODO: Change interface{} to ReadyPayload or Application and pass it to function
	or(c, i.(Ready))
}

func (or OnReady) Type() string {
	return EventReady
}

type OnResumed func(*Client)

func (or OnResumed) Handle(c *Client, i interface{}) {
	or(c)
}

func (or OnResumed) Type() string {
	return EventResumed
}

type OnSlashCommandCreate func(*Client, ApplicationCommand)

func (or OnSlashCommandCreate) Handle(c *Client, i interface{}) {
	or(c, i.(ApplicationCommand))
}

func (or OnSlashCommandCreate) Type() string {
	return EventApplicationCommandCreate
}

type OnSlashCommandUpdate func(*Client, ApplicationCommand)

func (or OnSlashCommandUpdate) Handle(c *Client, i interface{}) {
	or(c, i.(ApplicationCommand))
}

func (or OnSlashCommandUpdate) Type() string {
	return EventApplicationCommandUpdate
}

type OnSlashCommandDelete func(*Client, ApplicationCommand)

func (or OnSlashCommandDelete) Handle(c *Client, i interface{}) {
	or(c, i.(ApplicationCommand))
}

func (or OnSlashCommandDelete) Type() string {
	return EventApplicationCommandDelete
}

type OnChannelCreate func(*Client, Channel)

func (or OnChannelCreate) Handle(c *Client, i interface{}) {
	or(c, i.(Channel))
}

func (or OnChannelCreate) Type() string {
	return EventChannelCreate
}

type OnChannelUpdate func(*Client, Channel)

func (or OnChannelUpdate) Handle(c *Client, i interface{}) {
	or(c, i.(Channel))
}

func (or OnChannelUpdate) Type() string {
	return EventChannelUpdate
}

type OnChannelDelete func(*Client, Channel)

func (or OnChannelDelete) Handle(c *Client, i interface{}) {
	or(c, i.(Channel))
}

func (or OnChannelDelete) Type() string {
	return EventChannelDelete
}

type OnChannelPinsUpdate func(*Client, ChannelPinsUpdate)

func (or OnChannelPinsUpdate) Handle(c *Client, i interface{}) {
	or(c, i.(ChannelPinsUpdate))
}

func (or OnChannelPinsUpdate) Type() string {
	return EventChannelPinsUpdate
}

type OnThreadCreate func(*Client, Channel)

func (or OnThreadCreate) Handle(c *Client, i interface{}) {
	or(c, i.(Channel))
}

func (or OnThreadCreate) Type() string {
	return EventThreadCreate
}

type OnThreadUpdate func(*Client, Channel)

func (or OnThreadUpdate) Handle(c *Client, i interface{}) {
	or(c, i.(Channel))
}

func (or OnThreadUpdate) Type() string {
	return EventThreadUpdate
}

type OnThreadDelete func(*Client, Channel)

func (or OnThreadDelete) Handle(c *Client, i interface{}) {
	or(c, i.(Channel))
}

func (or OnThreadDelete) Type() string {
	return EventThreadDelete
}

type OnThreadListSync func(*Client, ThreadListSync)

func (or OnThreadListSync) Handle(c *Client, i interface{}) {
	or(c, i.(ThreadListSync))
}

func (or OnThreadListSync) Type() string {
	return EventThreadListSync
}
