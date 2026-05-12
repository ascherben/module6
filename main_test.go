package main

import "testing"

var testImagePaths = []string{
	"images/image1.jpeg",
	"images/image2.jpeg",
	"images/image3.jpeg",
	"images/image4.jpeg",
}

// runs concurrent pipeline benchmark
func BenchmarkConcurrentPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		success := concurrent(testImagePaths)
		if !success {
			b.Fatal("concurrent pipeline failed")
		}
	}
}

// runs the sequential pipeline benchmark
func BenchmarkSequentialPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		success := sequential(testImagePaths)
		if !success {
			b.Fatal("sequential pipeline failed")
		}
	}
}
