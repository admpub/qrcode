package qrcode

import (
	"fmt"
	"image/png"
	"os"
	"testing"
)

// go test -tags zbar

func TestQrcode(t *testing.T) {
	img, err := Encode("test qrcode", 300, 300)
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Create("./test.png")
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Fatal(err)
	}

	value, err := Decode("./test.png")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(value)
}
