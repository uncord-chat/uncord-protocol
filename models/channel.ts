/** Channel type constants matching the database CHECK constraint. */
export const ChannelType = {
  Text: "text",
  Voice: "voice",
  Announcement: "announcement",
  Forum: "forum",
  Stage: "stage",
} as const;
export type ChannelType = (typeof ChannelType)[keyof typeof ChannelType];

/** Channel returned by the API. */
export interface Channel {
  id: string;
  category_id: string | null;
  name: string;
  type: string;
  topic: string;
  position: number;
  slowmode_seconds: number;
  nsfw: boolean;
  created_at: string;
  updated_at: string;
}

/** Body for POST /api/v1/server/channels. */
export interface CreateChannelRequest {
  name: string;
  type?: string | null;
  category_id?: string | null;
  topic?: string | null;
  slowmode_seconds?: number | null;
  nsfw?: boolean | null;
}

/** Body for PATCH /api/v1/channels/:channelID. */
export interface UpdateChannelRequest {
  name?: string | null;
  category_id?: string | null;
  topic?: string | null;
  position?: number | null;
  slowmode_seconds?: number | null;
  nsfw?: boolean | null;
}
