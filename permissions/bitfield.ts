// General (bits 0-9)
export const ViewChannels = 1n << 0n;
export const ManageChannels = 1n << 1n;
export const ManageCategories = 1n << 2n;
export const ManageRoles = 1n << 3n;
export const ManageServer = 1n << 4n;
export const ManageInvites = 1n << 5n;
export const CreateInvites = 1n << 6n;
export const ViewAuditLog = 1n << 7n;
export const ManageWebhooks = 1n << 8n;
export const ManageEmoji = 1n << 9n;

// Membership (bits 10-15)
export const KickMembers = 1n << 10n;
export const BanMembers = 1n << 11n;
export const TimeoutMembers = 1n << 12n;
export const ChangeNicknames = 1n << 13n;
export const ManageNicknames = 1n << 14n;
export const AssignRoles = 1n << 15n;

// Text (bits 16-26)
export const SendMessages = 1n << 16n;
export const SendMessagesThreads = 1n << 17n;
export const CreateThreads = 1n << 18n;
export const EmbedLinks = 1n << 19n;
export const AttachFiles = 1n << 20n;
export const AddReactions = 1n << 21n;
export const MentionEveryone = 1n << 22n;
export const ManageMessages = 1n << 23n;
export const ReadMessageHistory = 1n << 24n;
export const UseCommands = 1n << 25n;

// Voice (bits 27-35)
export const VoiceConnect = 1n << 27n;
export const VoiceSpeak = 1n << 28n;
export const VoiceVideo = 1n << 29n;
export const VoiceScreenshare = 1n << 30n;
export const VoicePTT = 1n << 31n;
export const VoiceMuteMembers = 1n << 32n;
export const VoiceDeafenMembers = 1n << 33n;
export const VoiceMoveMembers = 1n << 34n;
export const VoicePriority = 1n << 35n;

// Moderation (bits 36-39)
export const ManageAutomod = 1n << 36n;
export const ViewReports = 1n << 37n;
export const ManageReports = 1n << 38n;
export const SlowmodeExempt = 1n << 39n;

export const AllPermissions = (1n << 40n) - 1n;

export function has(bitfield: bigint, check: bigint): boolean {
  return (bitfield & check) === check;
}

export function add(bitfield: bigint, perms: bigint): bigint {
  return bitfield | perms;
}

export function remove(bitfield: bigint, perms: bigint): bigint {
  return bitfield & ~perms;
}
