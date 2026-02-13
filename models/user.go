// Package models defines the shared entity and API types used by the Uncord
// server and clients. Both Go and TypeScript implementations import these
// types as the single source of truth for API request and response shapes.
package models

// User represents the public user profile returned by the API.
type User struct {
	ID            string  `json:"id"`
	Email         string  `json:"email"`
	Username      string  `json:"username"`
	DisplayName   *string `json:"display_name"`
	AvatarKey     *string `json:"avatar_key"`
	EmailVerified bool    `json:"email_verified"`
}

// UpdateUserRequest is the body for PATCH /api/v1/users/@me.
type UpdateUserRequest struct {
	DisplayName *string `json:"display_name"`
	AvatarKey   *string `json:"avatar_key"`
}
