package DiscordBotLib

type Message struct {
	ID                  Snowflake             `json:"ID"`
	ChannelID           int                   `json:"channel_ID"`
	GuildID             *int                  `json:"guild_ID,omitempty"`
	Author              User                  `json:"author"`
	GuildMember         *GuildMember          `json:"member,omitempty"`
	Content             string                `json:"content"`
	Timestamp           ISO8601Timestamp      `json:"timestamp"`
	EditedTimestamp     *ISO8601Timestamp     `json:"edited_timestamp"`
	TTS                 bool                  `json:"tts"`
	MentionEveryone     bool                  `json:"mention_everyone"`
	Mentions            []User                `json:"mentions"`
	MentionRoles        []int                 `json:"mention_roles"`
	MentionChannels     *[]ChannelMention     `json:"mention_channels,omitempty"`
	Attachments         []Attachment          `json:"attachments"`
	Embeds              []Embed               `json:"embeds"`
	Reactions           *[]Reaction           `json:"reactions,omitempty"`
	Nonce               string                `json:"nonce,omitempty"`
	Pinned              bool                  `json:"pinned"`
	WebhookID           *int                  `json:"webhook_ID,omitempty"`
	Type                int                   `json:"type"` // TODO: Make constants for message types
	Activity            *MessageActivity      `json:"activity,omitempty"`
	Application         *Application          `json:"application,omitempty"`
	ApplicationID       *int                  `json:"application_ID,omitempty"`
	MessageReference    *MessageReference     `json:"message_reference,omitempty"`
	Flags               *int                  `json:"flags,omitempty"`
	ReferencedMessage   *Message              `json:"referenced_message,omitempty"`
	Interaction         *MessageInteraction   `json:"interaction,omitempty"`
	Thread              *ThreadChannel        `json:"thread,omitempty"`
	Components          *[]Component          `json:"components,omitempty"`
	MessageStickerItems *[]MessageStickerItem `json:"sticker_items,omitempty"`

	ClientPTR *Client `json:"-"`
}

type MessageUpdate struct {
	ID                  Snowflake             `json:"ID"`
	ChannelID           int                   `json:"channel_ID"`
	GuildID             *int                  `json:"guild_ID,omitempty"`
	Author              User                  `json:"author"`
	GuildMember         *GuildMember          `json:"member,omitempty"`
	Content             string                `json:"content"`
	Timestamp           ISO8601Timestamp      `json:"timestamp"`
	EditedTimestamp     *ISO8601Timestamp     `json:"edited_timestamp"`
	TTS                 bool                  `json:"tts"`
	MentionEveryone     bool                  `json:"mention_everyone"`
	Mentions            []User                `json:"mentions"`
	MentionRoles        []int                 `json:"mention_roles"`
	MentionChannels     *[]ChannelMention     `json:"mention_channels,omitempty"`
	Attachments         []Attachment          `json:"attachments"`
	Embeds              []Embed               `json:"embeds"`
	Reactions           *[]Reaction           `json:"reactions,omitempty"`
	Nonce               string                `json:"nonce,omitempty"`
	Pinned              bool                  `json:"pinned"`
	WebhookID           *int                  `json:"webhook_ID,omitempty"`
	Type                int                   `json:"type"` // TODO: Make constants for message types
	Activity            *MessageActivity      `json:"activity,omitempty"`
	Application         *Application          `json:"application,omitempty"`
	ApplicationID       *int                  `json:"application_ID,omitempty"`
	MessageReference    *MessageReference     `json:"message_reference,omitempty"`
	Flags               *int                  `json:"flags,omitempty"`
	ReferencedMessage   *Message              `json:"referenced_message,omitempty"`
	Interaction         *MessageInteraction   `json:"interaction,omitempty"`
	Thread              *ThreadChannel        `json:"thread,omitempty"`
	Components          *[]Component          `json:"components,omitempty"`
	MessageStickerItems *[]MessageStickerItem `json:"sticker_items,omitempty"`

	ClientPTR *Client `json:"-"`
}
