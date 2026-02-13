package models

// Category represents a channel category returned by the API.
type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Position  int    `json:"position"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateCategoryRequest is the body for POST /api/v1/server/categories.
type CreateCategoryRequest struct {
	Name string `json:"name"`
}

// UpdateCategoryRequest is the body for PATCH /api/v1/categories/:categoryID.
type UpdateCategoryRequest struct {
	Name     *string `json:"name"`
	Position *int    `json:"position"`
}
