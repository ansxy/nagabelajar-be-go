package usecase

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/pkg/hash"
)

// RegisterUser implements IFaceUsecase.
func (u *Usecase) RegisterUser(ctx context.Context, data *request.UpsertUserRequest) error {
	hashPassword, err := hash.HashPassword(data.Password)
	if err != nil {
		return err
	}
	req := &model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashPassword,
	}
	return u.Repo.CreateUser(ctx, req)
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, data *request.LoginRequest) (*model.User, error) {
	user, err := u.Repo.FindOneUser(ctx, &model.User{
		Email: data.Email,
	})

	if err != nil {
		return nil, err
	}

	err = hash.ComparePassword(user.Password, data.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
