package usecase

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

// CreateTransaction implements IFaceUsecase.
func (u *Usecase) CreateTransaction(ctx context.Context, data *request.InsertTransaction) error {
	return u.Repo.CreateTransaction(ctx, &model.Transaction{
		UserID:   data.UserID,
		CourseID: data.CourseID,
	})
}
