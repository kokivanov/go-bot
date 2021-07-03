package DiscordBotLib

type Role struct {
	ID          Snowflake  `json:"id"`
	Name        string     `json:"name"`
	Color       int        `json:"color"`
	Hoist       bool       `json:"hoist"`
	Position    int        `json:"position"`
	Permissions string     `json:"permissions"`
	Managed     bool       `json:"managed"`
	Metionable  bool       `json:"mentionable"`
	RoleTag     *[]RoleTag `json:"tags,omitempty"`

	ClientPTR *Client `json:"-"`
}

type RoleTag struct {
	BotID             Snowflake `json:"bot_id,omitempty"`
	IntegrationId     Snowflake `json:"integration_id,omitempty"`
	PremiumSubscriber *byte     `json:"jsonFieldName"`
}
