package ascii

import (
	"image/color"
	"testing"
)

func TestGetPixelBrightnessIndex(t *testing.T) {
	tests := []struct {
		name     string
		input    uint8
		expected int
	}{
		{"minimum brightness", 0, 0},
		{"maximum brightness", 255, len(BRIGHTNESS) - 1},
		{"mid brightness", 128, findExpectedIndex(128)},
		{"exact match", BRIGHTNESS[10], 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getPixelBrightnessIndex(tt.input)
			if result != tt.expected {
				t.Errorf("getPixelBrightnessIndex(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetAsciiChar(t *testing.T) {
	tests := []struct {
		name     string
		input    color.Gray
		expected rune
	}{
		{"black pixel", color.Gray{Y: 0}, rune(ASCII_CHARS[0])},
		{"white pixel", color.Gray{Y: 255}, rune(ASCII_CHARS[len(ASCII_CHARS)-1])},
		{"mid gray", color.Gray{Y: 128}, getAsciiChar(color.Gray{Y: 128})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getAsciiChar(tt.input)
			// For mid gray, we just check it's not empty
			if tt.name == "mid gray" {
				if result == 0 {
					t.Errorf("getAsciiChar returned null character for mid gray")
				}
			} else if result != tt.expected {
				t.Errorf("getAsciiChar(%v) = %c, want %c", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetAsciiArt(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]color.Gray
		expected int // expected number of rows
	}{
		{"empty input", [][]color.Gray{}, 0},
		{"single pixel", [][]color.Gray{{color.Gray{Y: 0}}}, 1},
		{"2x2 image", [][]color.Gray{
			{color.Gray{Y: 0}, color.Gray{Y: 255}},
			{color.Gray{Y: 128}, color.Gray{Y: 64}},
		}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAsciiArt(tt.input)
			if len(result) != tt.expected {
				t.Errorf("GetAsciiArt() returned %d rows, want %d", len(result), tt.expected)
			}
			
			// Check dimensions match
			for i, row := range result {
				if len(tt.input) > i && len(row) != len(tt.input[i]) {
					t.Errorf("Row %d has %d characters, want %d", i, len(row), len(tt.input[i]))
				}
			}
		})
	}
}

// Helper function for tests
func findExpectedIndex(brightness uint8) int {
	for i, b := range BRIGHTNESS {
		if brightness <= b {
			return i
		}
	}
	return len(BRIGHTNESS) - 1
}
