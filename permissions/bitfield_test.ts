import { strictEqual } from "node:assert";
import { describe, test } from "node:test";
import * as permissions from "./bitfield.ts";

describe("has", () => {
  const tests: { name: string; base: bigint; check: bigint; want: boolean }[] = [
    { name: "single bit set", base: permissions.ViewChannels, check: permissions.ViewChannels, want: true },
    { name: "single bit unset", base: 0n, check: permissions.ViewChannels, want: false },
    {
      name: "multiple bits all set",
      base: permissions.ViewChannels | permissions.SendMessages,
      check: permissions.ViewChannels | permissions.SendMessages,
      want: true,
    },
    {
      name: "multiple bits partial",
      base: permissions.ViewChannels,
      check: permissions.ViewChannels | permissions.SendMessages,
      want: false,
    },
    { name: "zero check always true", base: 0n, check: 0n, want: true },
    {
      name: "all permissions has every bit",
      base: permissions.AllPermissions,
      check: permissions.ManageServer,
      want: true,
    },
    {
      name: "all permissions has voice",
      base: permissions.AllPermissions,
      check: permissions.VoiceConnect | permissions.VoiceSpeak,
      want: true,
    },
    {
      name: "all permissions has moderation",
      base: permissions.AllPermissions,
      check: permissions.ManageAutomod | permissions.SlowmodeExempt,
      want: true,
    },
  ];

  for (const tt of tests) {
    test(tt.name, () => {
      strictEqual(permissions.has(tt.base, tt.check), tt.want);
    });
  }
});

describe("add", () => {
  const tests: { name: string; base: bigint; add: bigint; want: bigint }[] = [
    { name: "add to zero", base: 0n, add: permissions.ViewChannels, want: permissions.ViewChannels },
    {
      name: "add existing is idempotent",
      base: permissions.ViewChannels,
      add: permissions.ViewChannels,
      want: permissions.ViewChannels,
    },
    {
      name: "combine two",
      base: permissions.ViewChannels,
      add: permissions.SendMessages,
      want: permissions.ViewChannels | permissions.SendMessages,
    },
    {
      name: "add to complex",
      base: permissions.ViewChannels | permissions.ManageServer,
      add: permissions.BanMembers,
      want: permissions.ViewChannels | permissions.ManageServer | permissions.BanMembers,
    },
  ];

  for (const tt of tests) {
    test(tt.name, () => {
      strictEqual(permissions.add(tt.base, tt.add), tt.want);
    });
  }
});

describe("remove", () => {
  const tests: { name: string; base: bigint; remove: bigint; want: bigint }[] = [
    {
      name: "remove present bit",
      base: permissions.ViewChannels | permissions.SendMessages,
      remove: permissions.SendMessages,
      want: permissions.ViewChannels,
    },
    {
      name: "remove absent bit is no-op",
      base: permissions.ViewChannels,
      remove: permissions.SendMessages,
      want: permissions.ViewChannels,
    },
    { name: "remove from zero", base: 0n, remove: permissions.ViewChannels, want: 0n },
    {
      name: "remove all",
      base: permissions.ViewChannels | permissions.SendMessages,
      remove: permissions.ViewChannels | permissions.SendMessages,
      want: 0n,
    },
  ];

  for (const tt of tests) {
    test(tt.name, () => {
      strictEqual(permissions.remove(tt.base, tt.remove), tt.want);
    });
  }
});

describe("AllPermissions", () => {
  const all = [
    // General
    permissions.ViewChannels,
    permissions.ManageChannels,
    permissions.ManageCategories,
    permissions.ManageRoles,
    permissions.ManageServer,
    permissions.ManageInvites,
    permissions.CreateInvites,
    permissions.ViewAuditLog,
    permissions.ManageWebhooks,
    permissions.ManageEmoji,
    // Membership
    permissions.KickMembers,
    permissions.BanMembers,
    permissions.TimeoutMembers,
    permissions.ChangeNicknames,
    permissions.ManageNicknames,
    permissions.AssignRoles,
    // Text
    permissions.SendMessages,
    permissions.SendMessagesThreads,
    permissions.CreateThreads,
    permissions.EmbedLinks,
    permissions.AttachFiles,
    permissions.AddReactions,
    permissions.MentionEveryone,
    permissions.ManageMessages,
    permissions.ReadMessageHistory,
    permissions.UseCommands,
    permissions.UseExternalEmoji,
    // Voice
    permissions.VoiceConnect,
    permissions.VoiceSpeak,
    permissions.VoiceVideo,
    permissions.VoiceScreenshare,
    permissions.VoicePTT,
    permissions.VoiceMuteMembers,
    permissions.VoiceDeafenMembers,
    permissions.VoiceMoveMembers,
    permissions.VoicePriority,
    // Moderation
    permissions.ManageAutomod,
    permissions.ViewReports,
    permissions.ManageReports,
    permissions.SlowmodeExempt,
  ];

  test("covers every bit", () => {
    for (const perm of all) {
      strictEqual(permissions.has(permissions.AllPermissions, perm), true, `AllPermissions missing bit ${perm}`);
    }
  });

  test("union of all constants equals AllPermissions", () => {
    let union = 0n;
    for (const perm of all) {
      union = permissions.add(union, perm);
    }
    strictEqual(union, permissions.AllPermissions);
  });
});

test("bits are contiguous", () => {
  for (let i = 0; i < 40; i++) {
    const bit = 1n << BigInt(i);
    strictEqual(permissions.has(permissions.AllPermissions, bit), true, `AllPermissions missing bit ${i}`);
  }
});

test("no overlapping bits", () => {
  const perms: { name: string; perm: bigint }[] = [
    { name: "ViewChannels", perm: permissions.ViewChannels },
    { name: "ManageChannels", perm: permissions.ManageChannels },
    { name: "ManageCategories", perm: permissions.ManageCategories },
    { name: "ManageRoles", perm: permissions.ManageRoles },
    { name: "ManageServer", perm: permissions.ManageServer },
    { name: "ManageInvites", perm: permissions.ManageInvites },
    { name: "CreateInvites", perm: permissions.CreateInvites },
    { name: "ViewAuditLog", perm: permissions.ViewAuditLog },
    { name: "ManageWebhooks", perm: permissions.ManageWebhooks },
    { name: "ManageEmoji", perm: permissions.ManageEmoji },
    { name: "KickMembers", perm: permissions.KickMembers },
    { name: "BanMembers", perm: permissions.BanMembers },
    { name: "TimeoutMembers", perm: permissions.TimeoutMembers },
    { name: "ChangeNicknames", perm: permissions.ChangeNicknames },
    { name: "ManageNicknames", perm: permissions.ManageNicknames },
    { name: "AssignRoles", perm: permissions.AssignRoles },
    { name: "SendMessages", perm: permissions.SendMessages },
    { name: "SendMessagesThreads", perm: permissions.SendMessagesThreads },
    { name: "CreateThreads", perm: permissions.CreateThreads },
    { name: "EmbedLinks", perm: permissions.EmbedLinks },
    { name: "AttachFiles", perm: permissions.AttachFiles },
    { name: "AddReactions", perm: permissions.AddReactions },
    { name: "MentionEveryone", perm: permissions.MentionEveryone },
    { name: "ManageMessages", perm: permissions.ManageMessages },
    { name: "ReadMessageHistory", perm: permissions.ReadMessageHistory },
    { name: "UseCommands", perm: permissions.UseCommands },
    { name: "UseExternalEmoji", perm: permissions.UseExternalEmoji },
    { name: "VoiceConnect", perm: permissions.VoiceConnect },
    { name: "VoiceSpeak", perm: permissions.VoiceSpeak },
    { name: "VoiceVideo", perm: permissions.VoiceVideo },
    { name: "VoiceScreenshare", perm: permissions.VoiceScreenshare },
    { name: "VoicePTT", perm: permissions.VoicePTT },
    { name: "VoiceMuteMembers", perm: permissions.VoiceMuteMembers },
    { name: "VoiceDeafenMembers", perm: permissions.VoiceDeafenMembers },
    { name: "VoiceMoveMembers", perm: permissions.VoiceMoveMembers },
    { name: "VoicePriority", perm: permissions.VoicePriority },
    { name: "ManageAutomod", perm: permissions.ManageAutomod },
    { name: "ViewReports", perm: permissions.ViewReports },
    { name: "ManageReports", perm: permissions.ManageReports },
    { name: "SlowmodeExempt", perm: permissions.SlowmodeExempt },
  ];

  for (let i = 0; i < perms.length; i++) {
    for (let j = i + 1; j < perms.length; j++) {
      strictEqual((perms[i].perm & perms[j].perm) === 0n, true, `${perms[i].name} and ${perms[j].name} share bits`);
    }
  }
});

test("deny wins pattern", () => {
  const base = permissions.ViewChannels | permissions.SendMessages | permissions.AddReactions;
  const roleAllow = permissions.ManageMessages;
  const roleDeny = permissions.SendMessages;

  const result = permissions.remove(permissions.add(base, roleAllow), roleDeny);

  strictEqual(permissions.has(result, permissions.SendMessages), false, "SendMessages should be denied");
  strictEqual(permissions.has(result, permissions.ViewChannels), true, "ViewChannels should remain");
  strictEqual(permissions.has(result, permissions.ManageMessages), true, "ManageMessages should be added");
  strictEqual(permissions.has(result, permissions.AddReactions), true, "AddReactions should remain untouched");
});
