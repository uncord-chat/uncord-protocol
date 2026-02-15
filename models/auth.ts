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

/** Response when login succeeds but MFA verification is needed. */
export interface MFARequiredResponse {
  mfa_required: boolean;
  ticket: string;
}

/** JSON body for POST /api/v1/auth/mfa/verify. */
export interface MFAVerifyRequest {
  ticket: string;
  code: string;
}

/** JSON body for POST /api/v1/users/@me/mfa/enable. */
export interface MFAEnableRequest {
  password: string;
}

/** Response for the MFA enable endpoint with TOTP provisioning data. */
export interface MFASetupResponse {
  secret: string;
  uri: string;
}

/** JSON body for POST /api/v1/users/@me/mfa/confirm. */
export interface MFAConfirmRequest {
  code: string;
}

/** Response for the MFA confirm endpoint with recovery codes. */
export interface MFAConfirmResponse {
  recovery_codes: string[];
}

/** JSON body for POST /api/v1/users/@me/mfa/disable. */
export interface MFADisableRequest {
  password: string;
  code: string;
}

/** JSON body for POST /api/v1/users/@me/mfa/recovery-codes. */
export interface MFARegenerateCodesRequest {
  password: string;
}

/** Response for the recovery code regeneration endpoint. */
export interface MFARegenerateCodesResponse {
  recovery_codes: string[];
}

/** JSON body for POST /api/v1/auth/verify-password. */
export interface VerifyPasswordRequest {
  password: string;
}

/** JSON body for DELETE /api/v1/users/@me. */
export interface DeleteAccountRequest {
  password: string;
}
