import type { MemberUser } from "./member";

/** Message in a channel returned by the API. */
export interface Message {
  id: string;
  channel_id: string;
  author: MemberUser;
  content: string;
  attachments: string[];
  reply_to_id: string | null;
  pinned: boolean;
  edited_at: string | null;
  created_at: string;
}

/** Body for POST /api/v1/channels/:channelID/messages. */
export interface CreateMessageRequest {
  content: string;
  reply_to_id?: string | null;
}

/** Body for PATCH /api/v1/messages/:messageID. */
export interface UpdateMessageRequest {
  content: string;
}
