package main

import (
	"io"
	"os"

	"bitmap/internal/bitmap"
)

func main() {
	// Open the BMP file
	file, err := os.Open("grid.bmp")
	errNil(err)
	defer file.Close()

	// Read the entire file into memory
	data, err := io.ReadAll(file)
	errNil(err)

	var header bitmap.Header

	header.ReadHeader(data)
	
}

func errNil(err error) {
	if err != nil {
		panic(err)
	}
}