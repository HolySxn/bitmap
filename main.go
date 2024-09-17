package main

import (
	"bitmap/internal/bitmap"
	"fmt"
	"os"
)

func main() {
	// error handling
	if len(os.Args) < 2 {
		fmt.Println(os.Stderr, "ERROR: Emtpty input")
		os.Exit(1)
	}
	if os.Args[1] != "apply" {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid apply")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Println(os.Stderr, "ERROR: Empty flag")
		os.Exit(1)
	}

	readData := os.Args[1:]

	var color string
	piece := readData[1]
	for i := 0; i < len(readData[1]); i++ {
		if piece[i] == '=' {
			color = piece[i+1:]
			piece = piece[:i]
			break
		}
	}
	if piece == "--filter" {
		dataPic := readFile("sample.bmp")
		bmp, _ := bitmap.Decode(dataPic)
		bmp.Filt(color)
		dataPic, _ = bitmap.Encode(bmp)
		createFile(dataPic)
		os.Exit(0)
	}
	// Read file
	data := readFile("sample.bmp")

	// Get BMPFile struct
	bmp, _ := bitmap.Decode(data)
	bmp.HeaderInfo()

	// Some manipulations
	bmp.Crop()

	// Get array of bytes
	data, _ = bitmap.Encode(bmp)

	// Create new bmp file
	createFile(data)
}
