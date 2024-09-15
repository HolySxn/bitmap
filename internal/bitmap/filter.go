package bitmap

import (
	"fmt"
	"os"
)

func (bmp *BMPFile) Filt(piece string, data []byte) []byte {
	
	if piece == "blue" {
		bit := int(bmp.head.BitsPerPixel) / 8
		rowSize := int(bmp.head.Width) * int(bit)
		padding := (4 - (rowSize % 4)) % 4
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
			if i+rowSize == 0 {
				i += padding
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "red" {
		bit := bmp.head.BitsPerPixel / 8
		rowSize := uint16(bmp.head.Width) * bit
		padding := (4 - (rowSize % 4)) % 4
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
			if uint16(i)+rowSize == 0 {
				i += int(padding)
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "green" {
		bit := bmp.head.BitsPerPixel / 8
		rowSize := uint16(bmp.head.Width) * bit
		padding := (4 - (rowSize % 4)) % 4
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
			if uint16(i)+rowSize == 0 {
				i += int(padding)
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "negative" {
		bit := bmp.head.BitsPerPixel / 8
		rowSize := uint16(bmp.head.Width) * bit
		padding := (4 - (rowSize % 4)) % 4
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
			if uint16(i)+rowSize == 0 {
				i += int(padding)
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "pixelate" {
		bit := bmp.head.BitsPerPixel / 8
		rowSize := uint16(bmp.head.Width) * bit
		blockSize := 20
		newData := make([]byte, len(data))
		copy(newData, data)

		// Process the image by pixelating in blocks
		for y := 0; y < len(data)/int(rowSize); y += blockSize {
			for x := 0; x < int(bmp.head.Width); x += blockSize {
				// Average the pixel colors within the block
				var rSum, gSum, bSum, count int

				// Traverse each pixel in the block
				for by := 0; by < blockSize; by++ {
					for bx := 0; bx < blockSize; bx++ {
						// Calculate the pixel position in the image
						yi := y + by
						xi := x + bx

						// Make sure we don't exceed the image boundaries
						if yi >= len(data)/int(rowSize) || xi >= int(bmp.head.Width) {
							continue
						}

						// Calculate the position in the byte array
						i := yi*int(rowSize) + xi*int(bit)

						// Extract pixel values (BMP is stored in BGR format)
						b := int(newData[i])
						g := int(newData[i+1])
						r := int(newData[i+2])

						// Sum the pixel values
						rSum += r
						gSum += g
						bSum += b
						count++
					}
				}

				// Calculate the average color
				rAvg := byte(rSum / count)
				gAvg := byte(gSum / count)
				bAvg := byte(bSum / count)

				// Apply the average color to all pixels in the block
				for by := 0; by < blockSize; by++ {
					for bx := 0; bx < blockSize; bx++ {
						// Calculate the pixel position in the image
						yi := y + by
						xi := x + bx

						// Make sure we don't exceed the image boundaries
						if yi >= len(data)/int(rowSize) || xi >= int(bmp.head.Width) {
							continue
						}

						// Calculate the position in the byte array
						i := yi*int(rowSize) + xi*int(bit)

						// Set the pixel values to the average color
						newData[i] = bAvg   // Set blue
						newData[i+1] = gAvg // Set green
						newData[i+2] = rAvg // Set red
					}
				}
			}
		}
		return newData
	} else if piece == "blur" {
		bit := bmp.head.BitsPerPixel / 8
		rowSize := uint16(bmp.head.Width) * bit
		halfKernel := 20

		newData := make([]byte, len(data))
		copy(newData, data)

		for y := 0; y < int(bmp.head.Height); y++ {
			for x := 0; x < int(bmp.head.Width); x++ {
				var rSum, gSum, bSum, count int

				for ky := 0; ky <= halfKernel; ky++ {
					for kx := 0; kx <= halfKernel; kx++ {
						ny := clamp(y+ky, 0, int(bmp.head.Height)-1)
						nx := clamp(x+kx, 0, int(bmp.head.Width)-1)

						i := (ny*int(rowSize) + nx*3)

						b := int(data[i])
						g := int(data[i+1])
						r := int(data[i+2])

						rSum += r
						gSum += g
						bSum += b
						count++
					}
				}

				rAvg := byte(rSum / count)
				gAvg := byte(gSum / count)
				bAvg := byte(bSum / count)

				i := (y*int(rowSize) + x*3)
				newData[i] = bAvg
				newData[i+1] = gAvg
				newData[i+2] = rAvg
			}
		}
		return newData
	} else {
		fmt.Fprintln(os.Stderr, "Undefined filter")
		os.Exit(1)
	}
	return []byte{}
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
