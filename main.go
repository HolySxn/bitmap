package main

import (
	"os"
)

func main() {
	//var color string
	args := os.Args[1:]
	manage(args)
	// if readData[1] == "apply" {
	// 	piece := readData[2]
	// 	for i := 0; i < len(readData[2]); i++ {
	// 		if piece[i] == '=' {
	// 			color = piece[i+1:]
	// 			piece = piece[:i]
	// 			break
	// 		}
	// 	}
	// 	if piece == "--filter" {
	// 		dataPic := readFile("sample.bmp")
	// 		bmp := bitmap.Decode(dataPic)
	// 		bmp.Filt(color)
	// 		dataPic = bitmap.Encode(bmp)
	// 		createFile(dataPic)
	// 		os.Exit(0)
	// 	}
	// 	// Read file
	// 	data := readFile("sample.bmp")

	// 	// Get BMPFile struct
	// 	bmp := bitmap.Decode(data)
	// 	bmp.HeaderInfo()

	// 	// Some manipulations
	// 	bmp.Crop()

	// 	// Get array of bytes
	// 	data = bitmap.Encode(bmp)

	// 	// Create new bmp file
	// 	createFile(data)
	// } else {
	// 	fmt.Fprintln(os.Stderr, "Error apply")
	// 	os.Exit(1)
	// }
}
