package events

type DispatchEvent string

const (
	Ready             DispatchEvent = "READY"
	Resumed           DispatchEvent = "RESUMED"
	MessageCreate     DispatchEvent = "MESSAGE_CREATE"
	MessageUpdate     DispatchEvent = "MESSAGE_UPDATE"
	MessageDelete     DispatchEvent = "MESSAGE_DELETE"
	ChannelCreate     DispatchEvent = "CHANNEL_CREATE"
	ChannelUpdate     DispatchEvent = "CHANNEL_UPDATE"
	ChannelDelete     DispatchEvent = "CHANNEL_DELETE"
	MemberAdd         DispatchEvent = "MEMBER_ADD"
	MemberUpdate      DispatchEvent = "MEMBER_UPDATE"
	MemberRemove      DispatchEvent = "MEMBER_REMOVE"
	RoleCreate        DispatchEvent = "ROLE_CREATE"
	RoleUpdate        DispatchEvent = "ROLE_UPDATE"
	RoleDelete        DispatchEvent = "ROLE_DELETE"
	PresenceUpdate    DispatchEvent = "PRESENCE_UPDATE"
	TypingStart       DispatchEvent = "TYPING_START"
	ReactionAdd       DispatchEvent = "REACTION_ADD"
	ReactionRemove    DispatchEvent = "REACTION_REMOVE"
	ThreadCreate      DispatchEvent = "THREAD_CREATE"
	ThreadUpdate      DispatchEvent = "THREAD_UPDATE"
	ThreadDelete      DispatchEvent = "THREAD_DELETE"
	InviteCreate      DispatchEvent = "INVITE_CREATE"
	InviteDelete      DispatchEvent = "INVITE_DELETE"
	WebhookUpdate     DispatchEvent = "WEBHOOK_UPDATE"
	EmojiUpdate       DispatchEvent = "EMOJI_UPDATE"
	ServerUpdate      DispatchEvent = "SERVER_UPDATE"
	SlashCommandInvoc DispatchEvent = "SLASH_COMMAND_INVOCATION"
)
