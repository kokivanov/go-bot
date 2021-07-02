package DiscordBotLib

type Message struct {
	ID                  Snowflake             `json:"id"`
	ChannelID           int                   `json:"channel_id"`
	GuildID             *int                  `json:"guild_id,omitempty"`
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
	WebhookID           *int                  `json:"webhook_id,omitempty"`
	Type                int                   `json:"type"` // TODO: Make constants for message types
	Activity            *MessageActivity      `json:"activity,omitempty"`
	Application         *Application          `json:"application,omitempty"`
	ApplicationID       *int                  `json:"application_id,omitempty"`
	MessageReference    *MessageReference     `json:"message_reference,omitempty"`
	Flags               *int                  `json:"flags,omitempty"`
	ReferencedMessage   *Message              `json:"referenced_message,omitempty"`
	Interaction         *Interaction          `json:"interaction,omitempty"`
	Thread              *Channel              `json:"thread,omitempty"`
	Components          *[]Component          `json:"components,omitempty"`
	MessageStickerItems *[]MessageStickerItem `json:"sticker_items,omitempty"`

	ClientPTR *Client `json:"-"`
}

type MessageStickerItem struct {
	ID         Snowflake `json:"id"`
	Name       string    `json:"name"`
	FormatType int       `json:"format_type"`
}

type Sticker struct {
	ID          Snowflake `json:"id"`
	PackID      Snowflake `json:"pack_id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`
	FormatType  int       `json:"format_type"`
	Available   *bool     `json:"available,omitempty"`
	GuildID     Snowflake `json:"guild_id,omitempty"`
	User        *User     `json:"user,omitempty"`
	SortValue   *int      `json:"sort_value,omitempty"`
}

type MessageReference struct {
	ChannelID       Snowflake `json:"channel_id,omitempty"`
	GuildID         Snowflake `json:"guild_id,omitempty"`
	MessageID       Snowflake `json:"message_id,omitempty"`
	FailIfNotExists *bool     `json:"fail_if_not_exists,omitempty"`

	ClientPTR *Client `json:"-"`
}

type MessageActivity struct {
	Type    int        `json:"type"`
	PartyID *Snowflake `json:"party_id"`
}

type Emoji struct { // TODO: omitempty
	ID            Snowflake `json:"id"`
	Name          string    `json:"name"`
	Roles         []RoleID  `json:"Roles"`
	User          User      `json:"user"`
	RequireColons bool      `json:"require_colons"`
	Managed       bool      `json:"managed"`
	Animated      bool      `json:"animated"`
	Available     bool      `json:"available"`

	ClientPTR *Client `json:"-"`
}

type Reaction struct { // TODO: omitempty
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

type Attachment struct { // TODO: omitempty
	ID          Snowflake `json:"id"`
	Filename    string    `json:"filename"`
	ContentType string    `json:"content_type"`
	Size        int       `json:"size"`
	URL         string    `json:"url"`
	ProxyURL    string    `json:"proxy_url"`
	Height      int       `json:"height"`
	WIDth       int       `json:"width"`
}
type ChannelMention struct { // TODO: omitempty
	ID      Snowflake `json:"id"`
	GuildID Snowflake `json:"guild_id"`
	Type    int       `json:"type"`
	Name    string    `json:"string"`
}
