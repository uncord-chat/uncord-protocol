package models

// SearchMessageHit is the API response shape for a single search result.
type SearchMessageHit struct {
	ID         string   `json:"id"`
	ChannelID  string   `json:"channel_id"`
	AuthorID   string   `json:"author_id"`
	Content    string   `json:"content"`
	CreatedAt  int64    `json:"created_at"`
	Highlights []string `json:"highlights,omitempty"`
}

// SearchResponse is the top-level search response returned inside the data envelope.
type SearchResponse struct {
	TotalCount int                `json:"total_count"`
	Page       int                `json:"page"`
	PerPage    int                `json:"per_page"`
	Hits       []SearchMessageHit `json:"hits"`
}
