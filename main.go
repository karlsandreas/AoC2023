package main

import (
	"fmt"
	"os"
)

func main() {
	var lines []string
	for i, arg := range os.Args {
		if arg == "-f" {
			var err error
			lines, err = readFile2Lines(os.Args[i+1])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	d4p2(lines)

}
