package errors

type Code string

const (
	Unauthorized            Code = "UNAUTHORIZED"
	TokenExpired            Code = "TOKEN_EXPIRED"
	MFARequired             Code = "MFA_REQUIRED"
	MissingPermissions      Code = "MISSING_PERMISSIONS"
	OwnerOnly               Code = "OWNER_ONLY"
	RoleHierarchy           Code = "ROLE_HIERARCHY"
	Banned                  Code = "BANNED"
	UserTimedOut            Code = "USER_TIMED_OUT"
	ValidationError         Code = "VALIDATION_ERROR"
	PayloadTooLarge         Code = "PAYLOAD_TOO_LARGE"
	ImageDimensionsExceeded Code = "IMAGE_DIMENSIONS_EXCEEDED"
	InvalidEmojiName        Code = "INVALID_EMOJI_NAME"
	NotFound                Code = "NOT_FOUND"
	UnknownChannel          Code = "UNKNOWN_CHANNEL"
	UnknownRole             Code = "UNKNOWN_ROLE"
	UnknownMessage          Code = "UNKNOWN_MESSAGE"
	UnknownInvite           Code = "UNKNOWN_INVITE"
	AlreadyExists           Code = "ALREADY_EXISTS"
	AlreadyMember           Code = "ALREADY_MEMBER"
	RateLimited             Code = "RATE_LIMITED"
	InternalError           Code = "INTERNAL_ERROR"
	SearchUnavailable       Code = "SEARCH_UNAVAILABLE"
	ServiceUnavailable      Code = "SERVICE_UNAVAILABLE"
)
