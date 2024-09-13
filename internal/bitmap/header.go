package bitmap

import (
	"fmt"
)

func (bmp *BMPFile) HeaderInfo() {
	fmt.Println("BMP Header:")

	// Display basic header information
	fmt.Printf("- File Type: %s\n", string(bmp.head.FileType[:]))
	fmt.Printf("- File Size: %d bytes\n", bmp.head.FileSize)
	fmt.Printf("- Starting Address: %d\n", bmp.head.StartingAddress)

	fmt.Println("DIB Header:")

	// Display DIB header information
	fmt.Printf("- DIB Header Size: %d bytes\n", bmp.head.HeaderSize)
	fmt.Printf("- Image Width: %d pixels\n", bmp.head.Width)
	fmt.Printf("- Image Height: %d pixels\n", bmp.head.Height)
	fmt.Printf("- Bits Per Pixel: %d\n", bmp.head.BitsPerPixel)
	fmt.Printf("- Image Data Size: %d bytes\n", bmp.head.FileSize-bmp.head.StartingAddress)
}
