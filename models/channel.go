package models

// Channel type constants matching the database CHECK constraint.
const (
	ChannelTypeText         = "text"
	ChannelTypeVoice        = "voice"
	ChannelTypeAnnouncement = "announcement"
	ChannelTypeForum        = "forum"
	ChannelTypeStage        = "stage"
)

// Channel represents a channel returned by the API.
type Channel struct {
	ID              string  `json:"id"`
	CategoryID      *string `json:"category_id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Topic           string  `json:"topic"`
	Position        int     `json:"position"`
	SlowmodeSeconds int     `json:"slowmode_seconds"`
	NSFW            bool    `json:"nsfw"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

// CreateChannelRequest is the body for POST /api/v1/server/channels.
type CreateChannelRequest struct {
	Name            string  `json:"name"`
	Type            *string `json:"type"`
	CategoryID      *string `json:"category_id"`
	Topic           *string `json:"topic"`
	SlowmodeSeconds *int    `json:"slowmode_seconds"`
	NSFW            *bool   `json:"nsfw"`
}

// UpdateChannelRequest is the body for PATCH /api/v1/channels/:channelID.
type UpdateChannelRequest struct {
	Name            *string `json:"name"`
	CategoryID      *string `json:"category_id"`
	Topic           *string `json:"topic"`
	Position        *int    `json:"position"`
	SlowmodeSeconds *int    `json:"slowmode_seconds"`
	NSFW            *bool   `json:"nsfw"`
}
