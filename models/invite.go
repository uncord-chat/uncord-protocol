package models

// Invite represents a server invite in API responses.
type Invite struct {
	ID            string  `json:"id"`
	Code          string  `json:"code"`
	ChannelID     string  `json:"channel_id"`
	CreatorID     string  `json:"creator_id"`
	MaxUses       *int    `json:"max_uses"`
	UseCount      int     `json:"use_count"`
	MaxAgeSeconds *int    `json:"max_age_seconds"`
	ExpiresAt     *string `json:"expires_at,omitempty"`
	CreatedAt     string  `json:"created_at"`
}

// CreateInviteRequest is the body for POST /api/v1/server/invites.
type CreateInviteRequest struct {
	ChannelID     string `json:"channel_id"`
	MaxUses       *int   `json:"max_uses"`
	MaxAgeSeconds *int   `json:"max_age_seconds"`
}

// AcceptOnboardingRequest is the body for POST /api/v1/onboarding/accept.
type AcceptOnboardingRequest struct {
	AcceptedRules bool `json:"accepted_rules"`
}
