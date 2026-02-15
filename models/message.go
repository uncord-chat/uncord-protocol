package models

// Attachment represents a file attached to a message.
type Attachment struct {
	ID           string  `json:"id"`
	Filename     string  `json:"filename"`
	URL          string  `json:"url"`
	Size         int64   `json:"size"`
	ContentType  string  `json:"content_type"`
	Width        *int    `json:"width,omitempty"`
	Height       *int    `json:"height,omitempty"`
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`
}

// Message represents a message in a channel returned by the API.
type Message struct {
	ID          string       `json:"id"`
	ChannelID   string       `json:"channel_id"`
	Author      MemberUser   `json:"author"`
	Content     string       `json:"content"`
	Attachments []Attachment `json:"attachments"`
	ReplyToID   *string      `json:"reply_to_id"`
	Pinned      bool         `json:"pinned"`
	EditedAt    *string      `json:"edited_at"`
	CreatedAt   string       `json:"created_at"`
}

// CreateMessageRequest is the body for POST /api/v1/channels/:channelID/messages.
type CreateMessageRequest struct {
	Content       string   `json:"content"`
	ReplyToID     *string  `json:"reply_to_id"`
	AttachmentIDs []string `json:"attachment_ids"`
}

// UpdateMessageRequest is the body for PATCH /api/v1/messages/:messageID.
type UpdateMessageRequest struct {
	Content string `json:"content"`
}
