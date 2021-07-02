package DiscordBotLib

type User struct {
	ID            Snowflake `json:"id"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Avatar        *string   `json:"avatar"`
	Bot           *bool     `json:"bot,omitempty"`
	System        *bool     `json:"system,omitempty"`
	MFAEnabled    *bool     `json:"mfa_enabled,omitempty"`
	Locale        string    `json:"locale,omitempty"`
	Verified      *bool     `json:"verified,omitempty"`
	Email         string    `json:"email,omitempty"`
	Flags         int       `json:"flags,omitempty"`
	PremiumType   int       `json:"premium_type,omitempty"`
	PublicFlags   int       `json:"public_flags,omitempty"`

	ClientPTR *Client `json:"-"`
}

type GuildMember struct {
	User         *User   `json:"user,omitempty"`
	Nick         *string `json:"nick,omitempty"`
	Roles        []int   `json:"roles"`
	JoinedAt     string  `json:"joined_at"`
	PremiumSince *string `json:"premium_since,omitempty"`
	Deaf         bool    `json:"deaf"`
	Mute         bool    `json:"mute"`
	Pending      *bool   `json:"pending,omitempty"`
	Permissions  string  `json:"permissions,omitempty"`

	ClientPTR *Client `json:"-"`
}
