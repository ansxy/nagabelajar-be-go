package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custome_jwt "github.com/ansxy/nagabelajar-be-go/pkg/jwt"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
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
			resJwt, err := jwt.ParseWithClaims(accessToken, &custome_jwt.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})

			if err != nil {
				response.UnauthorizedError(w)
				return
			}

			customClaims, ok := resJwt.Claims.(*custome_jwt.UserClaims)
			if !ok || !resJwt.Valid {
				response.UnauthorizedError(w)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, constant.UserID, customClaims.UserID)
			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
