package http

import (
	"net/http"
	"os"

	"github.com/ansxy/nagabelajar-be-go/config"
	v1 "github.com/ansxy/nagabelajar-be-go/gateway/http/v1"
	"github.com/ansxy/nagabelajar-be-go/internal/middleware"
	"github.com/ansxy/nagabelajar-be-go/internal/response"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	"github.com/ansxy/nagabelajar-be-go/pkg/firebase"
	custom_validator "github.com/ansxy/nagabelajar-be-go/pkg/validator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	go_validator "github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

func NewHTTPHandler(conf *config.Config, uc usecase.IFaceUsecase, fc firebase.IFaceFCM) http.Handler {
	r := chi.NewRouter()
	validator := custom_validator.NewValidator(go_validator.New())
	mw := middleware.Middleware{
		FCS: fc,
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
	r.Use(middleware.Logger(&logger))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response.Success(w, "OK")
	})

	r.Mount("/v1", v1.NewRoutes(uc, validator, mw))

	return r
}
