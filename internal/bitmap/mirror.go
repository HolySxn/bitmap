package bitmap

func (bmp *BMPFile) MirrorHorizontal() {
	// Мariable initialization
	data := bmp.image
	width := int(bmp.head.Width)
	height := int(bmp.head.Height)

	for y := 0; y < height; y++ {
		for x := 0; x < width/2; x++ {
			// Changes the opposite pixels in the row
			data[y][x], data[y][width-x-1] = data[y][width-x-1], data[y][x]
		}
	}
}

func (bmp *BMPFile) MirrorVertical() {
	// Мariable initialization
	data := bmp.image
	width := int(bmp.head.Width)
	height := int(bmp.head.Height)

	for y := 0; y < height/2; y++ {
		for x := 0; x < width; x++ {
			// Changes the opposite pixels in the column
			data[y][x], data[height-y-1][x] = data[height-y-1][x], data[y][x]
		}
	}
}
