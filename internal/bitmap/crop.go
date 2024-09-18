package bitmap

import (
	"errors"
)

func (bmp *BMPFile) Crop(OffsetX, OffsetY, Width, Height int) error {
	data := bmp.image
	if Width == -1 && Height == -1 {
		Width = int(bmp.head.Width) - 1 - OffsetX
		Height = int(bmp.head.Height) - 1 - OffsetY
	}
	OffsetX -= 1
	OffsetY -= 1
	if OffsetX < 0 || OffsetX > int(bmp.head.Width)-Width || OffsetY < 0 || OffsetY > int(bmp.head.Height)-Height || Width < 0 || Width > int(bmp.head.Width) || Height < 0 || Height > int(bmp.head.Height) {
		return errors.New("incorrect value for --crop")
	}
	newData := [][]Pixel{}
	for i := Height + OffsetY; i > OffsetY; i-- {
		row := []Pixel{}
		for j := OffsetX; j < Width+OffsetX; j++ {
			row = append(row, data[int(bmp.head.Height)-i][j])
		}
		newData = append(newData, row)
	}
	bmp.image = newData

	bmp.head.Width, bmp.head.Height = uint32(Width), uint32(Height)

	return nil
}
