package models

// Role represents a server role returned by the API.
type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Colour      int    `json:"colour"`
	Position    int    `json:"position"`
	Hoist       bool   `json:"hoist"`
	Permissions int64  `json:"permissions"`
	IsEveryone  bool   `json:"is_everyone"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// CreateRoleRequest is the body for POST /api/v1/server/roles.
type CreateRoleRequest struct {
	Name        string `json:"name"`
	Colour      *int   `json:"colour"`
	Permissions *int64 `json:"permissions"`
	Hoist       *bool  `json:"hoist"`
}

// UpdateRoleRequest is the body for PATCH /api/v1/server/roles/:roleID.
type UpdateRoleRequest struct {
	Name        *string `json:"name"`
	Colour      *int    `json:"colour"`
	Position    *int    `json:"position"`
	Permissions *int64  `json:"permissions"`
	Hoist       *bool   `json:"hoist"`
}
