package utils

import "fmt"

type AppError struct {
	Code int
	Message string
	Err error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s : %v", e.Message, e.Err)
	}
	return e.Message
}

func ErrorHandler(code int, msg string, err error) *AppError {
	return &AppError{ Code: code, Message: msg, Err: err}
}