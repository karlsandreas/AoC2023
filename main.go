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
		twoDigit := first + last
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

func findFirstLastDigit(str string) (string, string, bool) {
	var first string
	var last string

	wordNumber := make(map[string]string)

	wordNumber["zero"] = "0"
	wordNumber["one"] = "1"
	wordNumber["two"] = "2"
	wordNumber["three"] = "3"
	wordNumber["four"] = "4"
	wordNumber["five"] = "5"
	wordNumber["six"] = "6"
	wordNumber["seven"] = "7"
	wordNumber["eight"] = "8"
	wordNumber["nine"] = "9"

	fmt.Println("Input", str)

	strLen := len(str)

	for i, char := range str {
		if unicode.IsDigit(char) {
			if first == "" {
				first = string(char)
			}
			last = string(char)

		} else {
			for key, value := range wordNumber {
				keyLen := len(key)
				if i+keyLen <= strLen {
					if str[i:i+keyLen] == key {
						if first == "" {
							first = value
						}
						last = value
					}
				}
			}
		}
	}

	if first == "" || last == "" {
		return first, last, false
	}
	fmt.Println("First", string(first))
	fmt.Println("Last", string(last))
	return first, last, true

}
