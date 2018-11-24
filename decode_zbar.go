// +build zbar

package qrcode

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/admpub/qrcode/decode"
)

func Decode(imgPath string) (string, error) {
	var body string
	var img image.Image
	var err error
	file, err := os.Open(imgPath)
	if err != nil {
		return body, err
	}
	defer file.Close()

	imageTypeArr := strings.Split(file.Name(), ".")
	if len(imageTypeArr) <= 1 {
		fmt.Println("Image file format error")
		os.Exit(-1)
	}

	imageType := imageTypeArr[len(imageTypeArr)-1]

	switch imageType {
	case "jpeg", "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		fmt.Println("Image file format error")
		os.Exit(-1)
	}

	if err != nil {
		err = errors.New("decode failed: " + err.Error())
		return body, err
	}

	newImg := decode.NewImage(img)
	scanner := decode.NewScanner().SetEnabledAll(true)

	symbols, _ := scanner.ScanImage(newImg)
	for _, s := range symbols {
		body += s.Data
	}

	return body, err
}
