package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/ansxy/nagabelajar-be-go/internal/repository"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/ansxy/nagabelajar-be-go/pkg/firebase"
)

type Middleware struct {
	FCS  firebase.IFaceFCM
	Repo repository.IFaceRepository
}

func (m Middleware) AuthenticatedUser() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			lenToken := 2
			s := strings.Split(token, " ")

			if len(s) != lenToken {
				response.UnauthorizedError(w)
				return
			}

			accessToken := s[1]
			resJwt, err := m.FCS.VerifiyToken(r.Context(), accessToken)
			if err != nil {
				log.Println(err)
				response.UnauthorizedError(w)
				return
			}

			ctx := r.Context()
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
			queryUser := request.ListUserRequest{
				BaseQuery: request.BaseQuery{
					PerPage: -1,
				},
				Role: "admin",
			}

			listAdmin, _, err := m.Repo.FindUsers(ctx, &queryUser)
			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			if len(listAdmin) == 0 {
				response.UnauthorizedError(w)
				return
			}

			for _, admin := range listAdmin {
				if admin.FirebaseID != claims.UID {
					response.UnauthorizedError(w)
					return
				}
				ctx = context.WithValue(ctx, constant.UserID, admin.UserID)
			}

			ctx = context.WithValue(ctx, constant.FirebaseID, claims.UID)
			h.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
