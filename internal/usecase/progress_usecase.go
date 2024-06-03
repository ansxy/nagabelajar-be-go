package usecase

import (
	"context"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
)

// UpdateProgress implements IFaceUsecase.
func (u *Usecase) UpdateProgress(ctx context.Context, data *request.UpdateProgressRequest) error {
	return u.Repo.UpdateProgress(ctx, &model.Progress{
		ProgressID: data.ProgressID,
		UserID:     data.UserID,
		IsFinished: true,
	})
}
