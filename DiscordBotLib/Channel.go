package DiscordBotLib

type Channel struct {
	ID                         Snowflake         `json:"id"`
	Type                       int               `json:"type"`
	GuildID                    *Snowflake        `json:"guild_id,omitempty"`
	Position                   int               `json:"position,omitempty"`
	PermissionOverwrites       *[]Overwrite      `json:"permission_overwrites,omitempty"`
	Name                       string            `json:"name,omitempty"`
	Topic                      *string           `json:"topic,omitempty"`
	NSFW                       bool              `json:"nsfw"`
	LastMessageID              *Snowflake        `json:"last_message_id,omitempty"`
	Bitrate                    int               `json:"bitrate,omitempty"`
	UserLimit                  int               `json:"user_limit,omitempty"`
	RateLimitPerUser           int               `json:"rate_limit_per_user,omitempty"`
	Recipients                 *[]User           `json:"recipients,omitempty"`
	Icon                       *string           `json:"icon,omitempty"`
	OwnerID                    Snowflake         `json:"owner_id,omitempty"`
	ApplicationID              Snowflake         `json:"application_id,omitempty"`
	ParentID                   *Snowflake        `json:"parent_id,omitempty"`
	LastPinTimestamp           *ISO8601Timestamp `json:"last_pin_timestamp,omitempty"`
	RTCRegion                  *string           `json:"rtc_region,omitempty"`
	VideoQualityMode           int               `json:"video_quality_mode,omitempty"`
	MessageCount               *int              `json:"message_count,omitempty"`
	MemberCount                *int              `json:"member_count,omitempty"`
	ThreadMetadata             *ThreadMetadata   `json:"thread_metadata,omitempty"`
	Member                     *ThreadMember     `json:"member,omitempty"`
	DefaultAutoArchiveDuration int               `json:"default_auto_archive_duration,omitempty"`
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

// TODO: GetMessage, GetMessages, GetPins, Purge, Delete (Close if thread), Send
