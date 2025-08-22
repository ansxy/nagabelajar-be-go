package usecase

import (
	"bytes"
	"context"
	"mime/multipart"

	"github.com/ansxy/nagabelajar-be-go/internal/model"
)

// UploadMedia implements IFaceUsecase.
func (uc *Usecase) UploadMedia(ctx context.Context, data *multipart.FileHeader) (*model.Media, error) {

	file, err := data.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	pdfFileReader := bytes.NewReader(buf.Bytes())
	url, err := uc.FC.UploudFile(ctx, data, pdfFileReader)
	if err != nil {
		return nil, err
	}

	media := &model.Media{
		Name:     data.Filename,
		Type:     data.Header.Get("Content-Type"),
		UrlMedia: url,
	}

	err = uc.Repo.CreateMedia(ctx, media)
	if err != nil {
		return nil, err
	}

	return media, nil
}
