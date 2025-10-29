package ascii

import (
	"image/color"
)

func getPixelBrightnessIndex(y uint8) int {
	if len(BRIGHTNESS) == 0 {
		return -1
	}

	lo := 0
	hi := len(BRIGHTNESS) - 1
	result := -1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if y == BRIGHTNESS[mid] {
			return mid
		}
		if y < BRIGHTNESS[mid] {
			result = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return result
}

func getAsciiChar(gray color.Gray) rune {
	index := getPixelBrightnessIndex(gray.Y)
	ascii_chars := []rune(ASCII_CHARS)
	
	if index == -1 {
		// Fallback: use space for unmapped values
		return ' '
	}
	
	if index >= len(ascii_chars) {
		// Fallback: use the last character
		return ascii_chars[len(ascii_chars)-1]
	}
	
	return ascii_chars[index]
}

func GetAsciiArt(grayPixels [][]color.Gray) [][]rune {
	if len(grayPixels) == 0 {
		return [][]rune{}
	}

	asciiArt := make([][]rune, len(grayPixels))
	for y := 0; y < len(grayPixels); y++ {
		if len(grayPixels[y]) == 0 {
			asciiArt[y] = []rune{}
			continue
		}
		
		asciiArt[y] = make([]rune, len(grayPixels[y]))
		for x := 0; x < len(grayPixels[y]); x++ {
			asciiArt[y][x] = getAsciiChar(grayPixels[y][x])
		}
	}
	return asciiArt
}