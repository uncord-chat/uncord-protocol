package models

// Message represents a message in a channel returned by the API.
type Message struct {
	ID          string     `json:"id"`
	ChannelID   string     `json:"channel_id"`
	Author      MemberUser `json:"author"`
	Content     string     `json:"content"`
	Attachments []string   `json:"attachments"`
	ReplyToID   *string    `json:"reply_to_id"`
	Pinned      bool       `json:"pinned"`
	EditedAt    *string    `json:"edited_at"`
	CreatedAt   string     `json:"created_at"`
}

// CreateMessageRequest is the body for POST /api/v1/channels/:channelID/messages.
type CreateMessageRequest struct {
	Content   string  `json:"content"`
	ReplyToID *string `json:"reply_to_id"`
}

// UpdateMessageRequest is the body for PATCH /api/v1/messages/:messageID.
type UpdateMessageRequest struct {
	Content string `json:"content"`
}
