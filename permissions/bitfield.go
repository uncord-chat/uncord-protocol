package permissions

type Permission int64

const (
	// General (bits 0-9)
	ViewChannels     Permission = 1 << 0
	ManageChannels   Permission = 1 << 1
	ManageCategories Permission = 1 << 2
	ManageRoles      Permission = 1 << 3
	ManageServer     Permission = 1 << 4
	ManageInvites    Permission = 1 << 5
	CreateInvites    Permission = 1 << 6
	ViewAuditLog     Permission = 1 << 7
	ManageWebhooks   Permission = 1 << 8
	ManageEmoji      Permission = 1 << 9

	// Membership (bits 10-15)
	KickMembers     Permission = 1 << 10
	BanMembers      Permission = 1 << 11
	TimeoutMembers  Permission = 1 << 12
	ChangeNicknames Permission = 1 << 13
	ManageNicknames Permission = 1 << 14
	AssignRoles     Permission = 1 << 15

	// Text (bits 16-26)
	SendMessages        Permission = 1 << 16
	SendMessagesThreads Permission = 1 << 17
	CreateThreads       Permission = 1 << 18
	EmbedLinks          Permission = 1 << 19
	AttachFiles         Permission = 1 << 20
	AddReactions        Permission = 1 << 21
	MentionEveryone     Permission = 1 << 22
	ManageMessages      Permission = 1 << 23
	ReadMessageHistory  Permission = 1 << 24
	UseCommands         Permission = 1 << 25

	// Voice (bits 27-35)
	VoiceConnect       Permission = 1 << 27
	VoiceSpeak         Permission = 1 << 28
	VoiceVideo         Permission = 1 << 29
	VoiceScreenshare   Permission = 1 << 30
	VoicePTT           Permission = 1 << 31
	VoiceMuteMembers   Permission = 1 << 32
	VoiceDeafenMembers Permission = 1 << 33
	VoiceMoveMembers   Permission = 1 << 34
	VoicePriority      Permission = 1 << 35

	// Moderation (bits 36-39)
	ManageAutomod  Permission = 1 << 36
	ViewReports    Permission = 1 << 37
	ManageReports  Permission = 1 << 38
	SlowmodeExempt Permission = 1 << 39

	AllPermissions Permission = (1 << 40) - 1
)

// Has checks if the bitfield contains the given permission.
func (p Permission) Has(check Permission) bool {
	return p&check == check
}

// Add grants permissions to the bitfield.
func (p Permission) Add(perms Permission) Permission {
	return p | perms
}

// Remove revokes permissions from the bitfield.
func (p Permission) Remove(perms Permission) Permission {
	return p &^ perms
}
