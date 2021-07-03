package DiscordBotLib

import "time"

type Snowflake string
type RoleID int
type ISO8601Timestamp string // TODO: Add "toUnix" function

func (ts ISO8601Timestamp) ToString() string {
	return string(ts)
}

func (ts ISO8601Timestamp) ToUnix() int64 {
	t, err := time.Parse(time.RFC3339, string(ts))
	if err != nil {
		return 0
	}

	return t.Unix()
}

func (ts ISO8601Timestamp) ToTimeObject() *time.Time {
	t, err := time.Parse(time.RFC3339, string(ts))
	if err != nil {
		return nil
	}

	return &t
}
