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
		data := readFile("sample.bmp")
		// Get BMPFile struct
		bmp := bitmap.Decode(data)
		bmp.HeaderInfo()

		// Some manipulations
		bmp.MirrorVertical()
		bmp.RotateRight()

		// Get array of bytes
		data = bitmap.Encode(bmp)

		createFile(data)
	} else {
		fmt.Fprintln(os.Stderr, "Error apply")
		os.Exit(1)
	}
}

func readFile(name string) []byte {
	// Open the BMP file
	file, err := os.Open(name)
	errNil(err)
	defer file.Close()

	// Read the entire file into memory
	data, err := io.ReadAll(file)
	errNil(err)

	return data
}

func createFile(data []byte) {
	file, err := os.Create("output.bmp")
	errNil(err)
	defer file.Close()

	_, err = file.Write(data)
	errNil(err)
}

func errNil(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
