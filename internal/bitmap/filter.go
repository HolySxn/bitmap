package bitmap

import (
	"fmt"
	"os"
)

func Filt(data []byte, width int, BitsPerPixel int, readData []string, height int, pixel [][]Pixel) [][]Pixel {
	if len(readData) <= 2 {
		os.Exit(1)
	}
	piece := readData[2]
	for i := 0; i < len(readData[2]); i++ {
		if piece[i] == '=' {
			piece = piece[i+1:]
			break
		}
	}

	if piece == "blue" {
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				pixel[row][col].red = 0
				pixel[row][col].green = 0
			}
		}
		return pixel
	} else if piece == "red" {
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				pixel[row][col].green = 0
				pixel[row][col].blue = 0
			}
		}
		return pixel
	} else if piece == "green" {
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				pixel[row][col].red = 0
				pixel[row][col].blue = 0
			}
		}
		return pixel
	} else if piece == "negative" {
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				pixel[row][col].red = 255 - pixel[row][col].red
				pixel[row][col].green = 255 - pixel[row][col].green
				pixel[row][col].blue = 255 - pixel[row][col].blue
			}
		}
		return pixel
	} else if piece == "pixelate" {
		blockSize := 20
		for row := 0; row < height; row += blockSize {
			for col := 0; col < width; col += blockSize {
				var red, green, blue, count int
				for y := 0; y < blockSize; y++ {
					for x := 0; x < blockSize; x++ {
						yPixel := row + y
						xPixel := col + x
						if yPixel > height || xPixel > width {
							continue
						}
						red += int(pixel[yPixel][xPixel].red)
						blue += int(pixel[yPixel][xPixel].green)
						green += int(pixel[yPixel][xPixel].blue)
						count++
					}
				}
				for y := 0; y < blockSize; y++ {
					for x := 0; x < blockSize; x++ {
						yPixel := row + y
						xPixel := col + x
						if yPixel > height || xPixel > width {
							continue
						}
						pixel[yPixel][xPixel].red = byte(red / count)
						pixel[yPixel][xPixel].green = byte(green / count)
						pixel[yPixel][xPixel].blue = byte(blue / count)
					}
				}
			}
		}
		return pixel
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
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				var red, green, blue, count int

				for y := 0; y <= halfKernel; y++ {
					for x := 0; x <= halfKernel; x++ {
					}
				}
			}
		}
		return pixel
	} else {
		fmt.Fprintln(os.Stderr, "Undefined filter")
		os.Exit(1)
	}
	return [][]Pixel{}
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
