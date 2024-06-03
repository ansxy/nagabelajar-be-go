package usecase

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"mime/multipart"
	"strconv"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"
	goeth "github.com/ansxy/nagabelajar-be-go/pkg/go-eth/artifact"
	"github.com/ansxy/nagabelajar-be-go/pkg/html"
	"github.com/ansxy/nagabelajar-be-go/pkg/pdf"
	"github.com/ansxy/nagabelajar-be-go/pkg/qrcode"
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

	address, err := u.SM.Instance.CreateCertificate(u.SM.Auth, user.Name, user.Email, course.Name, course.Code)
	if err != nil {
		return err
	}

	//Wait till the certificate is created
	receipt, err := bind.WaitMined(ctx, u.SM.Client, address)

	if err != nil {
		return err
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	callOpts := &bind.CallOpts{
		From: u.SM.Auth.From,
	}

	count, err := u.SM.Instance.GetCertificateCount(callOpts)
	if err != nil {
		return err
	}

	newCertAddress, err := u.SM.Instance.CertificateAddresses(callOpts, big.NewInt(count.Int64()-1))
	if err != nil {
		return err
	}

	qrImage, err := qrcode.GenerateQRCode(fmt.Sprintf("%s/certificate?address=%s", u.Conf.FRONTENDURL, newCertAddress.String()))
	if err != nil {
		return err
	}

	time := address.Time().Format("2006-01-02")
	data := map[string]interface{}{
		"Address":    newCertAddress,
		"Name":       user.Name,
		"CourseName": course.Name,
		"QRImage":    qrImage,
		"IssuerAt":   time,
	}

	fileName := newCertAddress.String() + user.Name + course.Code + ".pdf"
	page, err := html.ParseTemplateHTML("certificate.html", data)
	if err != nil {
		return err
	}

	width := 8.27
	height := 10.52
	buf, err := pdf.GeneratePDFFromHtml(page.String(), width, height)
	if err != nil {
		return err
	}

	pdfFileHeader := &multipart.FileHeader{
		Filename: fileName,
		Size:     int64(len(buf)),
	}

	pdfFileReader := bytes.NewReader(buf)
	_, err = u.FC.UploudFile(ctx, pdfFileHeader, pdfFileReader)
	if err != nil {
		return err
	}

	md5File, err := u.FC.GetMd5Hash(ctx, fileName)

	if err != nil {
		return err
	}

	_, err = u.SM.Instance.UpdateMd5Certificate(u.SM.Auth, newCertAddress, md5File)

	if err != nil {
		return err
	}

	certificate := &model.Certificate{
		UserID:            user.UserID.String(),
		FileName:          fileName,
		FileUrl:           fmt.Sprintf("%s%s%s", constant.FirebaseStorageURL, fileName, constant.StorageMediaALT),
		MD5:               md5File,
		BlockchainAddress: newCertAddress.String(),
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
