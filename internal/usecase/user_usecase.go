package usecase

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/ansxy/nagabelajar-be-go/pkg/hash"
	"gorm.io/gorm"
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

// LoginWithGoogle implements IFaceUsecase.
func (u *Usecase) LoginWithGoogle(ctx context.Context, data *request.LoginWithGoogleRequest) (*model.User, error) {
	var user *model.User
	user, err := u.Repo.FindOneUser(ctx, &model.User{
		FirebaseID: data.FirebaseID,
	})

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			user = &model.User{
				Email:      data.Email,
				FirebaseID: data.FirebaseID,
				Role:       constant.Role.User,
				Name:       data.Email,
				IsGoogle:   true,
				Password:   "",
			}
			err = u.Repo.CreateUser(ctx, user)
			if err != nil {
				return nil, err
			}

			return user, nil
		}
	}
	return user, nil
}
