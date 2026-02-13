package models

// ServerConfig represents the server configuration returned by the API.
type ServerConfig struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	IconKey     *string `json:"icon_key"`
	BannerKey   *string `json:"banner_key"`
	OwnerID     string  `json:"owner_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// UpdateServerConfigRequest is the body for PATCH /api/v1/server.
type UpdateServerConfigRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IconKey     *string `json:"icon_key"`
	BannerKey   *string `json:"banner_key"`
}
