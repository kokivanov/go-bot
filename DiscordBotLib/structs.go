package DiscordBotLib

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
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
