package bitmap

import (
	"fmt"
	"os"
)


func (bmp *BMPFile) Filt(piece string) {
	data1 := bmp.image

	if piece == "blue" {
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				data1[i][j].g = 0
				data1[i][j].r = 0
			}
		}
	} else if piece == "red" {
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				data1[i][j].g = 0
				data1[i][j].b = 0
			}
		}
	} else if piece == "green" {
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				data1[i][j].b = 0
				data1[i][j].r = 0
			}
		}
	} else if piece == "negative" {
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				data1[i][j].g = 255 - data1[i][j].g
				data1[i][j].r = 255 - data1[i][j].r
				data1[i][j].b = 255 - data1[i][j].b
			}
		}
	} else if piece == "pixelate" {
		bit := bmp.head.BitsPerPixel / 8
		rowSize := uint16(bmp.head.Width) * bit
		blockSize := 20
		newData := make([]byte, len(data1))

		// Process the image by pixelating in blocks
		for y := 0; y < len(data1)/int(rowSize); y += blockSize {
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
						if yi >= len(data1)/int(rowSize) || xi >= int(bmp.head.Width) {
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
						if yi >= len(data1)/int(rowSize) || xi >= int(bmp.head.Width) {
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
	} else if piece == "blur" {
		bit := bmp.head.BitsPerPixel / 8
		rowSize := uint16(bmp.head.Width) * bit
		halfKernel := 20

		newData := make([]byte, len(data1))

		for y := 0; y < int(bmp.head.Height); y++ {
			for x := 0; x < int(bmp.head.Width); x++ {
				var rSum, gSum, bSum, count int

				for ky := 0; ky <= halfKernel; ky++ {
					for kx := 0; kx <= halfKernel; kx++ {
						// ny := clamp(y+ky, 0, int(bmp.head.Height)-1)
						// nx := clamp(x+kx, 0, int(bmp.head.Width)-1)

						// i := (ny*int(rowSize) + nx*3)

						// b := int(data1[i])
						// g := int(data1[i+1])
						// r := int(data1[i+2])

						// rSum += r
						// gSum += g
						// bSum += b
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
	} else {
		fmt.Fprintln(os.Stderr, "Undefined filter")
		os.Exit(1)
	}
}
