package bitmap

import (
	"image/color"
	"testing"
)

func TestPixelToGrayScale(t *testing.T) {
	tests := []struct {
		name     string
		pixel    color.NRGBA
		invert   bool
		expected uint8
	}{
		{
			name:     "black pixel normal",
			pixel:    color.NRGBA{R: 0, G: 0, B: 0, A: 255},
			invert:   false,
			expected: 0,
		},
		{
			name:     "black pixel inverted",
			pixel:    color.NRGBA{R: 0, G: 0, B: 0, A: 255},
			invert:   true,
			expected: 255,
		},
		{
			name:     "white pixel normal",
			pixel:    color.NRGBA{R: 255, G: 255, B: 255, A: 255},
			invert:   false,
			expected: 255,
		},
		{
			name:     "white pixel inverted",
			pixel:    color.NRGBA{R: 255, G: 255, B: 255, A: 255},
			invert:   true,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pixelToGrayScale(tt.pixel, tt.invert)
			if result.Y != tt.expected {
				t.Errorf("pixelToGrayScale(%v, %v) = %d, want %d", tt.pixel, tt.invert, result.Y, tt.expected)
			}
		})
	}
}

func TestBitmapToGrayscale(t *testing.T) {
	// Create a simple 2x2 bitmap
	bitmap := Bitmap{
		Width:  2,
		Height: 2,
		Pixels: [][]color.NRGBA{
			{
				{R: 0, G: 0, B: 0, A: 255},     // black
				{R: 255, G: 255, B: 255, A: 255}, // white
			},
			{
				{R: 128, G: 128, B: 128, A: 255}, // gray
				{R: 255, G: 0, B: 0, A: 255},     // red
			},
		},
	}

	t.Run("normal mode", func(t *testing.T) {
		result := bitmap.ToGrayscale(false)
		
		if len(result) != 2 {
			t.Errorf("Expected 2 rows, got %d", len(result))
		}
		
		if len(result[0]) != 2 {
			t.Errorf("Expected 2 columns in first row, got %d", len(result[0]))
		}
		
		// Check black pixel
		if result[0][0].Y != 0 {
			t.Errorf("Expected black pixel (0), got %d", result[0][0].Y)
		}
		
		// Check white pixel
		if result[0][1].Y != 255 {
			t.Errorf("Expected white pixel (255), got %d", result[0][1].Y)
		}
	})

	t.Run("inverted mode", func(t *testing.T) {
		result := bitmap.ToGrayscale(true)
		
		// Check inverted black pixel (should be white)
		if result[0][0].Y != 255 {
			t.Errorf("Expected inverted black pixel (255), got %d", result[0][0].Y)
		}
		
		// Check inverted white pixel (should be black)
		if result[0][1].Y != 0 {
			t.Errorf("Expected inverted white pixel (0), got %d", result[0][1].Y)
		}
	})
}
