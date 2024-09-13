package bitmap

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Get data and create BMPFile struct from a byte slice
func Decode(data []byte) BMPFile {
	// Create BMPFile
	var bmp BMPFile
	bmp.head = readHeader(data)
	bmp.image = readImage(data[bmp.head.StartingAddress:], int(bmp.head.Width), int(bmp.head.Height), int(bmp.head.BitsPerPixel))

	return bmp
}

// Create new a new byte slice from a BMPFile struct
func Encode(bmp BMPFile) []byte {
	head := bmp.convertHeaderToBytes()
	body := bmp.converBodyToBytes()

	file := append(head, body...)

	return file
}

// Read information from header
func readHeader(data []byte) Header {
	var head Header

	// Read all data from Header of the file
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, &head)
	errNil(err)

	return head
}

// Read pixels from data and create [][]Pixel array
func readImage(image []byte, width, height int, bpp int) [][]Pixel {
	bit := bpp / 8
	rowSize := width * bit
	padding := (4 - (rowSize % 4)) % 4

	pixelMap := [][]Pixel{}
	i := 0
	for row := 0; row < height; row++ {
		newRow := []Pixel{}
		for col := 0; col < width; col++ {
			// if i+bit > len(image) {
			// 	break
			// }
			// Extract the pixel and add it to the pixel map
			pixel := make([]byte, bit)
			copy(pixel, image[i:i+bit])
			newRow = append(newRow, Pixel{
				b: pixel[0],
				g: pixel[1],
				r: pixel[2]})

			i += bit
		}
		// After processing each row, skip the padding bytes
		pixelMap = append(pixelMap, newRow)
		i += padding
	}

	return pixelMap
}

// ConvertHeaderToBytes converts the Header struct to a byte slice
func (bmp *BMPFile) convertHeaderToBytes() []byte {
	// Create a buffer to hold the byte data
	buf := new(bytes.Buffer)

	// Write the Header struct into the buffer using little-endian encoding
	err := binary.Write(buf, binary.LittleEndian, bmp.head)
	if err != nil {
		log.Fatalf("Failed to convert header to bytes: %v", err)
	}

	// Return the byte slice
	return buf.Bytes()
}

// Convert [][]Pixel array to a byte slice
func (bmp *BMPFile) converBodyToBytes() []byte {
	data := bmp.image
	bpp := int(bmp.head.BitsPerPixel) / 8
	rowSize := int(bmp.head.Width) * bpp
	padding := (4 - (rowSize % 4)) % 4

	body := []byte{}

	for _, row := range data {
		for _, pixel := range row {
			// Add all bytes from row into array
			body = append(body, pixel.b)
			body = append(body, pixel.g)
			body = append(body, pixel.r)
		}
		// Add left bytes into row
		body = append(body, make([]byte, padding)...)
	}

	return body
}
