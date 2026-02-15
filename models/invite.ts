/** Server invite returned by the API. */
export interface Invite {
  id: string;
  code: string;
  channel_id: string;
  creator_id: string;
  max_uses: number | null;
  use_count: number;
  max_age_seconds: number | null;
  expires_at: string | null;
  created_at: string;
}

/** Body for POST /api/v1/server/invites. */
export interface CreateInviteRequest {
  channel_id: string;
  max_uses?: number | null;
  max_age_seconds?: number | null;
}

/** Body for POST /api/v1/onboarding/accept. */
export interface AcceptOnboardingRequest {
  accepted_rules: boolean;
}
