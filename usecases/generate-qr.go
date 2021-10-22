package usecases

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateQr(data string) (*barcode.Barcode, error) {
	qrCode, err := qr.Encode(data, qr.L, qr.Auto)

	if err != nil {
		return &qrCode, err
	}

	qrCode, err = barcode.Scale(qrCode, 512, 512)

	if err != nil {
		return &qrCode, err
	}

	return &qrCode, nil
}
