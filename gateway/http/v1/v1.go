package v1

import (
	"net/http"

	admin_handler "github.com/ansxy/nagabelajar-be-go/internal/handler/admin"
	user_handler "github.com/ansxy/nagabelajar-be-go/internal/handler/user"
	"github.com/ansxy/nagabelajar-be-go/internal/middleware"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	custome_validator "github.com/ansxy/nagabelajar-be-go/pkg/validator"
	"github.com/go-chi/chi/v5"
)

func NewRoutes(uc usecase.IFaceUsecase, v custome_validator.Validator, mw middleware.Middleware) http.Handler {
	router := chi.NewRouter()

	router.Mount("/", user_handler.NewRouter(uc, v, mw))
	router.Mount("/admin", admin_handler.NewRouter(uc, v, mw))

	return router
}
