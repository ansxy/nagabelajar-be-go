package usecase

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/internal/request"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"
	"github.com/ansxy/nagabelajar-be-go/pkg/html"
	"github.com/ansxy/nagabelajar-be-go/pkg/pdf"
	"github.com/google/uuid"
)

// ValidateCertificate implements IFaceUsecase.
func (u *Usecase) ValidateCertificate(ctx context.Context, file *multipart.FileHeader) error {
	userid := uuid.MustParse("01878c5c-98e1-4cbb-a86c-325442b69d22")
	user, err := u.Repo.FindOneUser(ctx, &model.User{
		UserID: userid,
	})

	if err != nil {
		return err
	}

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

	// _, err = u.SM.Instance.CreateCertificate(u.SM.Auth, md5Hash, user.Name,ema)
	// if err != nil {
	// 	return err
	// }

	file.Filename = md5Hash + "certification_" + user.Name
	certificate, err := u.FC.GetOneFile(ctx, file.Filename)
	if err != nil {
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultNotValidCertification],
			Message:  constant.ErrorMessageMap[constant.DefaultNotValidCertification],
		})
		return err
	}

	hashMd5Certificate := md5.New()
	if _, err := io.Copy(hashMd5Certificate, certificate); err != nil {
		return err
	}

	hashCertificate := hashMd5Certificate.Sum(nil)
	md5HashCertificate := hex.EncodeToString(hashCertificate)

	if md5Hash != md5HashCertificate {
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
	userid := uuid.MustParse("01878c5c-98e1-4cbb-a86c-325442b69d22")
	user, err := u.Repo.FindOneUser(ctx, &model.User{
		UserID: userid,
	})
	if err != nil {
		return err
	}

	_, err = u.SM.Instance.CreateCertificate(u.SM.Auth, user.Name, user.Email)
	if err != nil {
		return err
	}

	address, err := u.SM.Instance.GetAddressByName(nil, user.Name)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"Address": address,
		"Name":    user.Name,
	}

	fileName := "certification_421" + user.Name
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
	err = u.FC.UploudFile(ctx, pdfFileHeader, pdfFileReader)
	if err != nil {
		return err
	}

	md5File, err := u.FC.GetMd5Hash(ctx, fileName)
	if err != nil {
		return err
	}

	_, err = u.SM.Instance.UpdateMd5Certificate(u.SM.Auth, address, md5File)
	if err != nil {
		return err
	}

	certificate := &model.Certificate{
		UserID:            user.UserID.String(),
		FileName:          fileName,
		FileUrl:           "-",
		MD5:               md5File,
		BlockchainAddress: address.String(),
	}

	return u.Repo.CreateCertificate(ctx, certificate)
}
