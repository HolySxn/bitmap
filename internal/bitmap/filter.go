package bitmap

import "errors"

func (bmp *BMPFile) Filt(color string) error {
	data1 := bmp.image

	switch color {

	case "blue":
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				data1[i][j].g = 0
				data1[i][j].r = 0
			}
		}

	case "red":
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				data1[i][j].g = 0
				data1[i][j].b = 0
			}
		}

	case "green":
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				data1[i][j].b = 0
				data1[i][j].r = 0
			}
		}

	case "negative":
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				data1[i][j].g = 255 - data1[i][j].g
				data1[i][j].r = 255 - data1[i][j].r
				data1[i][j].b = 255 - data1[i][j].b
			}
		}

	case "pixelate":
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

				for by := 0; by < blockSize; by++ {
					for bx := 0; bx < blockSize; bx++ {
						data1[y+by][x+bx].r = RAvg
						data1[y+by][x+bx].g = GAvg
						data1[y+by][x+bx].b = BAvg
					}
				}
			}
		}

	case "blur":
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := -1; i < height-1; i++ {
			for j := -1; j < width-1; j++ {
				var rSum, gSum, bSum, count int

				for bi := 1; bi <= 21; bi++ {
					for bj := 1; bj <= 21; bj++ {
						if i+bi < height && j+bj < width {
							rSum += int(data1[i+bi][j+bj].r)
							gSum += int(data1[i+bi][j+bj].g)
							bSum += int(data1[i+bi][j+bj].b)
							count++
						}
					}
				}

				rCenterSum := rSum / count
				gCenterSum := gSum / count
				bCenterSum := bSum / count

				if i+1 >= height || j+1 >= width {
					continue
				}

				data1[i+1][j+1].r = byte(rCenterSum)
				data1[i+1][j+1].g = byte(gCenterSum)
				data1[i+1][j+1].b = byte(bCenterSum)

			}
		}

	case "grayscale":
		width := int(bmp.head.Width)
		height := int(bmp.head.Height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {

				R := float64(data1[i][j].r)
				G := float64(data1[i][j].g)
				B := float64(data1[i][j].b)

				grayscale := byte(R*0.299 + G*0.587 + B*0.114)

				data1[i][j].g = grayscale
				data1[i][j].b = grayscale
				data1[i][j].r = grayscale
			}
		}

	default:
		return errors.New("undefined filter")
	}
	return nil
}
