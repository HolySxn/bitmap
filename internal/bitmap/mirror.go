package bitmap

func MirrorHorizontal(data []byte, width int, height int, bpp int) []byte {
	bit := bpp / 8
	rowSize := width * bit
	padding := 0

	newData := make([]byte, len(data))
	copy(newData, data)

	for y := 0; y < height; y++ {
		start := y * (rowSize + padding)
		end := start + rowSize
		for i, j := start, end-bit; i < j; i, j = i+bit, j-bit {
			copy(newData[i:i+bit], data[j:j+bit])
			copy(newData[j:j+bit], data[i:i+bit])
		}
		padding = (4 - (rowSize % 4)) % 4
	}

	return newData
}

func MirrorVertical(data []byte, width int, height int, bpp int) []byte {
	bit := bpp / 8
	rowSize := width * bit
	padding := (4 - (rowSize % 4)) % 4
	rowSize += padding

	newData := make([]byte, len(data))
	copy(newData, data)

	for y := 0; y < height; y++ {
		top := y * rowSize
		bottom := height*rowSize - rowSize - top
		if bottom < top{
			break
		}
		for x := 0; x < rowSize; x++ {
			newData[top+x], newData[bottom+x] = newData[bottom+x], newData[top+x]
		}
	}

	return newData
}