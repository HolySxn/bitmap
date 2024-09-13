package main

import (
	"bitmap/internal/bitmap"
	"fmt"
	"io"
	"os"
)

func main() {
	var readData []string
	readData = append(readData, os.Args...)
	if readData[1] == "apply" {
		// Open the BMP file
		file, err := os.Open("grid.bmp")
		errNil(err)
		defer file.Close()

		// Read the entire file into memory
		data, err := io.ReadAll(file)
		errNil(err)

		// Read and output header
		var header bitmap.Header

		header.ReadHeader(data)
		header.HeaderInfo()

		pixelMap := bitmap.PixelMap(data[header.StartingAddress:], int(header.Width), int(header.Height), int(header.BitsPerPixel))
		// bitmap.MirrorVertical(pixelMap, int(header.Width), int(header.Height), int(header.BitsPerPixel))
		newMap := bitmap.Rotate90(pixelMap, int(header.Width), int(header.Height))
		fmt.Println(data[header.StartingAddress:])
		fmt.Println(pixelMap)
		fmt.Println(newMap)

		// Make new image

		// newFilter := bitmap.Filt(data[header.StartingAddress:], int(header.Width), int(header.BitsPerPixel), readData)
		// newImg := bitmap.MirrorVertical(data[header.StartingAddress:], int(header.Width), int(header.Height), int(header.BitsPerPixel))
		// fmt.Println(data[header.StartingAddress:])
		// fmt.Println(newImg)
		// Crete new BMP file
		// bitmap.CreateBMP(&header, data[header.StartingAddress:], "output.bmp")
		// bitmap.CreateBMP(&header, horiz, "outputFilter.bmp")
	} else {
		fmt.Fprintln(os.Stderr, "Error apply")
		os.Exit(1)
	}
}

func errNil(err error) {
	if err != nil {
		panic(err)
	}
}
