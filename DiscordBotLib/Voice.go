package DiscordBotLib

type VoiceState struct {
	GuildID                 Snowflake         `json:"guild_id,omitempty"`
	ChannelID               Snowflake         `json:"channel_id"`
	UserID                  Snowflake         `json:"user_id"`
	SessionID               string            `json:"session_id"`
	Member                  *GuildMember      `json:"member,omitempty"`
	Deaf                    bool              `json:"deaf"`
	Mute                    bool              `json:"mute"`
	SelfDeaf                bool              `json:"self_deaf"`
	SelfMute                bool              `json:"self_mute"`
	SelfStream              *bool             `json:"self_stream,omitempty"`
	SelfVideo               bool              `json:"self_video"`
	Suppress                bool              `json:"suppress"`
	RequestToSpeakTimestamp *ISO8601Timestamp `json:"request_to_speak_timestamp"`
}
