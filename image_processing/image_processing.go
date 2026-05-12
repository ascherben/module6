package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// ReadImage reads an image from the specified path and returns it as an image.Image.
func ReadImage(path string) (image.Image, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open image %s: %w", path, err)
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return nil, fmt.Errorf("could not decode image %s: %w", path, err)
	}
	return img, nil
}

// WriteImage writes the given image to the specified path in JPEG format.
func WriteImage(path string, img image.Image) error {
	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create output image %s: %w", path, err)
	}
	defer outputFile.Close()

	// Encode the image to the new file
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		return fmt.Errorf("could not write output image %s: %w", path, err)
	}
	return nil
}

// Grayscale converts the given image to grayscale and returns the new image.
func Grayscale(img image.Image) image.Image {
	// Create a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

// Resize resizes the given image to fit within a 500x500 pixel box while maintaining the aspect ratio.
func Resize(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	maxSize := uint(500)

	// Pass 0 for the unconstrained dimension so the library preserves aspect ratio.
	if width >= height {
		return resize.Resize(maxSize, 0, img, resize.Lanczos3)
	}

	return resize.Resize(0, maxSize, img, resize.Lanczos3)
}
