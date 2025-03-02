package main

import (
	bmap "asciigen/bitmap"
	"asciigen/readfile"
	"fmt"
)

func main() {
	filePath := "test.jpg"
	img, _ := readfile.FromPath(filePath)
	bitmap := bmap.FromImage(img)

	fmt.Println("Successfully read image", filePath)
	fmt.Println("First pixel:", bitmap.Pixels[0][0])
}
