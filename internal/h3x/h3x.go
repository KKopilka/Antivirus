package h3x

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func hexdump(s string) ([]byte, error) {
	content, err := ioutil.ReadFile(s)
	if err != nil {
		panic("Can't open file")
	}

	return content, nil
}

func resultOutput(arr []byte) {
	row_width := 16
	curr_row_width := 16
	seen_bytes := 0

	for i := 0; i < len(arr); i += row_width {
		if (len(arr) - seen_bytes) < row_width {
			curr_row_width = (len(arr) - seen_bytes)
		}
		row := arr[i:(seen_bytes + curr_row_width)]

		for i := 0; i < row_width; i++ {
			if i < curr_row_width {
				fmt.Printf("%02x ", row[i])
			} else {
				fmt.Print(strings.Repeat(" ", 3))
			}
		}

		fmt.Print("|")
		fmt.Print(" ")

		for i := 0; i < row_width; i++ {
			if i < curr_row_width {
				if row[i] >= 0x20 && row[i] < 0x7f {
					fmt.Print(string(row[i]))
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(strings.Repeat(" ", 3))
			}
		}
		fmt.Print("|")
		fmt.Print("\n")

		seen_bytes += row_width
	}
}

// func main() {
// 	if len(os.Args) <= 1 {
// 		fmt.Println("Usage: h3x <FILEPATH>")
// 		os.Exit(1)
// 	}

// 	var content, _ = hexdump(os.Args[1])
// 	resultOutput(content)
// }
