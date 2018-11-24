package qrcode

import (
	"errors"
	"image"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

const (
	VERSION = 1.0
)

func Encode(value string, width, height int) (image.Image, error) {
	code, err := qr.Encode(value, qr.L, qr.Unicode)
	if err != nil {
		return nil, err
	}

	if value != code.Content() {
		return nil, errors.New("data differs")
	}

	codeImg, err := barcode.Scale(code, width, height)
	return codeImg, err
}
