package DiscordBotLib

type Guild struct {
	ID                          Snowflake                `json:"id"`
	Name                        string                   `json:"name"`
	Icon                        *string                  `json:"icon"`
	IconHash                    *string                  `json:"icon_hash,omitempty"`
	Description                 string                   `json:"description"`
	Splash                      *string                  `json:"splash"`
	DiscoverySplash             *string                  `json:"discovery_splash"`
	Owner                       *bool                    `json:"owner,omitempty"`
	OwnerID                     Snowflake                `json:"owner_id"`
	Permissions                 *[]string                `json:"permissions,omitempty"`
	Region                      *[]string                `json:"region,omitempty"`
	AfkChannelID                *Snowflake               `json:"afk_channel_id"`
	AfkTimeout                  int64                    `json:"afk_timeout"`
	WidgetEnabled               *bool                    `json:"widget_enabled,omitempty"`
	WidgetChannelID             *Snowflake               `json:"widget_channel_id,omitempty"`
	VerificationLevel           int64                    `json:"verification_level"`
	DefaultMessageNotifications int64                    `json:"default_message_notifications"`
	ExplicitContentFilter       int64                    `json:"explicit_content_filter"`
	Roles                       []Role                   `json:"roles"`
	Emojis                      []Emoji                  `json:"emojis"`
	Features                    []string                 `json:"features"`
	MfaLevel                    int64                    `json:"mfa_level"`
	ApplicationID               *Snowflake               `json:"application_id"`
	SystemChannelID             *Snowflake               `json:"system_channel_id"`
	SystemChannelFlags          int64                    `json:"system_channel_flags"`
	RulesChannelID              *Snowflake               `json:"rules_channel_id"`
	JoinedAt                    *ISO8601Timestamp        `json:"joined_at"`
	Large                       *bool                    `json:"large,omitempty"`
	Unavailable                 *bool                    `json:"unavailable"`
	MemberCount                 *int                     `json:"member_count,omitempty"`
	VoiceStates                 *[]VoiceState            `json:"voice_states"`
	Members                     *[]GuildMember           `json:"members,omitempty"`
	Channels                    *[]Channel               `json:"channels,omitempty"`
	Threads                     *[]Channel               `json:"threads"`
	Presences                   *[]GatewayUpdatePresence `json:"presences,omitempty"`
	MaxPresences                *int                     `json:"max_presences,omitempty"`
	MaxMembers                  int64                    `json:"max_members,omitempty"`
	VanityURLCode               string                   `json:"vanity_url_code"`
	Banner                      string                   `json:"banner"`
	PremiumTier                 int64                    `json:"premium_tier"`
	PremiumSubscriptionCount    int64                    `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string                   `json:"preferred_locale"`
	PublicUpdatesChannelID      *Snowflake               `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        *int                     `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount      *int                     `json:"approximate_member_count,omitempty"`
	WelcomeScreen               []WelcomeScreen          `json:"welcome_screen"`
	NSFWLevel                   int                      `json:"nsfw_level"`
	StageInstances              *[]StageInstance         `json:"stage_instances"`

	ClientPTR *Client `json:"-"`
}

type WelcomeScreen struct {
	Description     string                   `json:"description"`
	WelcomeChannels [5]*WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelID   Snowflake  `json:"channel_id"`
	Description string     `json:"description"`
	EmojiID     *Snowflake `json:"emoji_id"`
	EmojiName   *string    `json:"emoji_name"`
}

type StageInstance struct {
	ID                   string `json:"id"`
	GuildID              string `json:"guild_id"`
	ChannelID            string `json:"channel_id"`
	Topic                string `json:"topic"`
	PrivacyLevel         int64  `json:"privacy_level"`
	DiscoverableDisabled bool   `json:"discoverable_disabled"`
}
