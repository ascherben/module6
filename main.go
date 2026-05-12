package main

import (
	"flag"
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"strings"
)

type Job struct {
	Image   image.Image
	OutPath string
	Err     error
}

func outputPath(inputPath string) string {
	return strings.Replace(inputPath, "images/", "images/output/", 1)
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)

	go func() {
		// For each input path create a job and add it to the out channel
		for _, p := range paths {
			job := Job{
				OutPath: outputPath(p),
			}

			img, err := imageprocessing.ReadImage(p)
			if err != nil {
				fmt.Println(err)
				job.Err = err
				out <- job
				continue
			}

			job.Image = img
			out <- job
		}

		close(out)
	}()

	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)

	go func() {
		// For each input job, create a new job after resize and add it to the out channel
		for job := range input { // Read from the channel
			if job.Err == nil {
				job.Image = imageprocessing.Resize(job.Image)
			}

			out <- job
		}

		close(out)
	}()

	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)

	go func() {
		for job := range input { // Read from the channel
			if job.Err == nil {
				job.Image = imageprocessing.Grayscale(job.Image)
			}

			out <- job
		}

		close(out)
	}()

	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)

	go func() {
		for job := range input { // Read from the channel
			if job.Err != nil {
				out <- false
				continue
			}

			err := imageprocessing.WriteImage(job.OutPath, job.Image)
			if err != nil {
				fmt.Println(err)
				out <- false
				continue
			}

			out <- true
		}

		close(out)
	}()

	return out
}

func main() {
	// Use a flag to choose between concurrent and sequential modes
	mode := flag.String("mode", "concurrent", "choose concurrent or sequential")
	flag.Parse()

	// List of image paths to process
	imagePaths := []string{
		"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	var success bool

	// Call the function based on the mode
	switch *mode {
	case "concurrent":
		success = concurrent(imagePaths)
	case "sequential":
		success = sequential(imagePaths)
	default:
		fmt.Println("Failed: unknown mode")
		return
	}

	if success {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failed!")
	}
}

// concurrent runs the pipeline with goroutines and channels
func concurrent(paths []string) bool {
	// Each stage runs in its own goroutine so load, resize, grayscale, and save
	// at the same time across images rather than waiting for each to fully complete.
	channel1 := loadImage(paths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	conSuccess := true

	for success := range writeResults {
		if !success {
			conSuccess = false
		}
	}

	return conSuccess
}

// sequential runs the pipeline one image at a time without goroutines
func sequential(paths []string) bool {
	// For each input path, read the image, resize it, convert it to grayscale, and save it to the output path
	for _, p := range paths {
		job := Job{
			OutPath: outputPath(p),
		}

		img, err := imageprocessing.ReadImage(p)
		if err != nil {
			fmt.Println(err)
			return false
		}

		job.Image = img
		job.Image = imageprocessing.Resize(job.Image)
		job.Image = imageprocessing.Grayscale(job.Image)

		err = imageprocessing.WriteImage(job.OutPath, job.Image)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	return true
}
