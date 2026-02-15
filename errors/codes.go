// Package errors defines the structured error codes returned in all Uncord
// API error responses. Both the server and any future clients import this
// package as the single source of truth, preventing code drift.
package errors

// Code is a machine-readable error identifier included in every API error
// response. Clients should match on Code rather than on HTTP status codes or
// human-readable messages.
type Code string

const (
	// --- Authentication & Authorization ---

	// Unauthorised indicates the request lacks valid authentication.
	Unauthorised Code = "UNAUTHORISED"
	// TokenExpired indicates the access token has expired.
	TokenExpired Code = "TOKEN_EXPIRED"
	// TokenReused indicates a refresh token was presented after it had
	// already been consumed, which may signal token theft.
	TokenReused Code = "TOKEN_REUSED"
	// InvalidCredentials indicates incorrect email or password.
	InvalidCredentials Code = "INVALID_CREDENTIALS"
	// InvalidToken indicates a verification or refresh token is invalid or expired.
	InvalidToken Code = "INVALID_TOKEN"
	// MFARequired indicates the account requires multi-factor authentication.
	MFARequired Code = "MFA_REQUIRED"
	// InvalidMFACode indicates the provided MFA code is incorrect.
	InvalidMFACode Code = "INVALID_MFA_CODE"
	// MFANotEnabled indicates MFA is not enabled on the account.
	MFANotEnabled Code = "MFA_NOT_ENABLED"
	// MFAAlreadyEnabled indicates MFA is already enabled on the account.
	MFAAlreadyEnabled Code = "MFA_ALREADY_ENABLED"
	// MissingPermissions indicates the user lacks a required permission.
	MissingPermissions Code = "MISSING_PERMISSIONS"
	// OwnerOnly indicates the action is restricted to the server owner.
	OwnerOnly Code = "OWNER_ONLY"
	// RoleHierarchy indicates the action is blocked by role hierarchy rules.
	RoleHierarchy Code = "ROLE_HIERARCHY"
	// Banned indicates the user is banned from the server.
	Banned Code = "BANNED"
	// UserTimedOut indicates the user is currently timed out.
	UserTimedOut Code = "USER_TIMED_OUT"
	// AccountDeleted indicates registration is blocked because the email or username was previously used by a deleted
	// account.
	AccountDeleted Code = "ACCOUNT_DELETED"
	// ServerOwner indicates the server owner must transfer ownership before performing this action.
	ServerOwner Code = "SERVER_OWNER"

	// --- Validation ---

	// InvalidBody indicates the request body could not be parsed.
	InvalidBody Code = "INVALID_BODY"
	// InvalidEmail indicates the email address is malformed.
	InvalidEmail Code = "INVALID_EMAIL"
	// InvalidUsername indicates the username does not meet requirements.
	InvalidUsername Code = "INVALID_USERNAME"
	// InvalidPassword indicates the password does not meet requirements.
	InvalidPassword Code = "INVALID_PASSWORD"
	// DisposableEmail indicates disposable email addresses are not allowed.
	DisposableEmail Code = "DISPOSABLE_EMAIL"
	// ValidationError indicates a generic input validation failure.
	ValidationError Code = "VALIDATION_ERROR"
	// PayloadTooLarge indicates the request body exceeds the size limit.
	PayloadTooLarge Code = "PAYLOAD_TOO_LARGE"
	// ImageDimensionsExceeded indicates the image exceeds dimension limits.
	ImageDimensionsExceeded Code = "IMAGE_DIMENSIONS_EXCEEDED"
	// InvalidEmojiName indicates the emoji name is invalid.
	InvalidEmojiName Code = "INVALID_EMOJI_NAME"
	// MissingChannelID indicates the required channel ID parameter is absent.
	MissingChannelID Code = "MISSING_CHANNEL_ID"
	// InvalidChannelID indicates the channel ID parameter is malformed.
	InvalidChannelID Code = "INVALID_CHANNEL_ID"

	// --- Not Found ---

	// NotFound indicates the requested resource does not exist.
	NotFound Code = "NOT_FOUND"
	// UnknownChannel indicates the referenced channel does not exist.
	UnknownChannel Code = "UNKNOWN_CHANNEL"
	// UnknownRole indicates the referenced role does not exist.
	UnknownRole Code = "UNKNOWN_ROLE"
	// UnknownMessage indicates the referenced message does not exist.
	UnknownMessage Code = "UNKNOWN_MESSAGE"
	// UnknownInvite indicates the referenced invite does not exist.
	UnknownInvite Code = "UNKNOWN_INVITE"
	// UnknownCategory indicates the referenced category does not exist.
	UnknownCategory Code = "UNKNOWN_CATEGORY"
	// UnknownOverride indicates the referenced permission override does not exist.
	UnknownOverride Code = "UNKNOWN_OVERRIDE"
	// UnknownMember indicates the referenced member does not exist.
	UnknownMember Code = "UNKNOWN_MEMBER"
	// UnknownBan indicates the referenced ban does not exist.
	UnknownBan Code = "UNKNOWN_BAN"

	// --- Limits ---

	// MaxChannelsReached indicates the server has reached the maximum number of channels.
	MaxChannelsReached Code = "MAX_CHANNELS_REACHED"
	// MaxCategoriesReached indicates the server has reached the maximum number of categories.
	MaxCategoriesReached Code = "MAX_CATEGORIES_REACHED"
	// MaxRolesReached indicates the server has reached the maximum number of roles.
	MaxRolesReached Code = "MAX_ROLES_REACHED"

	// --- Conflict ---

	// AlreadyExists indicates the resource already exists (e.g. duplicate email).
	AlreadyExists Code = "ALREADY_EXISTS"
	// AlreadyMember indicates the user is already a server member.
	AlreadyMember Code = "ALREADY_MEMBER"

	// --- Rate Limiting ---

	// RateLimited indicates the client has exceeded the request rate limit.
	RateLimited Code = "RATE_LIMITED"

	// --- Server Errors ---

	// InternalError indicates an unexpected server-side failure.
	InternalError Code = "INTERNAL_ERROR"
	// SearchUnavailable indicates the search backend is temporarily unavailable.
	SearchUnavailable Code = "SEARCH_UNAVAILABLE"
	// ServiceUnavailable indicates the service is temporarily unavailable.
	ServiceUnavailable Code = "SERVICE_UNAVAILABLE"
)
