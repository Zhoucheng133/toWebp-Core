package utils

import (
	"os"
	"path/filepath"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func Convert(path string, width, height int, output string, quality int) error {
	src, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	dst := imaging.Resize(src, width, height, imaging.Lanczos)

	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = webp.Encode(outFile, dst, &webp.Options{Lossless: false, Quality: float32(quality)})
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

func ConvertFromDir(path string, width, height int, output string, quality int) error {
	if err := os.MkdirAll(output, 0755); err != nil {
		return err
	}
	files := Scan(path)
	for _, file := range files {
		baseName := filepath.Base(file)
		ext := filepath.Ext(baseName)
		outName := baseName[:len(baseName)-len(ext)] + ".webp"
		targetPath := filepath.Join(output, outName)

		err := Convert(file, width, height, targetPath, quality)
		if err != nil {
			return err
		}
	}
	return nil
}
