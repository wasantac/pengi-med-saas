package core_errors

var (
	ErrInternal AppError = NewAppError("E-INT-001", "Internal server error.")

	ErrMessagesNotFound AppError = NewAppError("E-MES-001", "Messages not found.")

	ErrCompanyNotFound AppError = NewAppError("E-COMP-001", "Company not found.")

	ErrTenantNotFound AppError = NewAppError("E-TEN-001", "Tenant not found.")

	ErrUserNotFound AppError = NewAppError("E-USR-001", "User not found.")

	// Auth Errors
	ErrAuthInvalidRequest      AppError = NewAppError("E-AUTH-001", "Invalid authentication request.")
	ErrAuthUserCreateError     AppError = NewAppError("E-AUTH-002", "Error creating user.")
	ErrAuthInvalidCredentials  AppError = NewAppError("E-AUTH-003", "Invalid credentials.")
	ErrAuthTokenGenerateError  AppError = NewAppError("E-AUTH-004", "Error generating token.")
	ErrAuthInvalidRefreshToken AppError = NewAppError("E-AUTH-005", "Invalid refresh token.")
	ErrAuthUserInvalidID       AppError = NewAppError("E-AUTH-006", "Invalid user ID.")
)
