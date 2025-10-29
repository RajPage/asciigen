package main

import (
	"os"
	"testing"
)

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name      string
		config    Config
		shouldErr bool
	}{
		{
			name:      "empty input file",
			config:    Config{InputFile: ""},
			shouldErr: true,
		},
		{
			name:      "non-existent file",
			config:    Config{InputFile: "nonexistent.jpg"},
			shouldErr: true,
		},
		{
			name:      "zero width",
			config:    Config{InputFile: "test.jpg", Width: 0},
			shouldErr: true,
		},
		{
			name:      "negative width",
			config:    Config{InputFile: "test.jpg", Width: -10},
			shouldErr: true,
		},
		{
			name:      "width too large",
			config:    Config{InputFile: "test.jpg", Width: 1001},
			shouldErr: true,
		},
		{
			name:      "unsupported file type",
			config:    Config{InputFile: "test.txt", Width: 100},
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConfig(tt.config)
			if tt.shouldErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{"existing file", tmpFile.Name(), true},
		{"non-existing file", "definitely_does_not_exist.jpg", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fileExists(tt.filename)
			if result != tt.expected {
				t.Errorf("fileExists(%s) = %v, want %v", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestIsImageFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{"JPEG file", "photo.jpg", true},
		{"JPEG uppercase", "photo.JPG", true},
		{"JPEG long extension", "photo.jpeg", true},
		{"PNG file", "image.png", true},
		{"PNG uppercase", "image.PNG", true},
		{"text file", "document.txt", false},
		{"no extension", "filename", false},
		{"other image format", "image.gif", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isImageFile(tt.filename)
			if result != tt.expected {
				t.Errorf("isImageFile(%s) = %v, want %v", tt.filename, result, tt.expected)
			}
		})
	}
}
