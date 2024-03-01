package app

import (
	"flag"
	"fmt"

	"github.com/ansxy/nagabelajar-be-go/config"
	"github.com/ansxy/nagabelajar-be-go/database"
	"github.com/ansxy/nagabelajar-be-go/gateway/http"
	"github.com/ansxy/nagabelajar-be-go/internal/repository"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	custom_http "github.com/ansxy/nagabelajar-be-go/pkg/http"
	"github.com/ansxy/nagabelajar-be-go/pkg/postgres"
	"github.com/ansxy/nagabelajar-be-go/pkg/xendit"
)

func Run() (err error) {

	conf := config.SetConfig()
	db, err := postgres.NewPostgresClient(conf)
	if err != nil {
		return err
	}

	err = database.AutoMigrate(db)
	if err != nil {
		return err
	}
	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(&usecase.Usecase{
		Conf:   conf,
		Repo:   repo,
		Xendit: xendit.Xendit{Conf: conf},
	})
	addr := flag.String("http", fmt.Sprintf(":%d", 3000), "HTTP listen address")
	handler := http.NewHTTPHandler(conf, uc)
	err = custom_http.NewHTTPServer(*addr, handler, conf)
	if err != nil {
		return err
	}

	return
}
