/** Machine-readable error codes returned in all Uncord API error responses. */
export const ErrorCode = {
  // Authentication & Authorization
  Unauthorised: "UNAUTHORISED",
  TokenExpired: "TOKEN_EXPIRED",
  TokenReused: "TOKEN_REUSED",
  InvalidCredentials: "INVALID_CREDENTIALS",
  InvalidToken: "INVALID_TOKEN",
  MFARequired: "MFA_REQUIRED",
  InvalidMFACode: "INVALID_MFA_CODE",
  MFANotEnabled: "MFA_NOT_ENABLED",
  MFAAlreadyEnabled: "MFA_ALREADY_ENABLED",
  MissingPermissions: "MISSING_PERMISSIONS",
  OwnerOnly: "OWNER_ONLY",
  RoleHierarchy: "ROLE_HIERARCHY",
  Banned: "BANNED",
  UserTimedOut: "USER_TIMED_OUT",

  // Validation
  InvalidBody: "INVALID_BODY",
  InvalidEmail: "INVALID_EMAIL",
  InvalidUsername: "INVALID_USERNAME",
  InvalidPassword: "INVALID_PASSWORD",
  DisposableEmail: "DISPOSABLE_EMAIL",
  ValidationError: "VALIDATION_ERROR",
  PayloadTooLarge: "PAYLOAD_TOO_LARGE",
  ImageDimensionsExceeded: "IMAGE_DIMENSIONS_EXCEEDED",
  InvalidEmojiName: "INVALID_EMOJI_NAME",
  MissingChannelID: "MISSING_CHANNEL_ID",
  InvalidChannelID: "INVALID_CHANNEL_ID",

  // Not Found
  NotFound: "NOT_FOUND",
  UnknownChannel: "UNKNOWN_CHANNEL",
  UnknownRole: "UNKNOWN_ROLE",
  UnknownMessage: "UNKNOWN_MESSAGE",
  UnknownInvite: "UNKNOWN_INVITE",
  UnknownCategory: "UNKNOWN_CATEGORY",

  // Limits
  MaxChannelsReached: "MAX_CHANNELS_REACHED",
  MaxCategoriesReached: "MAX_CATEGORIES_REACHED",

  // Conflict
  AlreadyExists: "ALREADY_EXISTS",
  AlreadyMember: "ALREADY_MEMBER",

  // Rate Limiting
  RateLimited: "RATE_LIMITED",

  // Server Errors
  InternalError: "INTERNAL_ERROR",
  SearchUnavailable: "SEARCH_UNAVAILABLE",
  ServiceUnavailable: "SERVICE_UNAVAILABLE",
} as const;

export type ErrorCode = (typeof ErrorCode)[keyof typeof ErrorCode];
