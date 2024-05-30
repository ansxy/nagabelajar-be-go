package qrcode

import (
	"encoding/base64"

	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRCode(url string) (string, error) {
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(png), nil
}
