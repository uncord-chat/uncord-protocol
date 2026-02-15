import type { MemberUser } from "./member";

/** File attached to a message. */
export interface Attachment {
  id: string;
  filename: string;
  url: string;
  size: number;
  content_type: string;
  width?: number;
  height?: number;
  thumbnail_url?: string;
}

/** Message in a channel returned by the API. */
export interface Message {
  id: string;
  channel_id: string;
  author: MemberUser;
  content: string;
  attachments: Attachment[];
  reply_to_id: string | null;
  pinned: boolean;
  edited_at: string | null;
  created_at: string;
}

/** Body for POST /api/v1/channels/:channelID/messages. */
export interface CreateMessageRequest {
  content: string;
  reply_to_id?: string | null;
  attachment_ids?: string[];
}

/** Body for PATCH /api/v1/messages/:messageID. */
export interface UpdateMessageRequest {
  content: string;
}
