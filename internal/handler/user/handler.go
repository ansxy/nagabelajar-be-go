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
		route.Route("/google", func(route chi.Router) {
			route.Use(mw.AuthenticatedUser())
			route.Post("/", handler.LoginGoogle)
		})
	})

	routes.Route("/upload", func(route chi.Router) {
		route.Post("/", handler.UploudFile)
	})

	routes.Route("/user", func(route chi.Router) {
		route.Use(mw.AuthenticatedUser())
		route.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("profile"))
		})
	})

	routes.Route("/course", func(route chi.Router) {
		route.Use(mw.AuthenticatedUser())
		route.Get("/", handler.GetListCourse)
		route.Get("/{course_id}", handler.GetOneCourse)
		route.Post("/{course_id}/enroll", handler.CreateEnrollment)
	})

	routes.Route("/progress", func(route chi.Router) {
		route.Use(mw.AuthenticatedUser())
		route.Put("/{progress_id}", handler.UpdateProgress)
	})

	routes.Route("/course/detail", func(route chi.Router) {
		route.Use(mw.AuthenticatedUser())
		route.Get("/{course_detail_id}", handler.GetOneCourse)
	})

	routes.Route("/certificate", func(route chi.Router) {
		route.Use(mw.AuthenticatedUser())
		route.Post("/validate", handler.ValidateCertificate)
		route.Get("/validate", handler.ValidateCertificateByAddress)
		route.Post("/", handler.CreateCertificate)
		route.Get("/", handler.GetListCertificate)
	})

	return routes
}
