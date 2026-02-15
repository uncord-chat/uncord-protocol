package models

// SetOverrideRequest is the body for PUT /api/v1/channels/:channelID/overrides/:targetID.
type SetOverrideRequest struct {
	Type  string `json:"type"`  // "role" or "user"
	Allow int64  `json:"allow"`
	Deny  int64  `json:"deny"`
}

// PermissionOverride represents a channel-level permission override returned by the API.
type PermissionOverride struct {
	ID        string `json:"id"`
	Type      string `json:"type"`      // "role" or "user"
	TargetID  string `json:"target_id"` // the role or user ID
	Allow     int64  `json:"allow"`
	Deny      int64  `json:"deny"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ResolvedPermissions is the response for GET /api/v1/channels/:channelID/permissions/@me.
type ResolvedPermissions struct {
	Permissions int64 `json:"permissions"`
}
