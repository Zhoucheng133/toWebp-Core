package utils

import (
	"image"
	"os"
	"path/filepath"
	"strings"

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

func Scan(path string) []string {
	var files []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return files
	}
	extensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			ext := strings.ToLower(filepath.Ext(entry.Name()))
			if extensions[ext] {
				fullPath := filepath.Join(path, entry.Name())
				files = append(files, fullPath)
			}
		}
	}
	return files
}

func ConvertFromDir(path string, width, height int, output string) error {
	files := Scan(path)
	for _, file := range files {
		err := Convert(file, width, height, output)
		if err != nil {
			return err
		}
	}
	return nil
}
