package towebp

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

func Towebp(fileName string) (string, error) {
	inputFile, err := os.Open(fileName) // Change to "input.png" for PNG images
	if err != nil {
		return "", err
	}
	defer inputFile.Close()

	// Decode the image based on its format
	var img image.Image
	ext := strings.ToLower(filepath.Ext(fileName))

	switch ext {
	case ".png":
		img, err = png.Decode(inputFile)
		if err != nil {
			return "", err
		}
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(inputFile)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("unsupported image format: %s", ext)
	}

	if err != nil {
		return "", err
	}
	filename := filepath.Base(fileName)
	nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
	outputPath := filepath.Join(filepath.Dir(fileName), nameWithoutExt+".webp")
	// Create the output WebP file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	// Encode the image to WebP format
	err = webp.Encode(outputFile, img, &webp.Options{Quality: 90}) // Adjust quality as needed
	if err != nil {
		return "", err
	}
	return outputPath, nil
}
