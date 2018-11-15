package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	outputFormat := flag.String("o", "jpeg", "output image format")
	flag.Parse()

	if err := convert(*outputFormat, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func convert(format string, in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch format {
	case "jpeg":
		err = toJPEG(img, os.Stdout)
	case "png":
		err = toPNG(img, os.Stdout)
	case "gif":
		err = toGif(img, os.Stdout)
	default:
		return fmt.Errorf("unsupported format: %v", format)
	}

	if err != nil {
		return err
	}

	return nil
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

func toGif(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, nil)
}
