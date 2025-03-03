package ascii

import "image/color"

func getPixelBrightnessIndex(y uint8) int {
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
	if index != -1 && index < len(ascii_chars) {
		return ascii_chars[index]
	}
	return ' '
}

func GetAsciiArt(grayPixels [][]color.Gray) [][]rune {
	asciiArt := make([][]rune, len(grayPixels))
	for y := 0; y < len(grayPixels); y++ {
		asciiArt[y] = make([]rune, len(grayPixels[y]))
		for x := 0; x < len(grayPixels[y]); x++ {
			asciiArt[y][x] = getAsciiChar(grayPixels[y][x])
		}
	}
	return asciiArt
}
