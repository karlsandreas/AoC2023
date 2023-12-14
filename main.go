package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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
	sum := 0
	for _, l := range lines {
		first, last, found := findFirstLastDigit(l)
		if !found {
			fmt.Println("No digit found in line", l)
		}
		twoDigit := string(first) + string(last)
		x, err := strconv.Atoi(twoDigit)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += x
	}

	fmt.Println("The sum is:", sum)

}

func readFile2Lines(fileName string) ([]string, error) {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return lines, err
	}

	return lines, nil
}

func findFirstLastDigit(str string) (rune, rune, bool) {
	var first rune
	var last rune



	fmt.Println("Input", str)
	for _, char := range str {
		if unicode.IsDigit(char) {
			if first == 0 {
				first = char
			}
			last = char

		}
	}

	if first == 0 || last == 0 {
		return first, last, false
	}
	fmt.Println("First", string(first))
	fmt.Println("Last", string(last))
	return first, last, true

}
