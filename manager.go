package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

var applyHelp = `Usage:
	bitmap apply [options] <source_file> <output_file>

File options:
	source_file and output_file must be .bmp type

The options are:
	-h, --help      prints program usage information

	--mirror=:
	 	[horizontal], [h], [horizontally], [hor] 	mirror the image horizontally
		[vertical], [v], [vertically], [ver] 		mirror the image vertically
	
	--filter=:
		[blue] 		apply a filter that retains only the blue channel
		[red] 		apply a filter that retains only the red channel
		[green] 	apply a filter that retains only the green channel
		[grayscale] 	convert the image to grayscale
		[negative]	apply a negative filter
		[pixelate] 	apply a pixelation effect. Option pixelates the image with a block size of 20 pixels by default
		[blur] 		apply a blur effect
		
	--rotate=:
		[right] 					rotate the image right
		[left] 						rotate the image left
		[90], [180], [270], [-90], [-180], [-270] 	rotate the image on a certain number of degree
	
	--crop=OffsetX-OffsetY-Width-Height	crop image. Width and height are optional`


func manage(args []string) {
	if len(args) == 0 {
		fmt.Println("no command provided")
		fmt.Println(bitmapHelp)
		os.Exit(1)
	} else if args[0] == "--help" || args[0] == "-h" {
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
			if len(args[1:]) == 1 && (args[1] == "--help" || args[1] == "-h"){
				fmt.Println(applyHelp)
			}else if len(args[1:]) < 3 {
				fmt.Println("error: not enough arguments. See 'apply --help'")
				os.Exit(1)
			} else {
				dst := args[len(args)-1]
				src := args[len(args)-2]

				if !strings.HasSuffix(dst, ".bmp") {
					fmt.Println("Error: <output_file> must be .bmp type")
					os.Exit(1)
				}

				args = args[1 : len(args)-2]

				file := readFile(src)
				bmp, err := bitmap.Decode(file)
				errNil(err)

				for _, command := range args {
					applyManager(&bmp, command)
				}

				newData, err := bitmap.Encode(bmp)
				errNil(err)

				// Create new bmp file
				createFile(newData, dst)
				fmt.Printf("New file %v was successfully created!\n", dst)
			}
		default:
			fmt.Println("wrong input. See '--help'")
			os.Exit(1)
		}
	}
}

func applyManager(bmp *bitmap.BMPFile, command string) {
	com_val := strings.Split(command, "=")
	if len(com_val) != 2 {
		fmt.Printf("wrong input %v. See 'apply --help'\n", command)
		os.Exit(1)
	}

	com := com_val[0]
	value := com_val[1]

	switch com {
	case "--mirror":
		switch value {
		case "horizontal", "h", "horizontally", "hor":
			bmp.MirrorHorizontal()
		case "vertical", "v", "vertically", "ver":
			bmp.MirrorVertical()
		default:
			fmt.Printf("wrong input %v. See 'apply --help'\n", command)
			os.Exit(1)
		}
	case "--filter":
		err := bmp.Filt(value)
		if err != nil {
			fmt.Printf("%v. See 'apply --help'\n", err)
			os.Exit(1)
		}
	case "--rotate":
		switch value {
		case "right", "90", "-270":
			bmp.RotateRight()
		case "left", "-90", "270":
			bmp.RotateLeft()
		case "180", "-180":
			bmp.BottomUp()
		default:
			fmt.Printf("wrong input %v. See 'apply --help'\n", command)
			os.Exit(1)
		}
	case "--crop":
		xywd := strings.Split(value, "-")
		if len(xywd) != 2 && len(xywd) != 4 {
			fmt.Printf("wrong input %v. See 'apply --help'\n", command)
			os.Exit(1)
		}

		for _, v := range xywd {
			if v == "" {
				errApply(command)
			}
		}

		offsetX, offsetY, width, height := -1, -1, -1, -1
		var err error
		offsetX, err = strconv.Atoi(xywd[0])
		if err != nil {
			errApply(command)
		}
		offsetY, err = strconv.Atoi(xywd[1])
		if err != nil {
			errApply(command)
		}

		if len(xywd) == 4 {
			width, err = strconv.Atoi(xywd[2])
			if err != nil {
				errApply(command)
			}
			height, err = strconv.Atoi(xywd[3])
			if err != nil {
				errApply(command)
			}
		}

		err = bmp.Crop(offsetX, offsetY, width, height)
		if err != nil{
			errApply(command)
		}
	default:
		errApply(command)
	}
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

func createFile(data []byte, name string) {
	file, err := os.Create(name)
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

func errApply(command string) {
	fmt.Printf("wrong input %v. See 'apply --help'\n", command)
	os.Exit(1)
}
