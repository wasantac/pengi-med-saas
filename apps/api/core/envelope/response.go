package envelope

import (
	"fmt"
	"net/http"
	core_errors "pengi-med-saas/core/errors"
	"reflect"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func New(code int, message string, data interface{}) Response {
	if message == "" {
		message = http.StatusText(code)
	}

	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func SuccessResponse(data any, message string) Response {
	return Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(code int, message string, data core_errors.AppError) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (r Response) Unwrap() error {
	if r.Code > 399 {
		if reflect.TypeOf(r.Data) != reflect.TypeFor[core_errors.AppError]() {
			panic("Unwrap: data must be of type AppError")
		}
		return fmt.Errorf("description: %s, metadata: %s", r.Message, r.Data)
	}
	return nil
}

func (r Response) Error() string {
	return fmt.Sprintf("description: %s, metadata: %s", r.Message, r.Data)
}
