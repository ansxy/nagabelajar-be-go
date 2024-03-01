package user

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/middleware"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	custome_validator "github.com/ansxy/nagabelajar-be-go/pkg/validator"
	"github.com/go-chi/chi/v5"
)

type userHandler struct {
	uc usecase.IFaceUsecase
	v  custome_validator.Validator
}

func NewRouter(uc usecase.IFaceUsecase, v custome_validator.Validator, mw middleware.Middleware) http.Handler {
	routes := chi.NewRouter()
	handler := &userHandler{
		uc: uc,
		v:  v,
	}

	routes.Route("/auth", func(route chi.Router) {
		route.Post("/register", handler.RegisterUser)
		route.Post("/login", handler.Login)
	})

	routes.Route("/user", func(route chi.Router) {
		route.Use(mw.AuthenticatedUser())
		route.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("profile"))
		})
	})

	return routes
}
