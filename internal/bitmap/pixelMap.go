package bitmap

type Pixel struct {
	Blue  byte
	Green byte
	Red   byte
}

func PixelMap(image []byte, width, height int, bpp int) [][]Pixel {
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
			// pixel := make([]byte, bit)
			// copy(pixel, image[i:i+bit])
			newRow = append(newRow, Pixel{image[i], image[i+1], image[i+2]})

			i += bit
		}
		// After processing each row, skip the padding bytes
		pixelMap = append(pixelMap, newRow)
		i += padding
	}

	return pixelMap
}

// func ToBytes(pixelMap [][]Pixel, bbp int) []byte{
// 	data := []byte{}
// 	row

// 	for _, row := range pixelMap{
// 		for _, pixel := range row{
// 			data = append(data, pixel.Blue)
// 			data = append(data, pixel.Green)
// 			data = append(data, pixel.Red)
// 		}
// 		if len(data) %
// 	}

// }
