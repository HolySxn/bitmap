package main

import (
	"bitmap/internal/bitmap"
	"fmt"
	"os"
)

func main() {
	var color string
	readData := os.Args[1:]
	if readData[0] == "apply" {
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
	} else {
		fmt.Fprintln(os.Stderr, "Error apply")
		os.Exit(1)
	}
}
