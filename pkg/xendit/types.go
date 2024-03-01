package xendit

import "github.com/xendit/xendit-go"

type InvoiceData struct {
	ExternalID         string
	Amount             float64
	PayerEmail         string
	Description        string
	Items              []xendit.InvoiceItem
	Fees               []xendit.InvoiceFee
	PaymentMethods     []string
	SuccessRedirectURL string
	FailureRedirectURL string
}
