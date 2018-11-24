// +build !zbar

package qrcode

import (
	"os"

	"github.com/tuotoo/qrcode"
)

func Decode(imgPath string) (string, error) {
	fi, err := os.Open(imgPath)
	if err != nil {
		return ``, err
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		return ``, err
	}
	return qrmatrix.Content, err
}
