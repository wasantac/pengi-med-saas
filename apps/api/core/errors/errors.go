package core_errors

type AppError struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func NewAppError(code string, message string) AppError {
	return AppError{
		ErrorCode:    code,
		ErrorMessage: message,
	}
}
