package xendit

import (
	"github.com/ansxy/nagabelajar-be-go/config"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type IFaceXendit interface {
	CreateInvoice(data *InvoiceData) (*xendit.Invoice, error)
}

type Xendit struct {
	Conf *config.Config
}

func (x *Xendit) CreateInvoice(data *InvoiceData) (*xendit.Invoice, error) {
	xendit.Opt.SecretKey = x.Conf.XenditConfig.SecretKey

	invoiceParams := &invoice.CreateParams{
		ExternalID:         data.ExternalID,
		Amount:             data.Amount,
		PayerEmail:         data.PayerEmail,
		Description:        data.Description,
		SuccessRedirectURL: x.Conf.XenditConfig.SuccessRedirectURL,
		FailureRedirectURL: x.Conf.XenditConfig.FailureRedirectURL,
		Fees:               data.Fees,
		InvoiceDuration:    86400,
		Items:              data.Items,
	}

	invoice, err := invoice.Create(invoiceParams)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
