package DiscordBotLib

// TODO: add endpoints

const (
	GatewayOpDispatch = iota
	GatewayOpHeartbeat
	GatewayOpIdentify
	GatewayOpPresenceUpdate
	GatewayOpVoiceStateUpdate
	_
	GatewayOpResume
	GatewayOpReconnect
	GatewayOpRequestGuildMembers
	GatewayOpInvalidSession
	GatewayOpHello
	GatewayOpHeartbeatACK
)

const (
	GatewayCloseUnknownError = iota + 4000
	GatewayCloseUnknownOpcode
	GatewayCloseDecodeError
	GatewayCloseNotAuthenticated
	GatewayCloseAuthenticationFailed
	GatewayCloseAlreadyAuthenticated
	GatewayCloseInvalidSeq
	GatewayCloseRateLimited
	GatewayCloseSessionTimedOut
	GatewayCloseInvalidShard
	GatewayCloseShardingRequired
	GatewayCloseInvalidAPIVersion
	GatewayCloseInvalidIntent
	GatewayCloseDisallowedIntent
)

const (
	VoiceOpIdentify = iota
	VoiceOpSelectProtocol
	VoiceOpReady
	VoiceOpHeartbeat
	VoiceOpSessionDescription
	VoiceOpSpeaking
	VoiceOpHeartbeatACK
	VoiceOpResume
	VoiceOpHello
	VoiceOpResumed

	VoiceOpClientDisconnect = 13
)

const (
	VoiceCloseUnknownOpcode = iota + 4001
	VoiceCloseFailedToDecodePayload
	VoiceCloseNotAuthenticated
	VoiceCloseAuthenticationFailed
	VoiceCloseAlreadyAuthenticated
	VoiceCloseSessionNoLongerValid

	VoiceCloseSessionTimeout = 4009

	VoiceCloseServerNotFound = iota + 4004
	VoiceCloseUnknownProtocol
	VoiceCloseDisconnected
	_
	VoiceCloseVoiceServerCrashed
	VoiceCloseUnknownEncryptionMode
)

const (
	HTTPOk                 = 200
	HTTPCreated            = 201
	HTTPNoContent          = 204
	HTTPNotModified        = 304
	HTTPBadREquest         = 400
	HTTPUnauthorized       = 401
	HTTPForbidden          = 403
	HTTPNotFound           = 404
	HTTPMethodNotAllowed   = 405
	HTTPTooManyRequests    = 429
	HTTPGatewayUnavailable = 502
	HTTPServerError        = "5XX"
)

const (
	RPCUnknownError   = iota + 1000
	RPCInvalidPayload = iota + 3999
	RPCInvalidCommand = iota + 4000
	RPCInvalidGuild
	RPCInvalidEvent
	RPCInvalidChannel
	RPCInvalidPersmission
	RPCInvalidInvalidClientID
	RPCInvalidOrigin
	RPCInvalidToken
	RPCInvalidUser

	RPCOAuthError = iota + 4989
	RPCSelectChannelTimedOut
	RPCGetGuildTimedOut
	RPCSelectVoiceForceREquired
	RPCInvalidCaptureShortcutAlreadyListenig
)

const (
	IntentGuild = 1 << iota
	IntentGuildMembers
	IntentGuildBans
	IntentGuildEmojis
	IntentGuildIntegrations
	IntentGuildWebhooks
	IntentGuildInvites
	IntentGuildVoiceStates
	IntentGuildPresences
	IntentGuildMessage
	IntentGuildMessageReactions
	IntentGuildMessageTyping
	IntentDirectMessage
	IntentDirectMessageReaction
	IntentDirectMessageTyping

	IntentALL = IntentGuild | IntentGuildMembers | IntentGuildBans | IntentGuildEmojis | IntentGuildIntegrations | IntentGuildWebhooks | IntentGuildInvites | IntentGuildVoiceStates | IntentGuildPresences | IntentGuildMessage | IntentGuildMessageReactions | IntentGuildMessageTyping | IntentDirectMessage | IntentDirectMessageReaction | IntentDirectMessageTyping
)

const (
	TeamMemberInvited = iota + 1
	TeamMemberAccepted
)
const (
	UserFlagNone            = 0
	UserFlagDiscordEmployee = 1 << (iota - 1)
	UserFlagPartnerServerOwner
	UserFlagHypeSquadEvents
	UserFlagBugHunterOne
	UserFlagHouseBravery
	UserFlagHouseBrilliance
	UserFlagHouseBalance
	UserFlagEarlySuppotrer
	UserFlagTeamUser
	UserFlagBugHunterTwo
	UserFlag
	_
	ApplicationFlagGatewayPresence
	ApplicationFlagGatewayPresenceLimited
	ApplicationFlagGatewayGuildMembers
	ApplicationFlagGatewayGuildMembersLimited
	ApplicationFlagVertificationPendingGuildLimit
	ApplicationFlagEmbedded
)

const (
	MessageFlagCrossposted = 1 << iota
	MessageFlagIsCrosspostedMessageFlagCrossposted
	MessageFlagSupressEmbeds
	MessageFlagSourceMessageDeleted
	MessageFlagUrgent
	MessageFlagHasThread
	MessageFlagEphemeral
	MessageFlagLoading
)

const (
	LogErrors = 1 << iota
	LogWarnings
	LogMessages
	LogAll
)

const (
	MessageNotificationLevelAll = iota
	MessageNotificationLevelMentions
)

const (
	ExplicitContentFilterDisabled = iota
	ExplicitContentFilterNoRoles
	ExplicitContentFilterAll
)

const (
	MFANone = iota
	MFAElevated
)

const (
	GuildVertificationLevelNone = iota
	GuildVertificationLevelLow
	GuildVertificationLevelMedium
	GuildVertificationLevelHigh
	GuildVertificationLevelVEryHigh
)

const (
	GuildNSFWDefault = iota
	GuildNSFWExplicit
	GuildNSFWSafe
	GuildNSFWAgeRestricted
)

const (
	GuildPremiumTierNone = iota
	GuildPremiumTier1
	GuildPremiumTier2
	GuildPremiumTier3
)

const (
	SystemChannelSupressJoin = 1 << iota
	SystemChannelSupressBoost
	SystemChannelSupressTips
)

const (
	ApplicationCommandSubcommand = iota + 1
	ApplicationCommandSubcommandGroup
	ApplicationCommandString
	ApplicationCommandInteger
	ApplicationCommandBoolean
	ApplicationCommandUser
	ApplicationCommandChannel
	ApplicationCommandRole
	ApplicationCommandMentionable
)

const (
	ChannelGuildText = iota
	ChannelDM
	ChannelGuildVoice
	ChannelGroupDM
	ChannelGuildCategory
	ChannelGuildNews
	ChannelGuildStore
	_
	_
	_
	ChannelGuildNewsThread
	ChannelPublicThread
	ChannelPrivateThread
	ChannelStageVoice
)

const (
	VideoQualityAuto = 1
	VideoQualityFull = 2
)

const (
	MessageDefault = iota
	MessageRecipientAdd
	MessageRecipientRemove
	MessageCall
	MessageChannelNameChange
	MessageChannelIconChange
	MessageChannelPinnedMessage
	MessageGuildMemberJoin
	MessageUserPremiumGuilSub
	MessageUserPremiumGuilSub1
	MessageUserPremiumGuilSub2
	MessageUserPremiumGuilSub3
	MessageChannelFollowAdd
	_
	MessageGuildDiscoveryDisqualified
	MessageGuildDiscoveryRequalified
	MessageGuildDiscoveryGracePeriodInitialWarning
	MessageGuildDiscoveryGracePeriodFinalWarning
	MessageThreadCreated
	MessageReply
	MessageApplicationCommand
	MessageThreadStartedMessage
	MessageGuildInviteReminder
)

const (
	MessageStickerPNG = iota + 1
	MessageStickerAPNG
	MessageStickerLOTTIE
)

const (
	InteractionCallbackPong                     = iota + 1
	InteractionCallbackChannelMessageWithSource = iota + 3
	InteractionCallbackDeferredChannelMessageWithSource
	InteractionCallbackDefferedUpdateMessage
	InteractionCallbackUpdateMessage
)

const (
	ButtonPrimary = iota + 1
	ButtonSecondary
	ButtonSuccess
	ButtonDanger
	ButtonLink
)

const (
	ComponentActionRow = iota + 1
	ComponentButton
	ComponentSelectMenu
)
