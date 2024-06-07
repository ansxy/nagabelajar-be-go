package public

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	custom_validator "github.com/ansxy/nagabelajar-be-go/pkg/validator"
	"github.com/go-chi/chi/v5"
)

type publicHandler struct {
	uc usecase.IFaceUsecase
	v  custom_validator.Validator
}

func NewRouter(uc usecase.IFaceUsecase, v custom_validator.Validator) http.Handler {
	routes := chi.NewRouter()
	handler := &publicHandler{
		uc: uc,
		v:  v,
	}

	routes.Route("/course", func(route chi.Router) {
		route.Get("/", handler.GetListCourse)
		route.Get("/{course_id}", handler.GetOneCourse)
	})

	routes.Route("/certificate", func(route chi.Router) {
		route.Post("/validate", handler.ValidateCertificate)
		route.Get("/validate", handler.ValidateCertificateByAddress)
	})

	return routes
}
