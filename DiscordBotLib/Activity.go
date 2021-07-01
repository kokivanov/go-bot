package DiscordBotLib

type ActivityObject struct {
	Name          string              `json:"name"`
	Type          int                 `json:"type"`
	URL           string              `json:"url,omitempty"`
	ApplicationID *Snowflake          `json:"application_id,omitempty"`
	State         *string             `json:"state,omitempty"`
	Details       *string             `json:"details,omitempty"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Secrets       *ActivitySecrets    `json:"secrets,omitempty"`
	Emoji         *ActivityEmoji      `json:"emoji,omitempty"`
	Instance      bool                `json:"instance,omitempty"`
	Flags         int                 `json:"flags,omitempty"`
	Buttons       *[2]ActivityButton  `json:"buttons,omitempty"`
}

type ActivityEmoji struct {
	Name     string    `json:"name"`
	ID       Snowflake `json:"id,omitempty"`
	Animated bool      `json:"animated,omitempty"`
}

type ActivityButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivityParty struct {
	ID   string  `json:"id,omitempty"`
	Size *[2]int `json:"size,omitempty"`
}

type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type ActivityTimestamps struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}
