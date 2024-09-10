package main

import (
	"io"
	"os"

	"bitmap/internal/bitmap"
)

func main() {
	// Open the BMP file
	file, err := os.Open("sample.bmp")
	errNil(err)
	defer file.Close()

	// Read the entire file into memory
	data, err := io.ReadAll(file)
	errNil(err)

	// Read and output header
	var header bitmap.Header
	header.ReadHeader(data)
	header.HeaderInfo()

	// Make new image
	newImg := bitmap.MirrorVertical(data[header.StartingAddress:], int(header.Width), int(header.Height), int(header.BitsPerPixel))
	// fmt.Println(data[header.StartingAddress:])
	// fmt.Println(newImg)
	// Crete new BMP file
	bitmap.CreateBMP(&header, newImg, "output.bmp")
}

func errNil(err error) {
	if err != nil {
		panic(err)
	}
}
