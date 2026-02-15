/** Body for PUT /api/v1/channels/:channelID/overrides/:targetID. */
export interface SetOverrideRequest {
  type: "role" | "user";
  allow: number;
  deny: number;
}

/** Channel-level permission override returned by the API. */
export interface PermissionOverride {
  id: string;
  type: "role" | "user";
  target_id: string;
  allow: number;
  deny: number;
  created_at: string;
  updated_at: string;
}

/** Response for GET /api/v1/channels/:channelID/permissions/@me. */
export interface ResolvedPermissions {
  permissions: number;
}
