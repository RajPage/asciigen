package main

import (
	"asciigen/ascii"
	bmap "asciigen/bitmap"
	"asciigen/readfile"
	"fmt"
	"image/color"
)

func main() {
	// TODO(#4): Handle Flag inputs
	filePath := "raj.jpg"
	img, _ := readfile.FromPath(filePath)
	bitmap := bmap.FromImage(img)

	fmt.Println("Successfully read image", filePath)
	fmt.Println("First pixel:", bitmap.Pixels[0][0])
	gray := color.GrayModel.Convert(bitmap.Pixels[0][0])
	fmt.Println("First pixel value:", gray.(color.Gray).Y)

	// resized := bitmap.ResizeByInterpolation(100)
	resized := bitmap
	fmt.Println("Resized bitmap width:", resized.Width)
	fmt.Println("Resized bitmap height:", resized.Height)

	grayBMap := resized.ToGrayscale()
	fmt.Println("First gray pixel:", grayBMap[0][0])
	fmt.Println("First gray pixel value:", grayBMap[0][0].Y)

	asciiArt := ascii.GetAsciiArt(grayBMap)
	for y := 0; y < len(asciiArt); y++ {
		for x := 0; x < len(asciiArt[y]); x++ {
			fmt.Printf("%c", asciiArt[y][x])
			// TODO(#5): Add Support for color output in ascii art
			// This might be helpful https://github.com/fatih/color?tab=readme-ov-file#rgb-colors
			// fmt.Printf("%c%c", asciiArt[y][x], asciiArt[y][x])
		}
		fmt.Println()
	}
}
