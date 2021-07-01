package DiscordBotLib

import (
	"encoding/json"
)

// ================================ Utility section =====================================

// ------------------------------------------
//                    Utility
// ------------------------------------------
// Contains all utility functions that are will be used during event handling, parsing, making requests, etc.
//

/*
Returns intent (int) depending on handler type.
Used for automatically adding intents for bot to listen to available events that can be handled.*/
func getEventIntent(t string) int {
	switch t {
	case EventMessageCreate, EventMessageDelete, EventMessageUpdate:
		return (IntentGuildMessage | IntentDirectMessage)

	case EventChannelPinsUpdate:
		return (IntentDirectMessage | IntentGuild)

	case EventMessageDeleteBulk:
		return IntentGuildMessage

	case EventTypingStart:
		return (IntentGuildMessageTyping | IntentDirectMessageTyping)
	case EventGuildUpdate, EventGuildCreate, EventGuildDelete, EventGuildRoleCreate, EventGuildRoleUpdate, EventGuildRoleDelete, EventChannelCreate, EventChannelUpdate, EventChannelDelete, EventThreadCreate, EventThreadDelete, EventThreadUpdate, EventThreadListSync, EventThreadMembersUpdate, EventStageInstanceCreate, EventStageInstanceDelete, EventStageInstanceUpdate:
		return IntentGuild
	case EventThreadMemberUpdate:
		return (IntentGuild | IntentGuildMembers)
	case EventGuildBanAdd, EventGuildBanRemove:
		return IntentGuildBans
	case EventGuildEmojisUpdate:
		return IntentGuildEmojis
	case EventGuildIntegrationsUpdate, EventIntegrationCreate, EventIntegrationDelete, EventIntegrationUpdate:
		return IntentGuildIntegrations
	case EventWebhooksUpdate:
		return IntentGuildWebhooks
	case EventInviteCreate, EventInviteDelete:
		return IntentGuildInvites
	case EventVoiceStateUpdate:
		return IntentGuildVoiceStates
	case EventPresenceUpdate:
		return IntentGuildPresences
	case EventMessageReactionAdd, EventMessageReactionRemove, EventMessageReactionRemoveAll, EventMessageReactionRemoveEmoji:
		return (IntentGuildMessageReactions | IntentDirectMessageReaction)
	default:
		return 0
	}
}

/* Returns EventHandler interface{} depending on provided function
In order to add custom events you must add context calling of functions:
	 (T *YourHandler) Handle(*Client, interface{}) // Passes *Client and passes and adapts interface{} to your function
	 (T *YourHandler) Type() // Returns type of function
Example:
	type OnMessage func(*Client, GuildMessage)
	func (om OnMessage) Handle(c *Client, i interface{}) {
		om(c, i.(GuildMessage))
	}
	func (om OnMessage) Type() string {
		return "MESSAGE_CREATE"
	}*/
func getEventHandler(h interface{}, Type string) EventHandler { // TODO: Complete
	switch v := h.(type) {
	case func(*Client):
		switch Type {
		case EventReady:
			return OnReady(v)
		default:
			return nil
		}
	case func(*Client, Message):
		switch Type {
		case EventMessageCreate:
			return OnMessage(v)
		default:
			return nil
		}
	default:
		return nil
	}
}

/* Returns Object that represents payload content (field d) */
func (c *Client) getEventPayload(p Payload) interface{} { // TODO: Complete
	switch p.Type {
	case EventMessageCreate:
		m := Message{
			ClientPTR:   c,
			Author:      User{ClientPTR: c},
			GuildMember: &GuildMember{ClientPTR: c},
		}
		json.Unmarshal(p.RawData, &m)
		return m
	default:
		return nil
	}
}
