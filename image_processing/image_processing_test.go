package imageprocessing

import (
	"os"
	"path/filepath"
	"testing"
)

// Test functions for image processing
func TestRead(t *testing.T) {
	img, err := ReadImage(filepath.Join("..", "images", "image1.jpeg"))

	if err != nil {
		t.Fatalf("expected ReadImage to work, got error: %v", err)
	}

	if img == nil {
		t.Fatal("expected image, got nil")
	}
}

func TestWrite(t *testing.T) {
	img, err := ReadImage(filepath.Join("..", "images", "image1.jpeg"))
	if err != nil {
		t.Fatalf("expected ReadImage to work, got error: %v", err)
	}

	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "test-output.jpeg")

	err = WriteImage(outputPath, img)
	if err != nil {
		t.Fatalf("expected WriteImage to work, got error: %v", err)
	}

	if _, err := os.Stat(outputPath); err != nil {
		t.Fatalf("expected output image file to exist, got error: %v", err)
	}
}

// Test the Resize function
func TestResize(t *testing.T) {
	img, err := ReadImage(filepath.Join("..", "images", "image1.jpeg"))
	if err != nil {
		t.Fatalf("expected ReadImage to work, got error: %v", err)
	}

	resized := Resize(img)

	if resized == nil {
		t.Fatal("expected resized image, got nil")
	}

	bounds := resized.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width > 500 || height > 500 {
		t.Fatalf("expected longest side <= 500, got %dx%d", width, height)
	}
}

// Test the Grayscale function
func TestGray(t *testing.T) {
	img, err := ReadImage(filepath.Join("..", "images", "image1.jpeg"))
	if err != nil {
		t.Fatalf("expected ReadImage to work, got error: %v", err)
	}

	grayImg := Grayscale(img)

	if grayImg == nil {
		t.Fatal("expected grayscale image, got nil")
	}

	// Find the center pixel
	bounds := grayImg.Bounds()
	x := bounds.Dx() / 2
	y := bounds.Dy() / 2

	// Verify only the center pixel is gray
	r, g, b, _ := grayImg.At(x, y).RGBA()
	if r != g || g != b {
		t.Fatalf("center pixel is not gray")
	}
}
