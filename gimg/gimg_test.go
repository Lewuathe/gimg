package gimg

import (
	"os"
	"testing"
	"image/jpeg"
)

func TestAnsi(t *testing.T) {
	file, err := os.Open("./test.jpg")
	if err != nil {
		t.Error("Fail to open test.jpg")
	}
	defer file.Close()

	testImage, err := jpeg.Decode(file)
	if err != nil {
		t.Error("Fail to decode image file")
	}

	for i, s := range []int {
		10, 100, 400, 1000,
	} {
		r := ResizeImage(s, testImage)
		if r.Bounds().Max.X != s {
			t.Errorf("%d: Expected width %d, but got %d", i, s, r.Bounds().Max.X)
		}
	}
}