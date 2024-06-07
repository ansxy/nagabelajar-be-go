package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/ansxy/nagabelajar-be-go/pkg/firebase"
)

type Middleware struct {
	FCS firebase.IFaceFCM
	UC  usecase.IFaceUsecase
}

func NewMiddleware(fcm firebase.IFaceFCM, uc usecase.IFaceUsecase) Middleware {
	return Middleware{
		FCS: fcm,
		UC:  uc,
	}
}

func (m Middleware) AuthenticatedUser() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token := r.Header.Get("Authorization")
			lenToken := 2
			s := strings.Split(token, " ")

			if len(s) != lenToken {
				response.UnauthorizedError(w)
				return
			}

			accessToken := s[1]
			resJwt, err := m.FCS.VerifiyToken(ctx, accessToken)
			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			user, _ := m.UC.FindOneUser(ctx, &request.LoginRequest{
				FirebaseID: resJwt.UID,
			})

			if user != nil {
				ctx = context.WithValue(ctx, constant.UserID, user.UserID.String())
			}

			ctx = context.WithValue(ctx, constant.FirebaseID, resJwt.UID)

			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func (m Middleware) AuthenticatedAdmin() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			token := r.Header.Get("Authorization")
			letoken := 2
			s := strings.Split(token, " ")
			if len(s) != letoken {
				response.UnauthorizedError(w)
				return
			}

			firebaseToken := s[1]
			claims, err := m.FCS.VerifiyToken(ctx, firebaseToken)
			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			queryUser := request.LoginRequest{
				FirebaseID: claims.UID,
			}

			admin, err := m.UC.FindOneUser(ctx, &queryUser)
			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			if admin.Role != constant.Role.Admin {
				response.UnauthorizedError(w)
				return
			}

			ctx = context.WithValue(ctx, constant.UserID, admin.UserID)

			ctx = context.WithValue(ctx, constant.FirebaseID, claims.UID)
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
