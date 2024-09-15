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
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)
		blockSize := 20
		
		for y := 0; y < height; y += blockSize {
			for x := 0; x < width; x += blockSize {
				var RSum, GSum, BSum int
				count := 0
	
				for by := 0; by < blockSize; by++ {
					for bx := 0; bx < blockSize; bx++ {
						pixel := data1[y+by][x+bx]
						RSum += int(pixel.r)
						GSum += int(pixel.g)
						BSum += int(pixel.b)
						count++
					}
				}
	
				RAvg := byte(RSum / count)
				GAvg := byte(GSum / count)
				BAvg := byte(BSum / count)
	
				for by := 0; by < blockSize && y+by < height; by++ {
					for bx := 0; bx < blockSize && x+bx < width; bx++ {
						data1[y+by][x+bx].r = RAvg
						data1[y+by][x+bx].g = GAvg
						data1[y+by][x+bx].b = BAvg
					}
				}
			}
		}
	} else if piece == "blur" {
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)
		
	} else {
		fmt.Fprintln(os.Stderr, "Undefined filter")
		os.Exit(1)
	}

}