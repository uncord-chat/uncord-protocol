// Package permissions defines the bitfield type and constants used to
// represent server, category, and channel-level permissions in the Uncord
// protocol. Both the server and any future clients import this package as
// the single source of truth.
package permissions

// Permission is an int64 bitfield where each bit represents a distinct
// privilege. Bits are grouped by category: general, membership, text,
// voice, and moderation.
type Permission int64

const (
	// --- General (bits 0-9) ---

	// ViewChannels allows seeing channels in the channel list.
	ViewChannels Permission = 1 << 0
	// ManageChannels allows creating, editing, and deleting channels.
	ManageChannels Permission = 1 << 1
	// ManageCategories allows creating, editing, and deleting categories.
	ManageCategories Permission = 1 << 2
	// ManageRoles allows creating, editing, and deleting roles.
	ManageRoles Permission = 1 << 3
	// ManageServer grants full administrative access (treated as admin).
	ManageServer Permission = 1 << 4
	// ManageInvites allows revoking any invite link.
	ManageInvites Permission = 1 << 5
	// CreateInvites allows generating invite links.
	CreateInvites Permission = 1 << 6
	// ViewAuditLog allows viewing the server audit log.
	ViewAuditLog Permission = 1 << 7
	// ManageWebhooks allows creating, editing, and deleting webhooks.
	ManageWebhooks Permission = 1 << 8
	// ManageEmoji allows uploading and deleting custom emoji.
	ManageEmoji Permission = 1 << 9

	// --- Membership (bits 10-15) ---

	// KickMembers allows removing members from the server.
	KickMembers Permission = 1 << 10
	// BanMembers allows banning members from the server.
	BanMembers Permission = 1 << 11
	// TimeoutMembers allows temporarily muting members.
	TimeoutMembers Permission = 1 << 12
	// ChangeNicknames allows changing one's own nickname.
	ChangeNicknames Permission = 1 << 13
	// ManageNicknames allows changing other members' nicknames.
	ManageNicknames Permission = 1 << 14
	// AssignRoles allows assigning and removing roles from members.
	AssignRoles Permission = 1 << 15

	// --- Text (bits 16-26) ---

	// SendMessages allows sending messages in text channels.
	SendMessages Permission = 1 << 16
	// SendMessagesThreads allows sending messages in threads.
	SendMessagesThreads Permission = 1 << 17
	// CreateThreads allows creating new threads.
	CreateThreads Permission = 1 << 18
	// EmbedLinks allows link previews to be displayed.
	EmbedLinks Permission = 1 << 19
	// AttachFiles allows uploading files in messages.
	AttachFiles Permission = 1 << 20
	// AddReactions allows adding reactions to messages.
	AddReactions Permission = 1 << 21
	// MentionEveryone allows using @everyone and @here mentions.
	MentionEveryone Permission = 1 << 22
	// ManageMessages allows deleting and pinning other members' messages.
	ManageMessages Permission = 1 << 23
	// ReadMessageHistory allows reading message history in channels.
	ReadMessageHistory Permission = 1 << 24
	// UseCommands allows using slash commands and bot commands.
	UseCommands Permission = 1 << 25
	// UseExternalEmoji allows using emoji from other servers.
	UseExternalEmoji Permission = 1 << 26

	// --- Voice (bits 27-35) ---

	// VoiceConnect allows joining voice channels.
	VoiceConnect Permission = 1 << 27
	// VoiceSpeak allows speaking in voice channels.
	VoiceSpeak Permission = 1 << 28
	// VoiceVideo allows sharing video in voice channels.
	VoiceVideo Permission = 1 << 29
	// VoiceScreenshare allows screen sharing in voice channels.
	VoiceScreenshare Permission = 1 << 30
	// VoicePTT allows using push-to-talk in voice channels.
	VoicePTT Permission = 1 << 31
	// VoiceMuteMembers allows server-muting other members in voice.
	VoiceMuteMembers Permission = 1 << 32
	// VoiceDeafenMembers allows server-deafening other members in voice.
	VoiceDeafenMembers Permission = 1 << 33
	// VoiceMoveMembers allows moving members between voice channels.
	VoiceMoveMembers Permission = 1 << 34
	// VoicePriority allows priority speaker mode in voice channels.
	VoicePriority Permission = 1 << 35

	// --- Moderation (bits 36-39) ---

	// ManageAutomod allows configuring auto-moderation rules.
	ManageAutomod Permission = 1 << 36
	// ViewReports allows viewing user reports.
	ViewReports Permission = 1 << 37
	// ManageReports allows resolving and dismissing reports.
	ManageReports Permission = 1 << 38
	// SlowmodeExempt allows bypassing channel slow mode.
	SlowmodeExempt Permission = 1 << 39

	// AllPermissions has every permission bit set (bits 0-39).
	AllPermissions Permission = (1 << 40) - 1
)

// Has reports whether the bitfield contains all bits in check.
func (p Permission) Has(check Permission) bool {
	return p&check == check
}

// Add returns a new Permission with the given bits granted (bitwise OR).
func (p Permission) Add(perms Permission) Permission {
	return p | perms
}

// Remove returns a new Permission with the given bits revoked (bitwise AND-NOT).
func (p Permission) Remove(perms Permission) Permission {
	return p &^ perms
}
