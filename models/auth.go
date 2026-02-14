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

// MFARequiredResponse is the response when login succeeds but MFA verification is needed. The ticket is a single-use
// opaque token that the client must present to the MFA verify endpoint.
type MFARequiredResponse struct {
	MFARequired bool   `json:"mfa_required"`
	Ticket      string `json:"ticket"`
}

// MFAVerifyRequest is the JSON body for POST /api/v1/auth/mfa/verify.
type MFAVerifyRequest struct {
	Ticket string `json:"ticket"`
	Code   string `json:"code"`
}

// MFAEnableRequest is the JSON body for POST /api/v1/users/@me/mfa/enable. The password field confirms the user's
// identity before generating TOTP credentials.
type MFAEnableRequest struct {
	Password string `json:"password"`
}

// MFASetupResponse is the response for the MFA enable endpoint, containing the TOTP secret and provisioning URI for
// authenticator app registration.
type MFASetupResponse struct {
	Secret string `json:"secret"`
	URI    string `json:"uri"`
}

// MFAConfirmRequest is the JSON body for POST /api/v1/users/@me/mfa/confirm. The code is a six-digit TOTP code from
// the user's authenticator app, proving they successfully saved the secret.
type MFAConfirmRequest struct {
	Code string `json:"code"`
}

// MFAConfirmResponse is the response for the MFA confirm endpoint, containing the one-time-use recovery codes that the
// user must store securely.
type MFAConfirmResponse struct {
	RecoveryCodes []string `json:"recovery_codes"`
}

// MFADisableRequest is the JSON body for POST /api/v1/users/@me/mfa/disable. Both password and a valid TOTP code are
// required to disable MFA.
type MFADisableRequest struct {
	Password string `json:"password"`
	Code     string `json:"code"`
}

// MFARegenerateCodesRequest is the JSON body for POST /api/v1/users/@me/mfa/recovery-codes.
type MFARegenerateCodesRequest struct {
	Password string `json:"password"`
}

// MFARegenerateCodesResponse is the response for the recovery code regeneration endpoint.
type MFARegenerateCodesResponse struct {
	RecoveryCodes []string `json:"recovery_codes"`
}

// VerifyPasswordRequest is the JSON body for POST /api/v1/auth/verify-password. The authenticated user's password is
// verified without performing any other action, allowing clients to gate sensitive workflows behind a password prompt.
type VerifyPasswordRequest struct {
	Password string `json:"password"`
}
