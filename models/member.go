package models

// Member status constants matching the database CHECK constraint.
const (
	MemberStatusPending  = "pending"
	MemberStatusActive   = "active"
	MemberStatusTimedOut = "timed_out"
)

// MemberUser is the subset of user fields exposed in member listings and gateway events.
type MemberUser struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	DisplayName *string `json:"display_name"`
	AvatarKey   *string `json:"avatar_key"`
}

// Member represents a server member in API responses.
type Member struct {
	User         MemberUser `json:"user"`
	Nickname     *string    `json:"nickname"`
	JoinedAt     string     `json:"joined_at"`
	Roles        []string   `json:"roles"`
	Status       string     `json:"status"`
	TimeoutUntil *string    `json:"timeout_until"`
}

// Ban represents a server ban in API responses.
type Ban struct {
	User      MemberUser `json:"user"`
	Reason    *string    `json:"reason"`
	BannedBy  *string    `json:"banned_by"`
	ExpiresAt *string    `json:"expires_at"`
	CreatedAt string     `json:"created_at"`
}

// UpdateMemberRequest is the body for PATCH /api/v1/server/members/:userID or /@me.
type UpdateMemberRequest struct {
	Nickname *string `json:"nickname"`
}

// TimeoutMemberRequest is the body for PUT /api/v1/server/members/:userID/timeout.
type TimeoutMemberRequest struct {
	Until string `json:"until"`
}

// BanMemberRequest is the body for PUT /api/v1/server/bans/:userID.
type BanMemberRequest struct {
	Reason    *string `json:"reason"`
	ExpiresAt *string `json:"expires_at"`
}
