package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the BMP file
	file, err := os.Open("sample.bmp")
	errNil(err)
	defer file.Close()

	// Read the entire file into memory
	data, err := io.ReadAll(file)
	errNil(err)

	// Check if it's a BMP file by looking at the signature ("BM")
	fmt.Println("BMP Header:")
	if string(data[:2]) != "BM" {
		fmt.Println("Not a BMP file")
		return
	}

	// Extract file type
	fmt.Println("- FileType", string(data[:2]))

	// Extract total file size (optional, just as an example)
	fileSize := binary.LittleEndian.Uint32(data[2:6])
	fmt.Println("- FileSizeInBytes", fileSize)

	// Extract the size of the DIB header (which starts at byte 14 in the file)
	headerSize := binary.LittleEndian.Uint32(data[10:14]) 
	fmt.Println("- HeaderSize", headerSize)

	
	fmt.Println("DIB Header:")

	dibHeaderSize := binary.LittleEndian.Uint32(data[14:18]) 
	fmt.Println("- DibHeaderSize", dibHeaderSize)
	
	w := binary.LittleEndian.Uint32(data[18:22])
	fmt.Println("- WidthInPixels", w)

	h := binary.LittleEndian.Uint32(data[22:26])
	fmt.Println("- HeightInPixels", h)

	pixelSize := binary.LittleEndian.Uint32(data[28:32])
	fmt.Println("- PixelSizeInBits", pixelSize)

	imageSize := fileSize - headerSize
	fmt.Println("- ImageSizeInBytes", imageSize)
}

func errNil(err error) {
	if err != nil {
		panic(err)
	}
}
