package service

import (
	"github.com/ansxy/nagabelajar-be-go/pkg/firebase"
	goeth "github.com/ansxy/nagabelajar-be-go/pkg/go-eth"
)

type Service struct {
	SM  *goeth.GoethClient
	FCM firebase.IFaceFCM
}

func NewService(s *Service) IService {
	return &Service{
		SM:  s.SM,
		FCM: s.FCM,
	}
}
