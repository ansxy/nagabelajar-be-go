package usecase

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"mime/multipart"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/google/uuid"
)

// UploadFile implements IFaceUsecase.
func (u *Usecase) UploadFile(ctx context.Context, file *multipart.FileHeader) error {
	userid := uuid.MustParse(ctx.Value("user_id").(string))
	user, err := u.Repo.FindOneUser(ctx,
		&model.User{
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

	file.Filename = md5Hash + "certification_" + user.Name

	metaData := &model.Certificate{
		FileName:          file.Filename,
		UserID:            user.UserID.String(),
		MD5:               md5Hash,
		FileUrl:           constant.FirebaseStorageURL + constant.BucketName + file.Filename + constant.StorageMediaALT,
		BlockchainAddress: "asd",
	}

	err = u.Repo.CreateCertificate(ctx, metaData)
	if err != nil {

		return err
	}

	err = u.FC.UploudFile(ctx, file)
	if err != nil {
		log.Println("Error uploading file: ", err)
		return err
	}

	return nil
}
