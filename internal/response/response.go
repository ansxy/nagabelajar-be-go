package response

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"
)

type JSONResponse struct {
	Success bool                `json:"success"`
	Paging  *PaginationResponse `json:"paging"`
	Data    interface{}         `json:"data"`
	Error   *ErrorResponse      `json:"error"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	Page      int    `json:"page"`
	PerPage   int    `json:"per_page"`
	Count     int64  `json:"count"`
	PageCount int64  `json:"page_count"`
	Next      string `json:"next"`
	Previous  string `json:"previous"`
}

func Success(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JSONResponse{
		Success: true,
		Data:    data,
	})
}

func Pagination(w http.ResponseWriter, list interface{}, page, perPage int, count int64) {
	var paging *PaginationResponse
	total := math.Ceil(float64(count) / float64(perPage))

	if page > 0 {
		paging = &PaginationResponse{
			Page:      page,
			PerPage:   perPage,
			Count:     count,
			PageCount: int64(total),
			Next:      fmt.Sprint(page < int(total)),
			Previous:  fmt.Sprint(page > 1),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JSONResponse{
		Success: true,
		Paging:  paging,
		Data:    list,
		Error:   nil,
	})
}

func Error(w http.ResponseWriter, err error) {

	//Check errror if ccontais "invalid syntax"
	if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrSyntax {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			Code:     constant.DefaultBadRequestError,
			Message:  constant.ErrorMessageMap[constant.DefaultBadRequestError],
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultBadRequestError],
		})
	}

	if _, ok := err.(*custom_error.CustomError); !ok {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			Code:     constant.DefaultUnhandleError,
			Message:  err.Error(),
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultUnhandleError],
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.(*custom_error.CustomError).ErrorContext.HTTPCode)
	json.NewEncoder(w).Encode(JSONResponse{
		Success: false,
		Error: &ErrorResponse{
			Code:    err.(*custom_error.CustomError).ErrorContext.Code,
			Status:  err.(*custom_error.CustomError).ErrorContext.HTTPCode,
			Message: err.(*custom_error.CustomError).ErrorContext.Message,
		},
	})
}

func UnauthorizedError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(JSONResponse{
		Success: false,
		Error: &ErrorResponse{
			Code:    constant.DefaultUnauthorizedError,
			Status:  http.StatusUnauthorized,
			Message: constant.ErrorMessageMap[constant.DefaultUnauthorizedError],
		},
	})
}
