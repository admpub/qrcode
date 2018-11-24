package qrcode

import (
	"errors"
	"image"
	"image/png"
	"os"

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

func EncodeToFile(value string, width, height int, qrcodeFile string) error {
	img, err := Encode(value, width, height)
	if err != nil {
		return err
	}
	file, err := os.Create(qrcodeFile)
	if err != nil {
		return err
	}
	defer file.Close()
	err = png.Encode(file, img)
	return err
}
