package bitmap

import (
	"encoding/binary"
	"fmt"
)

func Header(data []byte) {
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

	// Extract the size of the header
	headerSize := binary.LittleEndian.Uint32(data[10:14])
	fmt.Println("- HeaderSize", headerSize)

	fmt.Println("DIB Header:")

	// Extract size of the DIB header
	dibHeaderSize := binary.LittleEndian.Uint32(data[14:18])
	fmt.Println("- DibHeaderSize", dibHeaderSize)

	// Extract image width
	w := binary.LittleEndian.Uint32(data[18:22])
	fmt.Println("- WidthInPixels", w)

	// Extract image height
	h := binary.LittleEndian.Uint32(data[22:26])
	fmt.Println("- HeightInPixels", h)

	// Extract pixel size in bits
	pixelSize := binary.LittleEndian.Uint32(data[28:32])
	fmt.Println("- PixelSizeInBits", pixelSize)

	// Extract image size in bytes
	imageSize := fileSize - headerSize
	fmt.Println("- ImageSizeInBytes", imageSize)
}
