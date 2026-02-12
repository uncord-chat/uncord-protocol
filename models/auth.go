package models

// RegisterRequest is the JSON body for POST /api/v1/auth/register.
type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRequest is the JSON body for POST /api/v1/auth/login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RefreshRequest is the JSON body for POST /api/v1/auth/refresh.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// VerifyEmailRequest is the JSON body for POST /api/v1/auth/verify-email.
type VerifyEmailRequest struct {
	Token string `json:"token"`
}

// AuthResponse is the response body for register and login endpoints.
type AuthResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenPairResponse is the response body for the refresh endpoint.
type TokenPairResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// MessageResponse is a generic response containing a status message.
type MessageResponse struct {
	Message string `json:"message"`
}
