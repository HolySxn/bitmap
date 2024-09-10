package bitmap

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type BMPHeader struct {
	FileType [2]byte
	FileSize uint32
	Reserved1 uint16
	Reserved2 uint16
	StartingAddress uint32
}

type DIBHeader struct{
	HeaderSize uint32
	Width uint32
	Height uint32
	ColorPlanes uint16
	BitsPerPixel uint16 
	CompressionMethod uint32
	ImageSize uint32
	XPixelsPerMeter   int32
	YPixelsPerMeter   int32
	TotalColors       uint32
	ImportantColors   uint32
}

type Header struct{
	BMPHeader BMPHeader
	DIBHeader DIBHeader
}

func (head *Header) ReadHeader(data []byte) {
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, head)
	errNil(err)
	// Check if it's a BMP file by looking at the signature ("BM")
	fmt.Println("BMP Header:")
	if string(head.BMPHeader.FileType[:]) != "BM" {
		fmt.Println("Not a BMP file")
		return
	}

	// Extract file type
	fmt.Println("- FileType", string(head.BMPHeader.FileType[:]))

	// Extract total file size (optional, just as an example)
	fmt.Println("- FileSizeInBytes", head.BMPHeader.FileSize)

	// Extract the size of the header
	fmt.Println("- HeaderSize", head.BMPHeader.StartingAddress)

	fmt.Println("DIB Header:")

	// Extract size of the DIB header
	fmt.Println("- DIBHeaderSize", head.DIBHeader.HeaderSize)

	// Extract image width
	fmt.Println("- WidthInPixels", head.DIBHeader.Width)

	// Extract image height
	fmt.Println("- HeightInPixels", head.DIBHeader.Height)

	// Extract pixel size in bits
	fmt.Println("- PixelSizeInBits", head.DIBHeader.BitsPerPixel)

	// Extract image size in bytes
	fmt.Println("- ImageSizeInBytes", head.BMPHeader.FileSize-head.BMPHeader.StartingAddress)
}
