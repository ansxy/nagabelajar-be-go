package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"
	goeth "github.com/ansxy/nagabelajar-be-go/pkg/go-eth/artifact"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ValidateCertificate implements IFaceUsecase.
func (u *Usecase) ValidateCertificate(ctx context.Context, file *multipart.FileHeader) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}

	defer fileReader.Close()
	hasher := md5.New()
	if _, err := io.Copy(hasher, fileReader); err != nil {
		return err
	}

	hash := hasher.Sum(nil)
	md5Hash := hex.EncodeToString(hash)
	callOpts := &bind.CallOpts{
		From: u.SM.Auth.From,
	}

	data, err := u.SM.Instance.GetCertificateByFileName(callOpts, md5Hash)
	if data.String() == "0x0000000000000000000000000000000000000000" {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultNotValidCertification],
			Message:  constant.ErrorMessageMap[constant.DefaultNotValidCertification],
		})
		return err
	}

	if err != nil {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultNotValidCertification],
			Message:  constant.ErrorMessageMap[constant.DefaultNotValidCertification],
		})
		return err
	}
	return nil
}

// CreateCertificate implements IFaceUsecase.
func (u *Usecase) CreateCertificate(ctx context.Context, req *request.CreateCertificateRequest) error {
	user, err := u.Repo.FindOneUser(ctx, &model.User{
		FirebaseID: req.FirebaseID,
	})

	if err != nil {
		return err
	}

	course, err := u.Repo.FindOneCourse(ctx, &request.GetOneCourseRequest{
		CourseID: strconv.Itoa(int(req.CourseId)),
	})
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

// ValidateCertificateByAddress implements IFaceUsecase.
func (u *Usecase) ValidateCertificateByAddress(ctx context.Context, address string) (*goeth.CertificateOfCompletionCertificateData, error) {
	callOpts := &bind.CallOpts{
		From: u.SM.Auth.From,
	}

	addr := common.HexToAddress(address)
	data, err := u.SM.Instance.GetCertificate(callOpts, addr)

	if data.Recipient.String() == "0x0000000000000000000000000000000000000000" {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultNotValidCertification],
			Message:  constant.ErrorMessageMap[constant.DefaultNotValidCertification],
		})
		return nil, err

	}

	if err != nil {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultNotValidCertification],
			Message:  constant.ErrorMessageMap[constant.DefaultNotValidCertification],
		})
		return nil, err
	}

	return &data, err
}

// GetListCertificate implements IFaceUsecase.
func (u *Usecase) GetListCertificate(ctx context.Context, params *request.ListCertificateRequest) ([]model.Certificate, int64, error) {
	var certificate []model.Certificate
	var count int64

	certificate, count, err := u.Repo.FindListCertificate(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return certificate, count, nil
}
