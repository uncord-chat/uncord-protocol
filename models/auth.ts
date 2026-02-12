import type { User } from "./user.ts";

/** JSON body for POST /api/v1/auth/register. */
export interface RegisterRequest {
  email: string;
  username: string;
  password: string;
}

/** JSON body for POST /api/v1/auth/login. */
export interface LoginRequest {
  email: string;
  password: string;
}

/** JSON body for POST /api/v1/auth/refresh. */
export interface RefreshRequest {
  refresh_token: string;
}

/** JSON body for POST /api/v1/auth/verify-email. */
export interface VerifyEmailRequest {
  token: string;
}

/** Response body for register and login endpoints. */
export interface AuthResponse {
  user: User;
  access_token: string;
  refresh_token: string;
}

/** Response body for the refresh endpoint. */
export interface TokenPairResponse {
  access_token: string;
  refresh_token: string;
}

/** Generic response containing a status message. */
export interface MessageResponse {
  message: string;
}
