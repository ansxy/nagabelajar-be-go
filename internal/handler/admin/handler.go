package admin

import (
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/internal/middleware"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	custom_validator "github.com/ansxy/nagabelajar-be-go/pkg/validator"
	"github.com/go-chi/chi/v5"
)

type adminHandler struct {
	uc usecase.IFaceUsecase
	v  custom_validator.Validator
}

func NewRouter(uc usecase.IFaceUsecase, v custom_validator.Validator, mw middleware.Middleware) http.Handler {
	routes := chi.NewRouter()
	handler := &adminHandler{
		uc: uc,
		v:  v,
	}

	// routes.Route("/admin", func(route chi.Router) {
	// 	// route.Post("/login", handler.Login)
	// })

	routes.Route("/category", func(route chi.Router) {
		route.Post("/", handler.CreateCategory)
		route.Delete("/{category_id}", handler.DeleteCategory)
		route.Get("/", handler.FindListCategory)
		// route.Get("/{category_id}", handler.FindOneCategory)
	})

	routes.Route("/course", func(route chi.Router) {
		route.Post("/", handler.CreateCourse)
		route.Get("/", handler.GetListCourse)
		route.Get("/{course_id}", handler.GetOneCourse)
		route.Group(func(route chi.Router) {
			// route.Use(mw.AuthenticatedAdmin())
			route.Post("/{course_id}/detail", handler.CreateCourseDetail)
			route.Get("/{course_id}/detail", handler.FindCourseDetail)

		})
	})

	routes.Route("/media", func(route chi.Router) {
		route.Post("/", handler.UploadMedia)
	})

	return routes
}
