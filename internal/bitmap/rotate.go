package bitmap

func Rotate90(data [][]Pixel, width int, height int) [][]Pixel {
	// Initialize the new matrix with swapped dimensions
	newData := [][]Pixel{}

	// Populate the new matrix with rotated values
	for i := 0; i < width; i++ {
		row := []Pixel{}
		for j := height - 1; j >= 0; j-- {
			row = append(row, data[i][j])
		}
		newData = append(newData, row)
	}

	return newData
}

// switch direction{
// case "90","Right","-270":
// 	return RotateRight(data[]byte)
// case "-90","Left","270":
// 	return RotateLeft(data[]byte)
// case "180","-180":
// 	return BottomUp(data[]byte)
// case "360","0","-360":
// 	return data[]byte
// }

// func Rotate90(data [][]byte, width int, height int) {
// 	Newdata := [][]byte{}
// 	for i := 0; i < len(data[0]); i++ {
// 		collumn := []byte{}
// 		for j := 0; j < len(data); j++ {
// 			pixel := data[len(data)-1-j][len(data[0])-1-i]
// 			collumn = append(collumn, pixel)
// 		}
// 		Newdata = append(Newdata, collumn)
// 	}
// 	fmt.Println(data)
// 	fmt.Println(Newdata)
// }

// package bitmap

// func Rotate90(data []byte, width int, height int, bpp int) []byte {
// 	bit := bpp / 8
// 	rowSize := width * bit
// 	padding := (4 - (rowSize % 4)) % 4
// 	rowSize += padding

// 	newRowSize := height * bit
// 	newPadding := (4 - (newRowSize % 4)) % 4
// 	newRowSize += newPadding

// 	newData := make([]byte, width*newRowSize)

// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			oldIndex := y*rowSize + x*bit

// 			newX := height - 1 - y
// 			newY := x
// 			newIndex := newY*newRowSize + newX*bit

// 			copy(newData[newIndex:newIndex+bit], data[oldIndex:oldIndex+bit])
// 		}
// 	}

// 	return newData
// }
