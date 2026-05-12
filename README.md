# Go Image Processing Pipeline with Concurrency

## Overview

This project replicates https://github.com/code-heim/go_21_goroutines_pipeline Go image processing pipeline example and extends it to compare concurrent and sequential execution times. The program loads images, resizes them, converts them to grayscale, and writes the processed images to an output folder.

### Results

Concurrent mode was faster than sequential mode when processing four images. I ran the benchmark for concurrent and sequential 1000x and I did an additional check with running it 5x. The results were all similar. There were significant improvements in performance using the concurrent method. 

| Mode | Time per Run |
|---|---:|
| Concurrent | ~120 ms |
| Sequential | ~163 ms |

Benchmarks were run with 1,000 iterations:

```bash
go test -bench=. -benchmem -benchtime=1000x ./...
```

### Recommendation

The benchmark results suggest that Go concurrency can improve pipeline throughput for this image processing task. The concurrent pipeline allows different stages of the pipeline to work across images without waiting for each image to complete every step before the next image begins. With a larger set of images, the performance difference may become more noticeable, although results can depend on image size, hardware, and the amount of processing being performed.

---

## Getting Started

### Project Files

- `ai/`
    - `claude.txt`
    - `deepseek.txt` 
    - `overview.md`
- `image_processing/`
    - `image_processing.go` - image processing functions (read, write, resize, grayscale)
    - `image_processing_test.go` - unit tests for image processing functions
- `images/` - input JPEG images
    - `output/` - processed output images
- `go.mod`
- `go.sum`
- `main.go` - pipeline program with concurrent and sequential modes
- `main_test.go` - benchmark tests
- `README.md`

### Requirements

- Go 1.20 or higher

### Download the Repository

Download or clone the repository

### Run with Go

**Concurrent Mode:**

```bash
go run . -mode=concurrent
```
**Sequential Mode:**

```bash
go run . -mode=sequential
```

Output images are written to `images/output/`.

### Build an Executable

#### Windows

```powershell
go build -o main.exe .
.\main.exe -mode=concurrent
```

#### MacOS / Linux

```bash
go build -o main .
./main -mode=concurrent
```

## Testing and Validation

Unit tests cover `ReadImage`, `WriteImage`, `Resize`, and `Grayscale`:

```bash
go test ./...
```

#### Benchmark

```bash
go test -bench=. -benchmem -benchtime=1000x ./...
```

## Modifications from Original

- Added descriptive error messages to `ReadImage` and `WriteImage`
- Added a `--mode` flag so the program can be run with or without goroutines
- Added separate `concurrent` and `sequential` functions
- Designed the concurrent function to use goroutines and channels
- Designed the sequential function to process images one at a time without goroutines
- Replaced the original input images with custom JPEG files
- Added unit tests for image processing helper functions
- Added benchmark tests to measure and compare concurrent vs. sequential throughput

## Sources and AI Disclosure

### Sources

- https://github.com/code-heim/go_21_goroutines_pipeline
- https://github.com/nfnt/resize

### GenAI Tools

This project used GenAI tools as supplemental support for README refinement, code review, debugging, and clarification of Go tooling and syntax.

GenAI was helpful for building understanding, especially when comparing Go syntax, goroutines, channels, tests, and benchmark methods. I also used AI tools to help refine grammar and improve the clarity of the README. The final code, testing, benchmark results, and README decisions were reviewed and edited by me.