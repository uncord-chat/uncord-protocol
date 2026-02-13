/** Public user profile returned by the API. */
export interface User {
  id: string;
  email: string;
  username: string;
  display_name: string | null;
  avatar_key: string | null;
  email_verified: boolean;
}

/** Body for PATCH /api/v1/users/@me. */
export interface UpdateUserRequest {
  display_name?: string | null;
  avatar_key?: string | null;
}
