package bitmap

func Rotate90(data []byte, width int, height int, bpp int) []byte {
	tempData := make([][]byte, height)
	counter := 0
	for i := 0; i < height; i++ {
		for j := 0; j <= height; j++ {
			tempData[i] = append(tempData[i], data[counter])
			counter++
		}
	}

	RotatedMatrix := make([][]byte, width)
	for i := 0; i < width; i++ {
		RotatedMatrix[i] = make([]byte, height)
	}
	m := 0
	k := 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			RotatedMatrix[m][k] = tempData[height-1-j][i]
			k++
		}
		k = 0
		m++
	}

	var newData []byte

	for i := 0; i < height; i++ {
		newData = append(newData, RotatedMatrix[i]...)
	}
	return newData
}
