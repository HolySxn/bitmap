package bitmap

import (
	"os"
)

func CreateBMP(head *Header, image []byte, name string) {
	newHead := head.ToBytes()

	file, err := os.Create(name)
	errNil(err)
	defer file.Close()

	_, err = file.Write(newHead)
	errNil(err)

	_, err = file.Write(image)
	errNil(err)
}
