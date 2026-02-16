import type { Channel } from "./channel";
import type { Member } from "./member";
import type { Role } from "./role";
import type { ServerConfig } from "./server";
import type { User } from "./user";

/** Payload for opcode 10 (Hello), sent by the server on connection. */
export interface HelloData {
  heartbeat_interval: number;
}

/** Payload for opcode 2 (Identify), sent by the client to authenticate. */
export interface IdentifyData {
  token: string;
}

/** Payload for opcode 6 (Resume), sent by the client to restore a session. */
export interface ResumeData {
  token: string;
  session_id: string;
  seq: number;
}

/** Payload for opcode 3 (PresenceUpdate), sent by the client to set its own online status. */
export interface PresenceUpdateRequest {
  status: string;
}

/** Payload for the PRESENCE_UPDATE dispatch event, broadcast when a user's presence changes. */
export interface PresenceUpdateData {
  user_id: string;
  status: string;
}

/** Snapshot of a single user's presence, included in the READY payload. */
export interface PresenceState {
  user_id: string;
  status: string;
}

/** Payload for the TYPING_START dispatch event. */
export interface TypingStartData {
  channel_id: string;
  user_id: string;
  timestamp: string;
}

/** Payload for the READY dispatch event, containing initial client state. */
export interface ReadyData {
  session_id: string;
  user: User;
  server: ServerConfig;
  channels: Channel[];
  roles: Role[];
  members: Member[];
  presences: PresenceState[];
}

/** Payload for the MESSAGE_DELETE dispatch event. */
export interface MessageDeleteData {
  id: string;
  channel_id: string;
}

/** Payload for the CHANNEL_DELETE dispatch event. */
export interface ChannelDeleteData {
  id: string;
}

/** Payload for the ROLE_DELETE dispatch event. */
export interface RoleDeleteData {
  id: string;
}

/** Payload for the MEMBER_REMOVE dispatch event. */
export interface MemberRemoveData {
  user_id: string;
}
