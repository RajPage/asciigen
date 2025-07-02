package main

import (
	"asciigen/ascii"
	bmap "asciigen/bitmap"
	"asciigen/readfile"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	InputFile        string
	OutputFile       string
	Width            int
	UseInterpolation bool
	InvertColors     bool
	Verbose          bool
}

func main() {
	config := parseFlags()

	if err := validateConfig(config); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	if err := processImage(config); err != nil {
		log.Fatalf("Processing error: %v", err)
	}
}

func parseFlags() Config {
	var config Config

	flag.StringVar(&config.InputFile, "input", "", "Input image file path (required)")
	flag.StringVar(&config.InputFile, "i", "", "Input image file path (shorthand)")
	flag.StringVar(&config.OutputFile, "output", "", "Output file path (default: stdout)")
	flag.StringVar(&config.OutputFile, "o", "", "Output file path (shorthand)")
	flag.IntVar(&config.Width, "width", 100, "Output width in characters")
	flag.IntVar(&config.Width, "w", 100, "Output width in characters (shorthand)")
	flag.BoolVar(&config.UseInterpolation, "interpolation", true, "Use bilinear interpolation for resizing")
	flag.BoolVar(&config.InvertColors, "invert", false, "Invert colors for light mode")
	flag.BoolVar(&config.InvertColors, "light", false, "Enable light mode (same as -invert)")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.Verbose, "v", false, "Enable verbose output (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "AsciiGen - Convert images to ASCII art\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -input photo.jpg -width 80\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -i photo.jpg -w 120 -invert -o output.txt\n", os.Args[0])
	}

	flag.Parse()

	// Handle the case where -light is used instead of -invert
	if config.InvertColors {
		config.InvertColors = true
	}

	return config
}

func validateConfig(config Config) error {
	if config.InputFile == "" {
		return fmt.Errorf("input file is required")
	}

	if !fileExists(config.InputFile) {
		return fmt.Errorf("input file does not exist: %s", config.InputFile)
	}

	if !isImageFile(config.InputFile) {
		return fmt.Errorf("input file is not a supported image format (jpg, jpeg, png): %s", config.InputFile)
	}

	if config.Width <= 0 {
		return fmt.Errorf("width must be greater than 0")
	}

	if config.Width > 1000 {
		return fmt.Errorf("width cannot exceed 1000 characters")
	}

	return nil
}

func processImage(config Config) error {
	if config.Verbose {
		fmt.Fprintf(os.Stderr, "Reading image: %s\n", config.InputFile)
	}

	img, err := readfile.FromPath(config.InputFile)
	if err != nil {
		return fmt.Errorf("failed to read image: %w", err)
	}

	if config.Verbose {
		bounds := img.Bounds()
		fmt.Fprintf(os.Stderr, "Image dimensions: %dx%d\n", bounds.Dx(), bounds.Dy())
	}

	bitmap := bmap.FromImage(img)

	// Resize if needed
	var resized *bmap.Bitmap
	if config.UseInterpolation && config.Width < bitmap.Width {
		if config.Verbose {
			fmt.Fprintf(os.Stderr, "Resizing to width: %d\n", config.Width)
		}
		resized = bitmap.ResizeByInterpolation(config.Width)
	} else {
		resized = &bitmap
	}

	if config.Verbose {
		fmt.Fprintf(os.Stderr, "Final dimensions: %dx%d\n", resized.Width, resized.Height)
	}

	// Convert to grayscale
	grayBMap := resized.ToGrayscale(config.InvertColors)

	// Generate ASCII art
	asciiArt := ascii.GetAsciiArt(grayBMap)

	// Output result
	return outputAsciiArt(asciiArt, config.OutputFile)
}

func outputAsciiArt(asciiArt [][]rune, outputFile string) error {
	var output *os.File
	var err error

	if outputFile == "" {
		output = os.Stdout
	} else {
		output, err = os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer output.Close()
	}

	for _, row := range asciiArt {
		for _, char := range row {
			if _, err := fmt.Fprintf(output, "%c", char); err != nil {
				return fmt.Errorf("failed to write output: %w", err)
			}
		}
		if _, err := fmt.Fprintln(output); err != nil {
			return fmt.Errorf("failed to write newline: %w", err)
		}
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}