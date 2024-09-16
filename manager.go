package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"bitmap/internal/bitmap"
)

var bitmapHelp = `Usage:
	bitmap <command> [arguments]

The commands are:
	header    prints bitmap file header information
	apply     applies processing to the image and saves it to the file

The options are:
	-h, --help      prints program usage information`

var headerHelp = `Usage:
	bitmap header <source_file>

Description:
	Prints bitmap file header information
	
The options are:
	-h, --help      prints program usage information`

var applyHeader = `Usage:
	bitmap apply [options] <source_file> <output_file>

File options:
	source_file and output_file must be .bmp type

The options are:
	-h, --help      prints program usage information
	...`

func manage(args []string) {
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		fmt.Println(bitmapHelp)
	} else if len(args) >= 1 {
		switch args[0] {
		case "header":
			if len(args[1:]) == 0 || len(args) > 2 {
				fmt.Println(headerHelp)
				os.Exit(1)
			} else if args[1] == "--help" || args[1] == "-h" {
				fmt.Println(headerHelp)
				os.Exit(0)
			} else if strings.HasPrefix(args[1], "--") || strings.HasPrefix(args[1], "-") {
				fmt.Println("unknown option:", args[1])
				fmt.Println(headerHelp)
				os.Exit(1)
			} else {
				data := readFile(args[1])
				bmp, err := bitmap.Decode(data)
				errNil(err)

				bmp.HeaderInfo()
			}
		case "apply":
			if len(args[1:]) < 3{
				fmt.Println("error: not enough arguments")
				fmt.Println(applyHeader)
				os.Exit(1)
			}else{
				dst := args[len(args)-1]
				src := args[len(args)-2]

				if !strings.HasSuffix(dst, ".bmp") || !strings.HasSuffix(src, ".bmp") {
					fmt.Println(src)
					fmt.Println(dst)
					fmt.Println("wrong file type")
					fmt.Println(applyHeader)
					os.Exit(1)
				}

				args = args[1:len(args)-2]

				applyManager(args, src, dst)
			}
		}
	}
}

func applyManager(args []string, src, dst string){
	fmt.Println("Did not do yet")
	fmt.Println(args)
	fmt.Println(src)
	fmt.Println(dst)
	os.Exit(0)
}

func readFile(name string) []byte {
	// Open the BMP file
	file, err := os.Open(name)
	errNil(err)
	defer file.Close()

	// Read the entire file into memory
	data, err := io.ReadAll(file)
	errNil(err)

	return data
}

func createFile(data []byte) {
	file, err := os.Create("output.bmp")
	errNil(err)
	defer file.Close()

	_, err = file.Write(data)
	errNil(err)
}

func errNil(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
