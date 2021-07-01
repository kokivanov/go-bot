package DiscordBotLib

type User struct { // TODO: omitempty
	ID            int    `json:"ID"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Verified      bool   `json:"verified"`
	Email         string `json:"email"`
	Flags         int    `json:"flags"`
	PremiumType   int    `json:"premium_type"`
	PublicFlags   int    `json:"public_flags"`

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
