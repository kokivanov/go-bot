package DiscordBotLib

const APIVersion = 9
const GatewayVersion = 9
const GatewayEncoding = "json"
const LibName = "Koki's discord lib"

const (
	APIURL             = "https://discord.com/api"
	ImageBaseUrl       = "https://cdn.discordapp.com/"
	GetGatewayEndpoint = "/gateway"
)

const (
	RichEmbed    = "rich"
	ImageEmbed   = "image"
	VideoEmbed   = "video"
	GIFVEmbed    = "gifv"
	ArticleEmbed = "article"
	LinkEmbed    = "link"
)

const (
	ActivityStatusOnline    = "online"
	ActivityStatusDnD       = "dnd"
	ActivityStatusAFK       = "idle"
	ActivityStatusInvisible = "invisible"
	ActivityStatusOffline   = "offline"
)

// TODO: Add Event types constants
const (
	EventReady   = "READY"
	EventResumed = "RESUMED"

	EventApplicationCommandCreate = "APPLICATION_COMMAND_CREATE"
	EventApplicationCommandUpdate = "APPLICATION_COMMAND_UPDATE"
	EventApplicationCommandDelete = "APPLICATION_COMMAND_DELETE"

	EventChannelCreate     = "CHANNEL_CREATE"
	EventChannelUpdate     = "CHANNEL_UPDATE"
	EventChannelDelete     = "CHANNEL_DELETE"
	EventChannelPinsUpdate = "CHANNEL_PINS_UPDATE"

	EventThreadCreate        = "THREAD_CREATE"
	EventThreadUpdate        = "THREAD_UPDATE"
	EventThreadDelete        = "THREAD_DELETE"
	EventThreadListSync      = "THREAD_LIST_SYNC"
	EventThreadMemberUpdate  = "THREAD_THREAD_MEMBER_UPDATE"
	EventThreadMembersUpdate = "THREAD_THREAD_MEMBERS_UPDATE"

	EventGuildCreate             = "GUILD_CREATE"
	EventGuildUpdate             = "GUILD_UPDATE"
	EventGuildDelete             = "GUILD_DELETE"
	EventGuildBanAdd             = "GUILD_BAN_ADD"
	EventGuildBanRemove          = "GUILD_BAN_REMOVE"
	EventGuildEmojisUpdate       = "GUILD_EMOJIS_UPDATE"
	EventGuildIntegrationsUpdate = "GUILD_INTEGRATIONS_UPDATE"
	EventGuildMemberAdd          = "GUILD_MEMBER_ADD"
	EventGuildMemberRemove       = "GUILD_MEMBER_REMOVE"
	EventGuildMemberUpdate       = "GUILD_MEMBER_UPDATE"
	EventGuildRoleCreate         = "GUILD_ROLE_CREATE"
	EventGuildRoleUpdate         = "GUILD_ROLE_UPDATE"
	EventGuildRoleDelete         = "GUILD_ROLE_DELETE"

	EventIntegrationCreate = "INTEGRATION_CREATE"
	EventIntegrationUpdate = "INTEGRATION_UPDATE"
	EventIntegrationDelete = "INTEGRATION_DELETE"

	EventInteractionCreate = "INTERACTION_CREATE"

	EventInviteCreate = "INVITE_CREATE"
	EventInviteDelete = "INVITE_DELETE"

	EventMessageCreate              = "MESSAGE_CREATE"
	EventMessageUpdate              = "MESSAGE_UPDATE"
	EventMessageDelete              = "MESSAGE_DELETE"
	EventMessageDeleteBulk          = "MESSAGE_DELETE_BULK"
	EventMessageReactionAdd         = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove      = "MESSAGE_REACTION_REMOVE"
	EventMessageReactionRemoveAll   = "MESSAGE_REACTION_REMOVE_ALL"
	EventMessageReactionRemoveEmoji = "MESSAGE_REACTION_REMOVE_EMOJI"

	EventPresenceUpdate = "PRESENCE_UPDATE"

	EventStageInstanceCreate = "STAGE_INSTANCE_CREATE"
	EventStageInstanceDelete = "STAGE_INSTANCE_DELETE"
	EventStageInstanceUpdate = "STAGE_INSTANCE_UPDATE"

	EventTypingStart = "TYPING_START"

	EventVoiceStateUpdate = "VOICE_STATE_UPDATE"

	EventVoiceServerUpdate = "VOICE_SERVER_UPDATE"

	EventWebhooksUpdate = "WEBHOOKS_UPDATE"

	EventGuildMembersChunk = "GUILD_MEMBERS_CHUNK"
)

const (
	GuildFetureAnimatedIcon     = "ANIMATED_ICON"
	GuildFetureBanner           = "BANNER"
	GuildFetureCommerce         = "COMMERCE"
	GuildFetureCommunity        = "COMMUNITY"
	GuildFetureDiscoverable     = "DISCOVERABLE"
	GuildFetureFeaturable       = "FEATURABLE"
	GuildFetureInviteSplash     = "INVITE_SPLASH"
	GuildFetureMembershipScreen = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFetureNews             = "NEWS"
	GuildFeturePartnered        = "PARTNERED"
	GuildFeturePreview          = "PREVIEW_ENABLED"
	GuildFetureVanityURL        = "VANITY_URL"
	GuildFetureVerified         = "VERIFIED"
	GuildFetureVIPRegion        = "VIP_REGIONS"
	GuildFetureWelcomeScreen    = "WELCOME_SCREEN_ENABLED"
	GuildFetureTicketedEvents   = "TICKETED_EVENTS_ENABLED"
	GuildFetureMonetization     = "MONETIZATION_ENABLED"
	GuildFetureMoreStickers     = "MORE_STICKERS"
	GuildFeture3DaysThreads     = "THREE_DAY_THREAD_ARCHIVE"
	GuildFeture7DaysThreads     = "SEVEN_DAY_THREAD_ARCHIVE"
	GuildFeturePrivateThreads   = "PRIVATE_THREADS"
)

// TODO: Add gateway constants
