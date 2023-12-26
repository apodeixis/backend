package types

type ErrorCode = string

const (
	ErrNotFoundUser                   ErrorCode = "err_not_found_user"
	ErrNotFoundRefreshToken           ErrorCode = "err_not_found_refresh_token"
	ErrNotFoundState                  ErrorCode = "err_not_found_state"
	ErrNotFoundPasswordRecoveryToken  ErrorCode = "err_not_found_password_recovery_token"
	ErrNotFoundEmailVerificationToken ErrorCode = "err_not_found_email_verification_token"

	ErrUnprocessableEntityNotOAuth2User ErrorCode = "err_unprocessable_entity_not_oauth2_user"
	ErrUnprocessableEntityOAuth2User    ErrorCode = "err_unprocessable_entity_oauth2_user"
	ErrEmailAlreadyVerified             ErrorCode = "err_email_already_verified"

	ErrInsufficientPermissions ErrorCode = "err_insufficient_permissions"
	ErrInvalidJWTToken         ErrorCode = "err_invalid_jwt_token"
	ErrInvalidPassword         ErrorCode = "err_invalid_password"
	ErrInvalidEmail            ErrorCode = "err_invalid_email"

	ErrInvalidOAuth2State ErrorCode = "err_invalid_oauth2_state"
	ErrExpiredOAuth2State ErrorCode = "err_expired_oauth2_state"

	ErrInvalidOrExpiredOAuth2Code ErrorCode = "err_invalid_or_expired_oauth2_code"

	ErrInvalidPasswordRecoveryToken ErrorCode = "err_invalid_password_recovery_token"
	ErrExpiredPasswordRecoveryToken ErrorCode = "err_expired_password_recovery_token"

	ErrInvalidEmailVerificationToken ErrorCode = "err_invalid_email_verification_token"
	ErrExpiredEmailVerificationToken ErrorCode = "err_expired_email_verification_token"

	ErrConflictUserEmail ErrorCode = "err_conflict_user_email"
)
