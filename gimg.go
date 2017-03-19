package main

import (
	"os"
	// TODO: Support PNG
	//"image/png"
	"flag"
	"github.com/Lewuathe/gimg/gimg"
	"image/jpeg"
	"log"
)

func display(filename string, xsize int) {
	in, err := os.Open(filename)
	if err != nil {
		log.Fatal("Fail to read image file")
	}
	defer in.Close()

	src, err := jpeg.Decode(in)

	// Resizing in order to fit into terminal size
	resized := gimg.ResizeImage(xsize, src)

	gimg.PrintImage(resized)
}

func main() {
	var (
		xsize int
	)
	flag.IntVar(&xsize, "size", 50, "The size of output image width (default 50)")
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("Filename must be specified")
	}
	// Ignore command file itself
	filenames := flag.Args()

	for _, f := range filenames {
		display(f, xsize)
	}
}
