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
	StatusOnline    = "online"
	StatusDnD       = "dnd"
	StatusAFK       = "idle"
	StatusInvisible = "invisible"
	StatusOffline   = "offline"
)

const (
	IntentGuild = 1 << iota
	IntentGuildMembers
	IntentGuildBans
	IntentGuildEmojis
	IntentGuildintegrations
	IntentGuildWebhooks
	IntentGuildInvites
	IntentGuildVoiceStates
	IntentGuildPresences
	IntentGuildMessage
	IntentGuildMessageReaction
	IntentGuildMEssageTyping
	IntentDirectMessage
	IntentDirectMessageReaction
	IntentDirectMessageTyping

	IntentALL = IntentGuild | IntentGuildMembers | IntentGuildBans | IntentGuildEmojis | IntentGuildintegrations | IntentGuildWebhooks | IntentGuildInvites | IntentGuildVoiceStates | IntentGuildPresences | IntentGuildMessage | IntentGuildMessageReaction | IntentGuildMEssageTyping | IntentDirectMessage | IntentDirectMessageReaction | IntentDirectMessageTyping
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
)

const (
	LogErrors = 1 << iota
	LogWarnings
	LogMessages
	LogAll
)
