package bitmap

import (
	"fmt"
	"os"
)

func Filt(data []byte, width int, BitsPerPixel int, readData []string) []byte {
	piece := readData[2]
	for i := 0; i < len(readData[2]); i++ {
		if piece[i] == '=' {
			piece = piece[i+1:]
			break
		}
	}

	if piece == "blue" {
		bit := BitsPerPixel / 8
		rowSize := width * bit
		padding := (4 - (rowSize % 4)) % 4
		// rowSize += padding
		newData := make([]byte, len(data))
		copy(newData, data)

		i := 0
		for {
			if i > len(newData) || i+3 > len(newData) {
				break
			}
			pix := newData[i : i+3]
			pix[1] = 0
			pix[2] = 0
			if i+1%rowSize == 0 {
				i += padding
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "red" {
		bit := BitsPerPixel / 8
		rowSize := width * bit
		padding := (4 - (rowSize % 4)) % 4
		// rowSize += padding
		newData := make([]byte, len(data))
		copy(newData, data)

		i := 0
		for {
			if i > len(newData) || i+3 > len(newData) {
				break
			}
			pix := newData[i : i+3]
			pix[0] = 0
			pix[1] = 0
			if i+1%rowSize == 0 {
				i += padding
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "green" {
		bit := BitsPerPixel / 8
		rowSize := width * bit
		padding := (4 - (rowSize % 4)) % 4
		// rowSize += padding
		newData := make([]byte, len(data))
		copy(newData, data)

		i := 0
		for {
			if i > len(newData) || i+3 > len(newData) {
				break
			}
			pix := newData[i : i+3]
			pix[0] = 0
			pix[2] = 0
			if i+1%rowSize == 0 {
				i += padding
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "negative" {
		bit := BitsPerPixel / 8
		rowSize := width * bit
		padding := (4 - (rowSize % 4)) % 4
		// rowSize += padding
		newData := make([]byte, len(data))
		copy(newData, data)

		i := 0
		for {
			if i > len(newData) || i+3 > len(newData) {
				break
			}
			pix := newData[i : i+3]
			pix[0] = 255 - pix[0]
			pix[1] = 255 - pix[1]
			pix[2] = 255 - pix[2]
			if i+1%rowSize == 0 {
				i += padding
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "pixelate" {
		bit := BitsPerPixel / 8
		rowSize := width * bit
		padding := (4 - (rowSize % 4)) % 4
		// rowSize += padding
		newData := make([]byte, len(data))
		copy(newData, data)

		i := 0
		for {
			if i > len(newData) || i+3 > len(newData) {
				break
			}
			if i+1%rowSize == 0 {
				i += padding
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "blur" {

	} else {
		fmt.Fprintln(os.Stderr, "Undefined filter")
		os.Exit(1)
	}

	return []byte{}
}
