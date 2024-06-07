package service

import (
	"bytes"
	"context"
	"mime/multipart"

	"github.com/ansxy/nagabelajar-be-go/pkg/html"
	"github.com/ansxy/nagabelajar-be-go/pkg/pdf"
	"github.com/ansxy/nagabelajar-be-go/pkg/qrcode"
)

// GeneratePDFFromHtml implements IService.
func (s *Service) GeneratePDFFromHtml(html string, width float64, heigth float64) ([]byte, error) {
	buf, err := pdf.GeneratePDFFromHtml(html, width, heigth)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// GetMd5Hash implements IService.
func (s *Service) GetMd5Hash(ctx context.Context, filename string) (string, error) {
	panic("unimplemented")
}

// QRCodeGenerator implements IService.
func (s *Service) QRCodeGenerator(text string) (string, error) {
	qrImage, err := qrcode.GenerateQRCode(text)

	if err != nil {
		return "", err
	}

	return qrImage, nil
}

// UploudFile implements IService.
func (s *Service) UploudFile(ctx context.Context, file []byte, filename string) (string, error) {
	panic("unimplemented")
}

// ParseTemplateHTML implements IService.
func (s *Service) ParseTemplateHTML(templateName string, data map[string]interface{}) (string, error) {
	page, err := html.ParseTemplateHTML(templateName, data)
	if err != nil {
		return "", err
	}

	return page.String(), nil
}

// Md5Reader implements IService.
func (s *Service) Md5Reader(ctx context.Context, fileName string, buf []byte) (string, error) {
	pdfFileHeader := &multipart.FileHeader{
		Filename: fileName,
		Size:     int64(len(buf)),
	}

	pdfFilereader := bytes.NewReader(buf)

	_, err := s.FCM.UploudFile(ctx, pdfFileHeader, pdfFilereader)
	if err != nil {
		return "", err
	}

	md5File, err := s.FCM.GetMd5Hash(ctx, fileName)

	if err != nil {
		return "", err
	}

	return md5File, nil
}
