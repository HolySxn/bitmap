package bitmap

import (
	"fmt"
	"os"
)

func errNil(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
