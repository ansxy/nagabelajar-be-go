package service

import (
	"context"
	"math/big"

	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CreateCertificate implements IService.
func (s *Service) CreateCertificate(ctx context.Context, opts *bind.TransactOpts, _recipientName string, _recipientEmail string, _courseName string, _courseCode string) (*common.Address, *types.Transaction, *big.Int, *big.Int, *big.Int, error) {
	address, err := s.SM.Instance.CreateCertificate(opts, _recipientName, _recipientEmail, _courseName, _courseCode)

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	receipt, err := bind.WaitMined(ctx, s.SM.Client, address)

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	if receipt.Status == 0 {
		return nil, nil, nil, nil, nil, err
	}

	gasUsed := new(big.Int).SetUint64(receipt.GasUsed)
	gasPrice := address.GasPrice()
	cost := new(big.Int).Mul(gasUsed, gasPrice)

	callOpts := &bind.CallOpts{
		From: opts.From,
	}

	count, err := s.SM.Instance.GetCertificateCount(callOpts)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	newCertAddress, err := s.SM.Instance.CertificateAddresses(callOpts, big.NewInt(count.Int64()-1))

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return &newCertAddress, address, cost, gasPrice, gasUsed, nil
}

// CertificatePDF implements IService.
func (s *Service) CertificatePDF(ctx context.Context, typeTransaction *types.Transaction, pdfData constant.CertificatePDF, address common.Address) ([]byte, error) {
	var data *constant.CertificatePDF

	time := typeTransaction.Time().Format("2006-01-02")
	data = &constant.CertificatePDF{
		QRImage:    pdfData.QRImage,
		Address:    pdfData.Address,
		Name:       pdfData.Name,
		CourseName: pdfData.CourseName,
		IssuerAt:   time,
	}

	htmlData := map[string]interface{}{
		"Address":    pdfData.Address,
		"Name":       data.Name,
		"CourseName": data.CourseName,
		"QRImage":    pdfData.QRImage,
		"IssuerAt":   time,
	}

	page, err := s.ParseTemplateHTML("certificate.html", htmlData)
	if err != nil {
		return nil, err
	}

	width := 10.52
	height := 8.27

	buf, err := s.GeneratePDFFromHtml(page, width, height)
	if err != nil {
		return nil, err
	}

	return buf, nil

}

// UpdateMd5Certificate implements IService.
func (s *Service) UpdateMd5Certificate(ctx context.Context, opts *bind.TransactOpts, fileName *common.Address, md5Hash string) error {
	_, err := s.SM.Instance.UpdateMd5Certificate(opts, *fileName, md5Hash)
	if err != nil {
		return err
	}

	return nil
}
