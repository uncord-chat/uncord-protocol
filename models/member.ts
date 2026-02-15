/** Member status constants matching the database CHECK constraint. */
export const MemberStatus = {
  Pending: "pending",
  Active: "active",
  TimedOut: "timed_out",
} as const;
export type MemberStatus = (typeof MemberStatus)[keyof typeof MemberStatus];

/** Subset of user fields exposed in member listings and gateway events. */
export interface MemberUser {
  id: string;
  username: string;
  display_name: string | null;
  avatar_key: string | null;
}

/** Server member returned by the API. */
export interface Member {
  user: MemberUser;
  nickname: string | null;
  joined_at: string;
  roles: string[];
  status: string;
  timeout_until: string | null;
}

/** Server ban returned by the API. */
export interface Ban {
  user: MemberUser;
  reason: string | null;
  banned_by: string | null;
  expires_at: string | null;
  created_at: string;
}

/** Body for PATCH /api/v1/server/members/:userID or /@me. */
export interface UpdateMemberRequest {
  nickname?: string | null;
}

/** Body for PUT /api/v1/server/members/:userID/timeout. */
export interface TimeoutMemberRequest {
  until: string;
}

/** Body for PUT /api/v1/server/bans/:userID. */
export interface BanMemberRequest {
  reason?: string | null;
  expires_at?: string | null;
}
