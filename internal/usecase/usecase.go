package usecase

import (
	"github.com/ansxy/nagabelajar-be-go/config"
	"github.com/ansxy/nagabelajar-be-go/internal/repository"
	"github.com/ansxy/nagabelajar-be-go/pkg/xendit"
)

type Usecase struct {
	Conf   *config.Config
	Repo   repository.IFaceRepository
	Xendit xendit.Xendit
}

func NewUsecase(u *Usecase) IFaceUsecase {
	return &Usecase{
		Repo: u.Repo,
		Conf: u.Conf,
	}
}
