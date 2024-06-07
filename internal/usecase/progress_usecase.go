package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
)

// UpdateProgress implements IFaceUsecase.
func (u *Usecase) UpdateProgress(ctx context.Context, data *request.UpdateProgressRequest) error {
	err := u.Repo.UpdateProgress(ctx, &model.Progress{
		ProgressID: data.ProgressID,
		UserID:     data.UserID,
		IsFinished: true,
	})

	if err != nil {
		return err
	}

	progress, err := u.Repo.FindOneProgress(ctx, "progress_id = ?", data.ProgressID)

	if err != nil {
		return err
	}

	listProgress, _, err := u.Repo.FindListProgress(ctx, progress.CourseID, progress.UserID)

	if err != nil {
		return err
	}

	for _, progress := range listProgress {
		if !progress.IsFinished {
			return nil
		}
	}
	courseId := strconv.Itoa(progress.CourseID)
	course, err := u.Repo.FindOneCourse(ctx, &request.GetOneCourseRequest{
		CourseID: courseId,
	})

	if err != nil {
		return err
	}
	user, err := u.Repo.FindOneUser(ctx, "user_id = ?", progress.UserID)

	if err != nil {
		return err
	}

	newCertAddress, address, cost, gasPrice, gasUsed, err := u.Service.CreateCertificate(ctx, u.SM.Auth, user.Name, user.Email, course.Name, course.Code)

	if err != nil {
		return err
	}

	qrImage, err := u.Service.QRCodeGenerator(fmt.Sprintf("%s/certificate?address=%s", u.Conf.FRONTENDURL, newCertAddress.String()))
	if err != nil {
		return err
	}

	pdf, err := u.Service.CertificatePDF(ctx, address, constant.CertificatePDF{
		QRImage:    qrImage,
		Address:    newCertAddress.String(),
		Name:       user.Name,
		CourseName: course.Name,
	}, *newCertAddress)

	if err != nil {
		return err
	}

	fileName := newCertAddress.String() + user.Name + course.Name + ".pdf"

	md5File, err := u.Service.Md5Reader(ctx, fileName, pdf)
	if err != nil {
		return err
	}

	err = u.Service.UpdateMd5Certificate(ctx, u.SM.Auth, newCertAddress, md5File)

	if err != nil {
		return err
	}

	certificate := &model.Certificate{
		UserID:            user.UserID.String(),
		FileName:          fileName,
		FileUrl:           fmt.Sprintf("%s%s%s", constant.FirebaseStorageURL, fileName, constant.StorageMediaALT),
		MD5:               md5File,
		BlockchainAddress: newCertAddress.String(),
		CourseID:          strconv.Itoa(int(course.CourseID)),
		GasUsed:           gasUsed.String(),
		GasPrice:          gasPrice.String(),
		Cost:              cost.String(),
	}

	return u.Repo.CreateCertificate(ctx, certificate)

}
