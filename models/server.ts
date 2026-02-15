/** Server configuration returned by the API. */
export interface ServerConfig {
  id: string;
  name: string;
  description: string;
  icon_key: string | null;
  banner_key: string | null;
  owner_id: string;
  created_at: string;
  updated_at: string;
}

/** Body for PATCH /api/v1/server. */
export interface UpdateServerConfigRequest {
  name?: string | null;
  description?: string | null;
  icon_key?: string | null;
  banner_key?: string | null;
}
