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
