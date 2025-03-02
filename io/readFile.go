package io

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func ReadImageFromPath(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	file.Seek(0, 0)
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return img, nil
}

// ReadImageFromUrl
