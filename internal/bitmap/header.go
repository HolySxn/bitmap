package bitmap

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Header struct {
	FileType          [2]byte
	FileSize          uint32
	Reserved1         uint16
	Reserved2         uint16
	StartingAddress   uint32
	HeaderSize        uint32
	Width             uint32
	Height            uint32
	ColorPlanes       uint16
	BitsPerPixel      uint16
	CompressionMethod uint32
	ImageSize         uint32
	XPixelsPerMeter   int32
	YPixelsPerMeter   int32
	TotalColors       uint32
	ImportantColors   uint32
}

func (head *Header) ReadHeader(data []byte) {
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, head)
	errNil(err)
}

func (head *Header) HeaderInfo() {
	// Check if it's a BMP file by looking at the signature ("BM")
	fmt.Println("BMP Header:")
	if string(head.FileType[:]) != "BM" {
		fmt.Println("Not a BMP file")
		return
	}

	// Extract file type
	fmt.Println("- FileType", string(head.FileType[:]))

	// Extract total file size (optional, just as an example)
	fmt.Println("- FileSizeInBytes", head.FileSize)

	// Extract the size of the header
	fmt.Println("- HeaderSize", head.StartingAddress)

	fmt.Println("DIB Header:")

	// Extract size of the DIB header
	fmt.Println("- DIBHeaderSize", head.HeaderSize)

	// Extract image width
	fmt.Println("- WidthInPixels", head.Width)

	// Extract image height
	fmt.Println("- HeightInPixels", head.Height)

	// Extract pixel size in bits
	fmt.Println("- PixelSizeInBits", head.BitsPerPixel)

	// Extract image size in bytes
	fmt.Println("- ImageSizeInBytes", head.FileSize-head.StartingAddress)
}

func (head *Header) ToBytes() []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, head)
	errNil(err)

	return buf.Bytes()
}
