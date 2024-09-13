package bitmap

func PixelMap(image []byte, width, height int, bpp int) [][]byte {
	bit := bpp / 8
	rowSize := width * bit
	padding := (4 - (rowSize % 4)) % 4

	pixelMap := [][]byte{}
	i := 0
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			// if i+bit > len(image) {
			// 	break
			// }
			// Extract the pixel and add it to the pixel map
			// pixel := make([]byte, bit)
			// copy(pixel, image[i:i+bit])
			pixelMap = append(pixelMap, image[i:i+bit])

			i += bit
		}
		// After processing each row, skip the padding bytes
		i += padding
	}

	return pixelMap
}
