package constant

import "net/http"

//Default error Code

const (
	DefaultUnhandleError = 1000 + iota
	DefaultNotfoundError
	DefaultValidationError
	DefaultDuplicateError
	DefaultUnauthorizedError
	DefaultBadRequestError
	DefaultLoginError
	DefaultNotValidCertification
)

var ErrorMessageMap = map[int]string{
	DefaultUnhandleError:         "Internal Server Error",
	DefaultNotfoundError:         "Data Not Found",
	DefaultValidationError:       "Validation Error",
	DefaultDuplicateError:        "Duplicate Data",
	DefaultUnauthorizedError:     "Unauthorized",
	DefaultBadRequestError:       "Bad Request",
	DefaultLoginError:            "Email or password wrong",
	DefaultNotValidCertification: "Not Valid Certificate / Diffrent MD5 hash with the original file",
}

var ErrorCodeResponseMap = map[int]int{

	// http.StatusUnprocessableEntity: DefaultValidationError,
	// http.StatusNotFound:            DefaultNotfoundError,
	// http.StatusConflict:            DefaultDuplicateError,
	// http.StatusUnauthorized:        DefaultUnauthorizedError,
	// http.StatusBadRequest:          DefaultBadRequestError,
	// http.StatusInternalServerError: DefaultUnhandleError,

	DefaultValidationError:       http.StatusUnprocessableEntity,
	DefaultNotfoundError:         http.StatusNotFound,
	DefaultDuplicateError:        http.StatusConflict,
	DefaultUnauthorizedError:     http.StatusUnauthorized,
	DefaultBadRequestError:       http.StatusBadRequest,
	DefaultUnhandleError:         http.StatusInternalServerError,
	DefaultLoginError:            http.StatusUnauthorized,
	DefaultNotValidCertification: http.StatusUnprocessableEntity,
}
