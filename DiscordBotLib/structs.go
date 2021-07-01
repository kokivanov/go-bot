package DiscordBotLib

import (
	"encoding/json"
)

//? Mb, swap all nullable fields with pointers

type GuildChannel struct { // TODO: omitempty
	ID                         Snowflake    `json:"id"`
	GuildID                    *Snowflake   `json:"guild_id"`
	Name                       string       `json:"name"`
	Type                       int          `json:"type"`
	Position                   *int         `json:"position"`
	PermissionOverwrites       *[]Overwrite `json:"permission_overwrites"`
	RateLimitPerUser           int          `json:"rate_limit_per_user"`
	Nsfw                       *bool        `json:"nsfw"`
	Topic                      *string      `json:"topic"`
	LastMessageID              *Snowflake   `json:"last_message_id"`
	ParentID                   Snowflake    `json:"parent_id"`
	DefaultAutoArchiveDuration int          `json:"default_auto_archive_duration"`

	ClientPTR *Client `json:"-"`
}

type VoiceChannel struct { // TODO: omitempty
	ID                   Snowflake    `json:"id"`
	GuildID              *Snowflake   `json:"guild_id"`
	Name                 string       `json:"name"`
	Type                 int          `json:"type"`
	Nsfw                 bool         `json:"nsfw"`
	Position             int          `json:"position"`
	PermissionOverwrites *[]Overwrite `json:"permission_overwrites"`
	Bitrate              int          `json:"bitrate"`
	UserLimit            int          `json:"user_limit"`
	ParentID             *Snowflake   `json:"parent_id"`
	RTCRegion            *Snowflake   `json:"rtc_region"`

	ClientPTR *Client `json:"-"`
}

type ThreadChannel struct { // TODO: omitempty
	ID                         Snowflake       `json:"id"`
	GuildID                    Snowflake       `json:"guild_id"`
	ParentID                   *Snowflake      `json:"parent_id"`
	OwnerID                    Snowflake       `json:"owner_id"`
	Name                       string          `json:"name"`
	Type                       int             `json:"type"`
	LastMessageID              *Snowflake      `json:"last_message_id"`
	MessageCount               int             `json:"message_count"`
	MemberCount                int             `json:"member_count"`
	RateLimitPerUser           *int            `json:"rate_limit_per_user"`
	Member                     *ThreadMember   `json:"member"`
	ThreadMetadata             *ThreadMetadata `json:"thread_metadata"`
	DefaultAutoArchiveDuration int             `json:"default_auto_archive_duration"`

	ClientPTR *Client `json:"-"`
}

type DMChannel struct { // TODO: omitempty
	LastMessageID *Snowflake  `json:"last_message_id"`
	Type          int         `json:"type"`
	ID            Snowflake   `json:"id"`
	Recipients    []Recipient `json:"recipients"`

	ClientPTR *Client `json:"-"`
}

type GroupDMChannel struct { // TODO: omitempty
	Name          string       `json:"name"`
	Icon          *string      `json:"icon"`
	Recipients    *[]Recipient `json:"recipients"`
	LastMessageID *Snowflake   `json:"last_message_id"`
	Type          int          `json:"type"`
	ID            Snowflake    `json:"id"`
	OwnerID       Snowflake    `json:"owner_id"`

	ClientPTR *Client `json:"-"`
}

type ThreadMember struct { // TODO: omitempty
	ID            *Snowflake       `json:"id"`
	UserID        *Snowflake       `json:"user_id"`
	JoinTimestamp ISO8601Timestamp `json:"join_timestamp"`
	Flags         int              `json:"flags"`
}

type ThreadMetadata struct { // TODO: omitempty
	Archived            bool   `json:"archived"`
	AutoArchiveDuration int64  `json:"auto_archive_duration"`
	ArchiveTimestamp    string `json:"archive_timestamp"`
	Locked              bool   `json:"locked"`
}

type StoreChannel struct { // TODO: omitempty
	ID                   Snowflake    `json:"id"`
	GuildID              Snowflake    `json:"guild_id"`
	Name                 string       `json:"name"`
	Type                 int          `json:"type"`
	Position             int          `json:"position"`
	PermissionOverwrites *[]Overwrite `json:"permission_overwrites"`
	Nsfw                 *bool        `json:"nsfw"`
	ParentID             *Snowflake   `json:"parent_id"`

	ClientPTR *Client `json:"-"`
}

type ChannelCategory struct { // TODO: omitempty
	PermissionOverwrites *[]Overwrite `json:"permission_overwrites"`
	Name                 string       `json:"name"`
	ParentID             *Snowflake   `json:"parent_id"`
	Nsfw                 bool         `json:"nsfw"`
	Position             int          `json:"position"`
	GuildID              Snowflake    `json:"guild_id"`
	Type                 int          `json:"type"`
	ID                   Snowflake    `json:"id"`

	ClientPTR *Client `json:"-"`
}

type Recipient struct { // TODO: omitempty
	Username      string    `json:"username"`
	Discriminator int       `json:"discriminator"`
	ID            Snowflake `json:"id"`
	Avatar        *string   `json:"avatar"`

	ClientPTR *Client `json:"-"`
}

type Overwrite struct { // TODO: omitempty
	ID    Snowflake `json:"id"`
	Type  int       `json:"type"`
	Allow string    `json:"allow"`
	Deny  string    `json:"deny"`
}

type MessageStickerItem struct { // TODO: omitempty
	ID         Snowflake `json:"ID"`
	Name       string    `json:"name"`
	FormatType int       `json:"format_type"`
}

type Sticker struct { // TODO: omitempty
	ID          Snowflake  `json:"ID"`
	PackID      *Snowflake `json:"pack_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Tags        string     `json:"tags"`
	FormatType  int        `json:"format_type"`
	Available   *bool      `json:"available"`
	GuildID     *int       `json:"guild_ID"`
	User        *User      `json:"user"`
	SortValue   *int       `json:"sort_value"`
}

type MessageInteraction struct { // TODO: omitempty
	TTS            *bool        `json:"tts"`
	Content        *string      `json:"content"`
	Embeds         *[]Embed     `json:"embeds"`
	AllowdMentions *bool        `json:"allowed_mentions"`
	Flags          *int         `json:"flags"`
	Components     *[]Component `json:"components"`
}

type Component struct { // TODO: omitempty
	Type       int          `json:"type"`
	Style      *int         `json:"style"`
	Label      *string      `json:"label"`
	Emoji      *Emoji       `json:"emoji"`
	CustomID   *int         `json:"custom_ID"`
	URL        *string      `json:"url"`
	Disabled   *bool        `json:"daisabled"`
	Components *[]Component `json:"components"`
}

type MessageReference struct { // TODO: omitempty
	ChannelID       *string `json:"channel_ID"`
	GuildID         *string `json:"guild_ID"`
	MessageID       *string `json:"message_ID"`
	FailIfNotExists *bool   `json:"fail_if_not_exists"`

	ClientPTR *Client `json:"-"`
}

type MessageActivity struct { // TODO: omitempty
	Type    int    `json:"type"`
	PartyID string `json:"party_ID"`
}

type Emoji struct { // TODO: omitempty
	ID            int      `json:"ID"`
	Name          string   `json:"name"`
	Roles         []RoleID `json:"Roles"`
	User          User     `json:"user"`
	RequireColons bool     `json:"require_colons"`
	Managed       bool     `json:"managed"`
	Animated      bool     `json:"animated"`
	Available     bool     `json:"available"`

	ClientPTR *Client `json:"-"`
}

type Reaction struct { // TODO: omitempty
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

type EmbedFooter struct { // TODO: omitempty
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedField struct { // TODO: omitempty
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type EmbedAuthor struct { // TODO: omitempty
	Name         string `json:"name"`
	URL          string `json:"url"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedProvIDer struct { // TODO: omitempty
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EmbedImage struct { // TODO: omitempty
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	WIDth    int    `json:"wIDth"`
}

type EmbedVIDeo struct { // TODO: omitempty
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	WIDth    int    `json:"wIDth"`
}

type EmbedThumbnail struct { // TODO: omitempty
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	WIDth    int    `json:"wIDth"`
}

// TODO: omitempty
type Embed struct { // TODO: Make functions to create and maintain Embeds
	Title       string           `json:"title"`
	Type        string           `json:"type"`
	Description string           `json:"description"`
	URL         string           `json:"url"`
	Timestamp   ISO8601Timestamp `json:"timestamp"`
	Color       int              `json:"color"`
	Footer      *EmbedFooter     `json:"footer"`
	Image       *EmbedImage      `json:"image"`
	Thumbnail   *EmbedThumbnail  `json:"thumbnail"`
	VIDeo       *EmbedVIDeo      `json:"vIDeo"`
	ProvIDer    *EmbedProvIDer   `json:"provIDer"`
	Author      *EmbedAuthor     `json:"author"`
	Fields      []EmbedField     `json:"fields"`
}

type IdentifyProperties struct { // TODO: omitempty
	Os      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

type GatewayUpdatePresence struct { // TODO: omitempty
	Since      int              `json:"since"`
	Activities []ActivityObject `json:"activities"`
	Status     string           `json:"status"`
	AFK        bool             `json:"afk"`
}

type Attachment struct { // TODO: omitempty
	ID          int    `json:"ID"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	ProxyURL    string `json:"proxy_url"`
	Height      int    `json:"height"`
	WIDth       int    `json:"wIDth"`
}
type ChannelMention struct { // TODO: omitempty
	ID      int    `json:"ID"`
	GuildID int    `json:"guild_ID"`
	Type    int    `json:"type"`
	Name    string `json:"string"`
}

type Resume struct { // TODO: omitempty
	Op   int           `json:"op"`
	Data ResumePayload `json:"d"`
}

type ResumePayload struct { // TODO: omitempty
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Sequence  int    `json:"seq"`
}

type RequestGuildMembersQuery struct { // TODO: omitempty
	Op   int                             `json:"op"`
	Data RequestGuildMembersPayloadQuery `json:"d"`
}

type RequestGuildMembersID struct { // TODO: omitempty
	Op   int                             `json:"op"`
	Data RequestGuildMembersPayloadQuery `json:"d"`
}

type RequestGuildMembersPayloadQuery struct { // TODO: omitempty
	GuildID   Snowflake `json:"guild_id"`
	Query     string    `json:"query"`
	Limit     int       `json:"limit"`
	Presences bool      `json:"presences"`
}

type Payload struct { // TODO: omitempty
	Operation int             `json:"op"`
	Sequence  int             `json:"s"`
	Type      string          `json:"t"`
	RawData   json.RawMessage `json:"d"`
	Struct    interface{}     `json:"-"`
}

type Heartbeat struct { // TODO: omitempty
	Op int  `json:"op"`
	D  *int `json:"d"`

	ClientPTR *Client `json:"-"`
}

type TeamMember struct { // TODO: omitempty
	MembershipState int      `json:"membership_state"`
	Permissions     []string `json:"permissions"`
	TeamID          int      `json:"team_ID"`
	User            User     `json:"user"`
}

type Team struct { // TODO: omitempty
	Icon        string       `json:"icon"`
	ID          int          `json:"ID"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID int          `json:"owner_user_ID"`
}

type Application struct { // TODO: omitempty
	ID                  int      `json:"ID"`
	Name                string   `json:"name"`
	Icon                string   `json:"icon"`
	Description         string   `json:"description"`
	RPCOrigins          []string `json:"rpc_origins"`
	BotPublic           bool     `json:"bot_public"`
	BotRequireCodeGrant bool     `json:"bot_require_code_grant"`
	TermsOfServiceURL   string   `json:"terms_of_service_url"`
	PrivacyPolicyURL    string   `json:"privacy_policy_url"`
	Owner               User     `json:"owner"`
	Summary             string   `json:"summary"`
	VertifyKey          string   `json:"verify_key"`
	Team                Team     `json:"team"`
	GuildID             int      `json:"guild_ID"`
	PrimarySkuID        int      `json:"primary_sku_ID"`
	Slug                string   `json:"slug"`
	CoverImage          string   `json:"cover_image"`
	Flags               int      `json:"flags"`
}

// Represents User object
// For more information look: https://discord.com/developers/docs/resources/user

// Client is a base structure that represents your whole bot and methods that are allowed to it

// TODO (Structs)
// - Guild struct
// - Message struct
// - User struct
// - Member struct
