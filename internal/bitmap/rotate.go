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

func (bmp *BMPFile) RotateLeft() {
	data := bmp.image
	width := int(bmp.head.Width)
	height := int(bmp.head.Height)
	newData := [][]Pixel{}

	for i := 0; i < width; i++ {
		row := []Pixel{}
		for j := height - 1; j >= 0; j-- {
			row = append(row, data[j][i])
		}
		newData = append(newData, row)
	}
	bmp.image = newData

	bmp.head.Width, bmp.head.Height = bmp.head.Height, bmp.head.Width
}

func (bmp *BMPFile) BottomUp() {
	data := bmp.image
	height := int(bmp.head.Height)
	width := int(bmp.head.Width)
	newData := [][]Pixel{}

	for i := height - 1; i >= 0; i-- {
		row := []Pixel{}
		for j := width - 1; j >= 0; j-- {
			row = append(row, data[i][j])
		}
		newData = append(newData, row)
	}

	bmp.image = newData
}
