package events

// DispatchEvent identifies the type of a gateway dispatch event. Each
// constant corresponds to a specific real-time event delivered over the
// WebSocket connection.
type DispatchEvent string

const (
	// Ready is sent after a successful identify, containing initial state.
	Ready DispatchEvent = "READY"
	// Resumed is sent after a successful session resume.
	Resumed DispatchEvent = "RESUMED"

	// MessageCreate is sent when a message is created in a channel.
	MessageCreate DispatchEvent = "MESSAGE_CREATE"
	// MessageUpdate is sent when a message is edited.
	MessageUpdate DispatchEvent = "MESSAGE_UPDATE"
	// MessageDelete is sent when a message is deleted.
	MessageDelete DispatchEvent = "MESSAGE_DELETE"

	// ChannelCreate is sent when a channel is created.
	ChannelCreate DispatchEvent = "CHANNEL_CREATE"
	// ChannelUpdate is sent when a channel is edited.
	ChannelUpdate DispatchEvent = "CHANNEL_UPDATE"
	// ChannelDelete is sent when a channel is deleted.
	ChannelDelete DispatchEvent = "CHANNEL_DELETE"

	// MemberAdd is sent when a user joins the server.
	MemberAdd DispatchEvent = "MEMBER_ADD"
	// MemberUpdate is sent when a member's roles or nickname change.
	MemberUpdate DispatchEvent = "MEMBER_UPDATE"
	// MemberRemove is sent when a member leaves or is removed.
	MemberRemove DispatchEvent = "MEMBER_REMOVE"

	// RoleCreate is sent when a role is created.
	RoleCreate DispatchEvent = "ROLE_CREATE"
	// RoleUpdate is sent when a role is edited.
	RoleUpdate DispatchEvent = "ROLE_UPDATE"
	// RoleDelete is sent when a role is deleted.
	RoleDelete DispatchEvent = "ROLE_DELETE"

	// PresenceUpdate is sent when a user's online status changes.
	PresenceUpdate DispatchEvent = "PRESENCE_UPDATE"
	// TypingStart is sent when a user begins typing in a channel.
	TypingStart DispatchEvent = "TYPING_START"

	// ReactionAdd is sent when a reaction is added to a message.
	ReactionAdd DispatchEvent = "REACTION_ADD"
	// ReactionRemove is sent when a reaction is removed from a message.
	ReactionRemove DispatchEvent = "REACTION_REMOVE"

	// ThreadCreate is sent when a thread is created.
	ThreadCreate DispatchEvent = "THREAD_CREATE"
	// ThreadUpdate is sent when a thread is edited.
	ThreadUpdate DispatchEvent = "THREAD_UPDATE"
	// ThreadDelete is sent when a thread is deleted.
	ThreadDelete DispatchEvent = "THREAD_DELETE"

	// InviteCreate is sent when an invite is created.
	InviteCreate DispatchEvent = "INVITE_CREATE"
	// InviteDelete is sent when an invite is revoked or expires.
	InviteDelete DispatchEvent = "INVITE_DELETE"

	// WebhookUpdate is sent when a webhook is created, edited, or deleted.
	WebhookUpdate DispatchEvent = "WEBHOOK_UPDATE"
	// EmojiUpdate is sent when custom emoji are added, edited, or removed.
	EmojiUpdate DispatchEvent = "EMOJI_UPDATE"
	// ServerUpdate is sent when server settings change.
	ServerUpdate DispatchEvent = "SERVER_UPDATE"
	// SlashCommandInvocation is sent when a slash command is invoked.
	SlashCommandInvocation DispatchEvent = "SLASH_COMMAND_INVOCATION"
)
