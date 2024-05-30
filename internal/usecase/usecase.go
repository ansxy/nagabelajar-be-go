package usecase

import (
	"github.com/ansxy/nagabelajar-be-go/config"
	"github.com/ansxy/nagabelajar-be-go/internal/repository"
	"github.com/ansxy/nagabelajar-be-go/pkg/firebase"
	goeth "github.com/ansxy/nagabelajar-be-go/pkg/go-eth"
	"github.com/ansxy/nagabelajar-be-go/pkg/xendit"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Usecase struct {
	Conf      *config.Config
	Repo      repository.IFaceRepository
	DB        *gorm.DB
	Xendit    xendit.Xendit
	FC        firebase.IFaceFCM
	SM        *goeth.GoethClient
	Validator *validator.Validate
}

func NewUsecase(u *Usecase) IFaceUsecase {
	return &Usecase{
		Repo: u.Repo,
		Conf: u.Conf,
		DB:   u.DB,
		FC:   u.FC,
		SM:   u.SM,
	}
}
