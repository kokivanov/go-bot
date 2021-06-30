package DiscordBotLib

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

//? Mb, swap all nullable fields with pointers
type Message struct { //TODO: Complete it
	Id              int               `json:"id"`
	ChannelId       int               `json:"channel_id"`
	GuildID         *int              `json:"guild_id"`
	Author          User              `json:"author"`
	GuildMember     *GuildMember      `json:"member"`
	Content         string            `json:"content"`
	Timestamp       ISO8601Timestamp  `json:"timestamp"`
	EditedTimestamp ISO8601Timestamp  `json:"edited_timestamp"`
	TTS             bool              `json:"tts"`
	MentionEveryone bool              `json:"mention_everyone"`
	Mentions        []User            `json:"mentions"`
	MentionRoles    []int             `json:"mention_roles"`
	MentionChannels *[]ChannelMention `json:"mention_channels"`
	Attachments     []Attachment      `json:"attachments"`
	Embeds          []Embed           `json:"embeds"`
	Reactions       *[]Reaction       `json:"reactions"`
	// Nonce           string           `json:"nonce"`
	// TODO: Make something with that
	Pinned    bool `json:"pinned"`
	WebhookId *int `json:"webhook_id"`
	Type      int  `json:"type"` // TODO: Make constants for message types

	ClientPTR *Client
}

type Emoji struct {
	Id            int      `json:"id"`
	Name          string   `json:"name"`
	Roles         []RoleID `json:"Roles"`
	User          User     `json:"user"`
	RequireColons bool     `json:"require_colons"`
	Managed       bool     `json:"managed"`
	Animated      bool     `json:"animated"`
	Available     bool     `json:"available"`
}

type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedVideo struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type Embed struct { // TODO: Make functions to create and maintain Embeds
	Title       string           `json:"title"`
	Type        string           `json:"type"`
	Description string           `json:"description"`
	URL         string           `json:"url"`
	Timestamp   ISO8601Timestamp `json:"timestamp"`
	Color       int              `json:"color"`
	Footer      EmbedFooter      `json:"footer"`
	Image       EmbedImage       `json:"image"`
	Thumbnail   EmbedThumbnail   `json:"thumbnail"`
	Video       EmbedVideo       `json:"video"`
	Provider    EmbedProvider    `json:"provider"`
	Author      EmbedAuthor      `json:"author"`
	Fields      []EmbedField     `json:"fields"`
}

type ActivityTimestampsObject struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type ActivityObject struct {
	Name          string                   `json:"name"`
	Type          int                      `json:"type"`
	Url           string                   `json:"url"`
	CreatedAt     int                      `json:"created_at"`
	Timestamps    ActivityTimestampsObject `json:"timestamps"`
	ApplicationId Snowflake                `json:"application_id"`
	Details       string                   `json:"details"`
	State         string                   `json:"state"`

	ClientPTR *Client
}

type IdentifyProperties struct {
	Os      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

type GatewayUpdatePresence struct {
	Since      int              `json:"since"`
	Activities []ActivityObject `json:"activities"`
	Status     string           `json:"status"`
	AFK        bool             `json:"afk"`
}

type Attachment struct {
	Id          int    `json:"id"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	ProxyURL    string `json:"proxy_url"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
}
type ChannelMention struct {
	Id      int    `json:"id"`
	GuildID int    `json:"guild_id"`
	Type    int    `json:"type"`
	Name    string `json:"string"`
}

type Identify struct {
	Token      string             `json:"token"`
	Properties IdentifyProperties `json:"properties"`
	Intents    int                `json:"intents"`
	// They are unnecessaray,
	// TODO: Do something with them
	// Compress       bool               `json:"compress"`
	// LargeThreshold int                `json:"large_threshold"`
	// Shard          [2]int             `json:"shard"`
	// Presence
}

type Payload struct {
	Operation int             `json:"op"`
	Sequence  int             `json:"s"`
	Type      string          `json:"t"`
	RawData   json.RawMessage `json:"d"`
	Struct    interface{}     `json:"-"`
}

type Heartbeat struct {
	Op int  `json:"op"`
	D  *int `json:"d"`

	ClientPTR *Client
}

type TeamMember struct {
	MembershipState int      `json:"membership_state"`
	Permissions     []string `json:"permissions"`
	TeamId          int      `json:"team_id"`
	User            User     `json:"user"`
}

type Team struct {
	Icon        string       `json:"icon"`
	Id          int          `json:"id"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserId int          `json:"owner_user_id"`
}

type Application struct {
	Id                  int      `json:"id"`
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
	GuildId             int      `json:"guild_id"`
	PrimarySkuId        int      `json:"primary_sku_id"`
	Slug                string   `json:"slug"`
	CoverImage          string   `json:"cover_image"`
	Flags               int      `json:"flags"`
}

// Represents User object
// For more information look: https://discord.com/developers/docs/resources/user
type User struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Verified      bool   `json:"verified"`
	Email         string `json:"email"`
	Flags         int    `json:"flags"`
	PremiumType   int    `json:"premium_type"`
	PublicFlags   int    `json:"public_flags"`

	ClientPTR *Client
}

type GuildMember struct {
	User         User   `json:"user"`
	Nick         string `json:"nick"`
	Roles        []int  `json:"roles"`
	JoinedAt     string `json:"joined_at"`
	PremiumSince string `json:"premium_since"`
	Deaf         bool   `json:"deaf"`
	Mute         bool   `json:"mute"`
	Pending      bool   `json:"pending"`
	Permissions  string `json:"permissions"`

	ClientPTR *Client
}

// Client is a base structure that represents your whole bot and methods that are allowed to it
type Client struct {
	// Mutex
	sync.RWMutex
	wsMutex sync.Mutex
	wG      sync.WaitGroup

	// Struncts that are used to work with api
	wsGateway  string // TODO: Make unexported!
	wsConn     *websocket.Conn
	httpClient *http.Client
	authHeader *http.Header

	// Discord related fields
	intent            int
	token             string // TODO: Make unexported!
	heartbeatInterval int    // TODO: Make unexported!
	lastSequence      int
	lastHeartbeatACK  uint64

	handlers map[string]*EventHandler

	// information about application and it's owner
	Owner User
	Me    Application

	// Functions that will be called on events
	OnMessage *OnMessage

	LogLevel  int
	state     uint8
	interrupt chan int

	// TODO: Task Queqe
}

// TODO (Structs)
// - Guild struct
// - Message struct
// - User struct
// - Member struct
