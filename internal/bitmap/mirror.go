package bitmap

func MirrorHorizontal(data [][]byte, width int, height int, bpp int) {
	bpp = bpp / 8

	for y := 0; y < height; y++ {
		start := y * width
		end := start + width - 1
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			for k := 0; k < bpp; k++ {
				data[i][k], data[j][k] = data[j][k], data[i][k]
			}
		}
	}
}


func MirrorVertical(data [][]byte, width int, height int, bpp int) {
	bpp = bpp / 8

	for y := 0; y < height; y++ {
		top := y * width
		bottom := height*width - width - top
		if bottom < top {
			break
		}
		for x := 0; x < width; x++ {
			for k := 0; k < bpp; k++ {
				data[top+x][k], data[bottom+x][k] = data[bottom+x][k], data[top+x][k]
			}
		}
	}
}
