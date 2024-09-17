package bitmap

import "fmt"

func (bmp *BMPFile)Crop() {
	data := bmp.image
	OffsetX := 0
	OffsetY := 0
	Width := 10
	Height := 100
	if OffsetX < 0 || OffsetX > int(bmp.head.Width) - Width || OffsetY < 0 || OffsetY > int(bmp.head.Height) - Height || Width < 0 || Width > int(bmp.head.Width) || Height < 0 || Height > int(bmp.head.Height) {
		fmt.Println("Incorrect value for CROP")
		bmp.image = nil
		bmp.head.Width, bmp.head.Height = 1, 1	
		return
	}
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
