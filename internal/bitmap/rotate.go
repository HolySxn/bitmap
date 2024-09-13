package bitmap

func (bmp *BMPFile) RotateRight() {
	// Initialize variables
	data := bmp.image
	width := int(bmp.head.Width)
	height := int(bmp.head.Height)
	newData := [][]Pixel{}

	// Populate the new matrix with rotated values
	for i := width - 1; i >= 0; i-- {
		row := []Pixel{}
		for j := 0; j < height; j++ {
			row = append(row, data[j][i])
		}
		newData = append(newData, row)
	}

	// Data overwrite
	bmp.image = newData
	// Height and width overwriting
	bmp.head.Width, bmp.head.Height = bmp.head.Height, bmp.head.Width
}
