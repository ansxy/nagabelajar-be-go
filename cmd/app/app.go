package app

import (
	"flag"
	"fmt"

	"github.com/ansxy/nagabelajar-be-go/config"
	"github.com/ansxy/nagabelajar-be-go/database"
	"github.com/ansxy/nagabelajar-be-go/gateway/http"
	"github.com/ansxy/nagabelajar-be-go/internal/repository"
	"github.com/ansxy/nagabelajar-be-go/internal/service"
	"github.com/ansxy/nagabelajar-be-go/internal/usecase"
	"github.com/ansxy/nagabelajar-be-go/pkg/firebase"
	goeth "github.com/ansxy/nagabelajar-be-go/pkg/go-eth"
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

	fc, err := firebase.NewFCMClient(conf.FirebaseConfig)
	if err != nil {
		return err
	}

	goeth, err := goeth.NewGoethClient(conf.SmartContractConfig)
	if err != nil {
		return err
	}

	err = database.AutoMigrate(db)
	if err != nil {
		return err
	}

	service := service.NewService(
		&service.Service{
			SM:  goeth,
			FCM: fc,
		},
	)
	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(&usecase.Usecase{
		Conf:    conf,
		FC:      fc,
		Repo:    repo,
		Xendit:  xendit.Xendit{Conf: conf},
		SM:      goeth,
		Service: service,
	})

	addr := flag.String("http", fmt.Sprintf(":%d", 3000), "HTTP listen address")
	handler := http.NewHTTPHandler(conf, uc, fc)
	err = custom_http.NewHTTPServer(*addr, handler, conf)
	if err != nil {
		return err
	}

	return
}
