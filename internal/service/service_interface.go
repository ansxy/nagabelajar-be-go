package service

import (
	"context"
	"math/big"

	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:generate mockgen -destination=../mock/service_mock.go -package=mock_service -source=service_interface.go
type IService interface {
	GeneratePDFFromHtml(html string, width, heigth float64) ([]byte, error)
	UploudFile(ctx context.Context, file []byte, filename string) (string, error)
	QRCodeGenerator(text string) (string, error)
	GetMd5Hash(ctx context.Context, filename string) (string, error)
	ParseTemplateHTML(templateName string, data map[string]interface{}) (string, error)
	CreateCertificate(ctx context.Context, opts *bind.TransactOpts, _recipientName string, _recipientEmail string, _courseName string, _courseCode string) (*common.Address, *types.Transaction, *big.Int, *big.Int, *big.Int, error)
	CertificatePDF(ctx context.Context, typeTransaction *types.Transaction, pdfData constant.CertificatePDF, address common.Address) ([]byte, error)
	Md5Reader(ctx context.Context, fileName string, buf []byte) (string, error)

	UpdateMd5Certificate(ctx context.Context, opts *bind.TransactOpts, fileName *common.Address, md5Hash string) error
}
