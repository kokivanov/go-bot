package DiscordBotLib

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

type EmbedProvider struct { // TODO: omitempty
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EmbedImage struct { // TODO: omitempty
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedVideo struct { // TODO: omitempty
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedThumbnail struct { // TODO: omitempty
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
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
	Video       *EmbedVideo      `json:"video"`
	Provider    *EmbedProvider   `json:"provider"`
	Author      *EmbedAuthor     `json:"author"`
	Fields      []EmbedField     `json:"fields"`
}
