package bitmap

import (
	"fmt"
	"os"
)

func Filt(data []byte, width int, BitsPerPixel int, readData []string, height int) []byte {
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
		bit := BitsPerPixel / 8
		rowSize := width * bit
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
			if i+rowSize == 0 {
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
			if i+rowSize == 0 {
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
			if i+rowSize == 0 {
				i += padding
			} else {
				i += 3
			}
		}
		return newData
	} else if piece == "pixelate" {
		bit := BitsPerPixel / 8
		rowSize := width * bit
		blockSize := 20
		newData := make([]byte, len(data))
		copy(newData, data)

		// Process the image by pixelating in blocks
		for y := 0; y < len(data)/rowSize; y += blockSize {
			for x := 0; x < width; x += blockSize {
				// Average the pixel colors within the block
				var rSum, gSum, bSum, count int

				// Traverse each pixel in the block
				for by := 0; by < blockSize; by++ {
					for bx := 0; bx < blockSize; bx++ {
						// Calculate the pixel position in the image
						yi := y + by
						xi := x + bx

						// Make sure we don't exceed the image boundaries
						if yi >= len(data)/rowSize || xi >= width {
							continue
						}

						// Calculate the position in the byte array
						i := yi*rowSize + xi*bit

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
						if yi >= len(data)/rowSize || xi >= width {
							continue
						}

						// Calculate the position in the byte array
						i := yi*rowSize + xi*bit

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
		bit := BitsPerPixel / 8
		rowSize := width * bit
		halfKernel := 20

		newData := make([]byte, len(data))
		copy(newData, data)

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				var rSum, gSum, bSum, count int

				for ky := 0; ky <= halfKernel; ky++ {
					for kx := 0; kx <= halfKernel; kx++ {
						ny := clamp(y+ky, 0, height-1)
						nx := clamp(x+kx, 0, width-1)

						i := (ny*rowSize + nx*3)

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

				i := (y*rowSize + x*3)
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
