package ansi

import (
	"testing"
)

func TestAnsi(t *testing.T) {
	testColors := []uint32{
		0,10,20,100,256,65536,
	}

	expected256Colors := []uint32{
		0,0,0,0,1,256,
	}

	for i, c := range testColors {
		if expected256Colors[i] != To256(c) {
			t.Errorf("%d: Expecting %d, but got %d", i, expected256Colors[i], To256(c))
		}
	}
}