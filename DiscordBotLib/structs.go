package DiscordBotLib

import (
	"encoding/json"
)

//? Mb, swap all nullable fields with pointers

type Interaction struct {
	TTS            *bool        `json:"tts,omitempty"`
	Content        string       `json:"content,omitempty"`
	Embeds         *[]Embed     `json:"embeds,omitempty"`
	AllowdMentions *bool        `json:"allowed_mentions,omitempty"`
	Flags          int          `json:"flags,omitempty"`
	Components     *[]Component `json:"components,omitempty"`
}

type Component struct {
	Type       int          `json:"type"`
	Style      *int         `json:"style,omitempty"`
	Label      string       `json:"label,omitempty"`
	Emoji      *Emoji       `json:"emoji,omitempty"`
	CustomID   string       `json:"custom_id,omitempty"`
	URL        *string      `json:"url,omitempty"`
	Disabled   *bool        `json:"daisabled,omitempty"`
	Components *[]Component `json:"components,omitempty"`
}

type GatewayUpdatePresence struct {
	Since      *int             `json:"since"`
	Activities []ActivityObject `json:"activities"`
	Status     string           `json:"status"`
	AFK        bool             `json:"afk"`
}

type RequestGuildMembers struct {
	Op   int                        `json:"op"`
	Data RequestGuildMembersPayload `json:"d"`
}

type RequestGuildMembersPayload struct {
	GuildID   Snowflake    `json:"guild_id"`
	Query     string       `json:"query,omitempty"`
	Limit     int          `json:"limit"`
	Presences *bool        `json:"presences,omitempty"`
	UserIDs   *[]Snowflake `json:"user_ids,omitempty"`
	Nonce     string       `json:"nonce,omitempty"`
}

type Payload struct { // TODO: omitempty
	Operation *int            `json:"op"`
	Sequence  int64           `json:"s,omitempty"`
	Type      string          `json:"t,omitempty"`
	RawData   json.RawMessage `json:"d,omitempty"`
	Struct    interface{}     `json:"-"`
}

type Heartbeat struct { // TODO: omitempty
	Op int   `json:"op"`
	D  int64 `json:"d"`

	ClientPTR *Client `json:"-"`
}

type TeamMember struct { // TODO: omitempty
	MembershipState int       `json:"membership_state"`
	Permissions     []string  `json:"permissions"`
	TeamID          Snowflake `json:"team_id"`
	User            User      `json:"user"`
}

type Team struct { // TODO: omitempty
	Icon        *string      `json:"icon"`
	ID          Snowflake    `json:"id"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID Snowflake    `json:"owner_user_id"`
}

type Application struct {
	ID                  Snowflake `json:"id"`
	Name                string    `json:"name"`
	Icon                string    `json:"icon"`
	Description         string    `json:"description"`
	RPCOrigins          *[]string `json:"rpc_origins,omitempty"`
	BotPublic           bool      `json:"bot_public"`
	BotRequireCodeGrant bool      `json:"bot_require_code_grant"`
	TermsOfServiceURL   string    `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL    string    `json:"privacy_policy_url,omitempty"`
	Owner               *User     `json:"owner,omitempty"`
	Summary             string    `json:"summary"`
	VertifyKey          string    `json:"verify_key"`
	Team                *Team     `json:"team"`
	GuildID             Snowflake `json:"guild_ID,omitempty"`
	PrimarySkuID        Snowflake `json:"primary_sku_ID,omitempty"`
	Slug                string    `json:"slug,omitempty"`
	CoverImage          string    `json:"cover_image,omitempty"`
	Flags               int       `json:"flags,omitempty"`
}

type ApplicationCommand struct {
	ID                Snowflake                   `json:"id"`
	ApplicationID     Snowflake                   `json:"application_id"`
	GuildID           *Snowflake                  `json:"guild_id,omitempty"`
	Name              string                      `json:"name"`              // TODO: validate with regex ("^[\w-]{1,32}$")
	Description       string                      `json:"description"`       // TODO: max 100 char
	Options           *[]ApplicationCommandOption `json:"options,omitempty"` // TODO: Required options must be listed before optional options
	DefaultPermission bool                        `json:"default_permission,omitempty"`
}

type ApplicationCommandOption struct {
	Type        int                               `json:"type"`
	Name        string                            `json:"name"`        // TODO: validate with regex ("^[\w-]{1,32}$")
	Description string                            `json:"description"` // TODO: max 100 char
	Requred     *bool                             `json:"required,omitempty"`
	Choices     *[]ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options     *[]ApplicationCommandOption       `json:"options,omitempty"`
}

type ApplicationCommandOptionChoice struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"` // TODO: Validate string or int
}

type UnavailableGuild struct {
	ID          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailable"`
}

type Ready struct {
	GatewayVersion int                 `json:"v"`
	User           *User               `json:"user"`
	Guilds         *[]UnavailableGuild `json:"guilds"`
	SessionID      string              `json:"session_id"`
	Shard          *[2]int             `json:"shard,omitempty"`
	Application    *Application        `json:"application"`
}

type ChannelPinsUpdate struct {
	GuildID          *Snowflake        `json:"guild_id,omitempty"`
	ChannelID        Snowflake         `json:"channel_id"`
	LastPinTimestamp *ISO8601Timestamp `json:"last_pin_timestamp"`
}

type ThreadListSync struct {
	GuildID    Snowflake      `json:"guild_id,omitempty"`
	ChannelIDs *[]Snowflake   `json:"channel_ids"`
	Threads    []Channel      `json:"threads"`
	Members    []ThreadMember `json:"members"`
}

type ThreadMemberUpdate struct {
	ID                Snowflake       `json:"id"`
	GuildID           Snowflake       `json:"guild_id"`
	MemberCount       int             `json:"member_counts"`
	AddedMembers      *[]ThreadMember `json:"added_members,omitempty"`
	RemovedMembersIDs *[]Snowflake    `json:"removed_member_ids,omitempty"`
}
