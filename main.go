package main

import (
	"os"
)

func main() {
	readData := os.Args[1:]
	manage(readData)
}
