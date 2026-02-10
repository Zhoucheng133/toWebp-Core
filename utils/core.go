package utils

import (
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func Convert(path string, width, height int, output string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = webp.Encode(outFile, newImg, &webp.Options{Lossless: false, Quality: 80})
	if err != nil {
		return err
	}

	return nil
}
