package bitmap

func (bmp *BMPFile) Crop() {
	data := bmp.image
	OffsetX := 20
	OffsetY := 20
	Width := 100
	Height := 100
	newData := [][]Pixel{}
	for i := Height + OffsetY; i > OffsetY; i-- {
		row := []Pixel{}
		for j := OffsetX; j < Width+OffsetX ; j++ {
			row = append(row, data[int(bmp.head.Height)-i-1][j])
		}
		newData = append(newData, row)
	}
	bmp.image = newData

	bmp.head.Width, bmp.head.Height = uint32(Width), uint32(Height)
}
