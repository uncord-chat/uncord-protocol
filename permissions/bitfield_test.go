package permissions

import "testing"

func TestHas(t *testing.T) {
	tests := []struct {
		name  string
		base  Permission
		check Permission
		want  bool
	}{
		{"single bit set", ViewChannels, ViewChannels, true},
		{"single bit unset", 0, ViewChannels, false},
		{"multiple bits all set", ViewChannels | SendMessages, ViewChannels | SendMessages, true},
		{"multiple bits partial", ViewChannels, ViewChannels | SendMessages, false},
		{"zero check always true", 0, 0, true},
		{"all permissions has every bit", AllPermissions, ManageServer, true},
		{"all permissions has voice", AllPermissions, VoiceConnect | VoiceSpeak, true},
		{"all permissions has moderation", AllPermissions, ManageAutomod | SlowmodeExempt, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.base.Has(tt.check); got != tt.want {
				t.Errorf("Permission(%d).Has(%d) = %v, want %v", tt.base, tt.check, got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		base Permission
		add  Permission
		want Permission
	}{
		{"add to zero", 0, ViewChannels, ViewChannels},
		{"add existing is idempotent", ViewChannels, ViewChannels, ViewChannels},
		{"combine two", ViewChannels, SendMessages, ViewChannels | SendMessages},
		{"add to complex", ViewChannels | ManageServer, BanMembers, ViewChannels | ManageServer | BanMembers},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.base.Add(tt.add); got != tt.want {
				t.Errorf("Permission(%d).Add(%d) = %d, want %d", tt.base, tt.add, got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name   string
		base   Permission
		remove Permission
		want   Permission
	}{
		{"remove present bit", ViewChannels | SendMessages, SendMessages, ViewChannels},
		{"remove absent bit is no-op", ViewChannels, SendMessages, ViewChannels},
		{"remove from zero", 0, ViewChannels, 0},
		{"remove all", ViewChannels | SendMessages, ViewChannels | SendMessages, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.base.Remove(tt.remove); got != tt.want {
				t.Errorf("Permission(%d).Remove(%d) = %d, want %d", tt.base, tt.remove, got, tt.want)
			}
		})
	}
}

func TestAllPermissionsCoversEveryBit(t *testing.T) {
	all := []Permission{
		// General
		ViewChannels, ManageChannels, ManageCategories, ManageRoles, ManageServer,
		ManageInvites, CreateInvites, ViewAuditLog, ManageWebhooks, ManageEmoji,
		// Membership
		KickMembers, BanMembers, TimeoutMembers, ChangeNicknames, ManageNicknames, AssignRoles,
		// Text
		SendMessages, SendMessagesThreads, CreateThreads, EmbedLinks, AttachFiles,
		AddReactions, MentionEveryone, ManageMessages, ReadMessageHistory, UseCommands, UseExternalEmoji,
		// Voice
		VoiceConnect, VoiceSpeak, VoiceVideo, VoiceScreenshare, VoicePTT,
		VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoicePriority,
		// Moderation
		ManageAutomod, ViewReports, ManageReports, SlowmodeExempt,
	}

	for _, perm := range all {
		if !AllPermissions.Has(perm) {
			t.Errorf("AllPermissions missing bit %d", perm)
		}
	}

	// Verify the union of all individual constants equals AllPermissions.
	var union Permission
	for _, perm := range all {
		union = union.Add(perm)
	}
	if union != AllPermissions {
		t.Errorf("union of all constants = %d, want AllPermissions = %d", union, AllPermissions)
	}
}

func TestBitsAreContiguous(t *testing.T) {
	for i := 0; i < 40; i++ {
		bit := Permission(1 << i)
		if !AllPermissions.Has(bit) {
			t.Errorf("AllPermissions missing bit %d", i)
		}
	}
}

func TestNoOverlappingBits(t *testing.T) {
	perms := []struct {
		name string
		perm Permission
	}{
		{"ViewChannels", ViewChannels},
		{"ManageChannels", ManageChannels},
		{"ManageCategories", ManageCategories},
		{"ManageRoles", ManageRoles},
		{"ManageServer", ManageServer},
		{"ManageInvites", ManageInvites},
		{"CreateInvites", CreateInvites},
		{"ViewAuditLog", ViewAuditLog},
		{"ManageWebhooks", ManageWebhooks},
		{"ManageEmoji", ManageEmoji},
		{"KickMembers", KickMembers},
		{"BanMembers", BanMembers},
		{"TimeoutMembers", TimeoutMembers},
		{"ChangeNicknames", ChangeNicknames},
		{"ManageNicknames", ManageNicknames},
		{"AssignRoles", AssignRoles},
		{"SendMessages", SendMessages},
		{"SendMessagesThreads", SendMessagesThreads},
		{"CreateThreads", CreateThreads},
		{"EmbedLinks", EmbedLinks},
		{"AttachFiles", AttachFiles},
		{"AddReactions", AddReactions},
		{"MentionEveryone", MentionEveryone},
		{"ManageMessages", ManageMessages},
		{"ReadMessageHistory", ReadMessageHistory},
		{"UseCommands", UseCommands},
		{"UseExternalEmoji", UseExternalEmoji},
		{"VoiceConnect", VoiceConnect},
		{"VoiceSpeak", VoiceSpeak},
		{"VoiceVideo", VoiceVideo},
		{"VoiceScreenshare", VoiceScreenshare},
		{"VoicePTT", VoicePTT},
		{"VoiceMuteMembers", VoiceMuteMembers},
		{"VoiceDeafenMembers", VoiceDeafenMembers},
		{"VoiceMoveMembers", VoiceMoveMembers},
		{"VoicePriority", VoicePriority},
		{"ManageAutomod", ManageAutomod},
		{"ViewReports", ViewReports},
		{"ManageReports", ManageReports},
		{"SlowmodeExempt", SlowmodeExempt},
	}

	for i := 0; i < len(perms); i++ {
		for j := i + 1; j < len(perms); j++ {
			if perms[i].perm&perms[j].perm != 0 {
				t.Errorf("%s and %s share bits", perms[i].name, perms[j].name)
			}
		}
	}
}

func TestDenyWinsPattern(t *testing.T) {
	base := ViewChannels | SendMessages | AddReactions

	roleAllow := ManageMessages
	roleDeny := SendMessages

	result := base.Add(roleAllow).Remove(roleDeny)

	if result.Has(SendMessages) {
		t.Error("SendMessages should be denied after deny override")
	}
	if !result.Has(ViewChannels) {
		t.Error("ViewChannels should remain after deny of unrelated permission")
	}
	if !result.Has(ManageMessages) {
		t.Error("ManageMessages should be added by role allow")
	}
	if !result.Has(AddReactions) {
		t.Error("AddReactions should remain untouched")
	}
}
