package user

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"
	"github.com/ansxy/nagabelajar-be-go/pkg/jwt"
)

func (h *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UpsertUserRequest

	if err := h.v.ValidateStruct(r, &req); err != nil {
		response.Error(w, err)
		return
	}

	err := h.uc.RegisterUser(ctx, &req)
	if err != nil {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			Code:     constant.DefaultDuplicateError,
			Message:  constant.ErrorMessageMap[constant.DefaultDuplicateError],
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultDuplicateError],
		})
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.LoginRequest

	if err := h.v.ValidateStruct(r, &req); err != nil {
		response.Error(w, err)
		return
	}

	res, err := h.uc.Login(ctx, &req)
	if err != nil {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			Code:     constant.DefaultLoginError,
			Message:  constant.ErrorMessageMap[constant.DefaultLoginError],
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultLoginError],
		})
		response.Error(w, err)
		return
	}

	jwt, err := jwt.GenereateToken(
		&jwt.UserClaims{
			UserID: res.UserID,
		},
	)

	if err != nil {
		response.Error(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "access_token",
		Value: jwt.AccessToken,
	})

	response.Success(w, res)
}
