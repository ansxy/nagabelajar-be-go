package custom_error

import "github.com/ansxy/nagabelajar-be-go/pkg/constant"

type ErrorContext struct {
	Code     int
	Message  string
	HTTPCode int
	Function string
}

type CustomError struct {
	ErrorContext *ErrorContext
}

func (c *CustomError) Error() string {
	if c.ErrorContext.HTTPCode == 0 {
		c.ErrorContext.HTTPCode = constant.ErrorCodeResponseMap[constant.DefaultUnhandleError]
	}
	return constant.ErrorMessageMap[constant.DefaultUnhandleError]
}

func SetCostumeError(cte *ErrorContext) *CustomError {
	return &CustomError{ErrorContext: cte}
}
