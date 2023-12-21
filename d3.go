package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type gear struct {
	num   int
	start int
	stop  int
}

type gearRatio struct {
	a     *gear
	b     *gear
	ratio int
}

func d3p2(lines []string) {
	//var nums []*gearRatio
	sum := 0

	for i, l := range lines {
		for j, c := range l {
			if c == '*' {
				gr, success := findAdjacentsGear(lines, i, j)
				if success {
					sum += gr.ratio
				}
			}
		}
	}
	fmt.Println("Gear ratio", sum)
}

func findAdjacentsGear(lines []string, row int, col int) (*gearRatio, bool) {
	gear1 := gear{}
	gear2 := gear{}
	g := &gear1
	gr := &gearRatio{a: &gear1, b: &gear2}

	count := 0

	for i := max(row-1, 0); i <= min(row+1, len(lines)-1); i++ {
		l := lines[i]
		for j := max(col-1, 0); j <= min(col+1, len(lines[i])-1); j++ {
			if unicode.IsDigit(rune(l[j])) {
				num, start, stop, success := findNum(l, j)
				if success {
					count += 1
					if count > 2 {
						fmt.Println("Too many adjacent nums")
						return gr, false
					}
					j = stop
					numInt, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println("Error converting to int")
						return gr, false
					}
					g.num = numInt
					g.start = start
					g.stop = stop
					g = &gear2
				}
			}
		}
	}
	gr.ratio = gr.a.num * gr.b.num

	return gr, true

}

func d3p1(lines []string) {
	nums := make([][]string, len(lines))
	sum := 0

	for i, l := range lines {
		j := 0
		for {
			//fmt.Println(string(l[j]), i, l)
			//fmt.Println(unicode.IsDigit(rune(l[j])))
			if unicode.IsDigit(rune(l[j])) {
				num, start, stop, success := findNum(l, j)
				fmt.Println("Processing", num)
				fmt.Println(i, j)
				if success {
					if isSymbolAdjacent(lines[max(i-1, 0):1+min(i+1, len(lines))], num, start, stop) {
						nums[i] = append(nums[i], num)
						numInt, err := strconv.Atoi(num)
						if err != nil {
							fmt.Println("Error in string to in conversion")
							return
						}
						fmt.Println("Adding", numInt, "to sum")
						sum += numInt
					}
					j = stop
				}

			}
			j++
			if j >= len(l)-1 {
				break
			}
		}
	}
	fmt.Println("Sum", sum)
}

func isSymbolAdjacent(lines []string, num string, start int, stop int) bool {

	for _, l := range lines {
		maxJ := max(start-1, 0)
		minJ := min(stop+1, len(l)-1)
		for j := maxJ; j <= minJ; j++ {
			if (unicode.IsPunct(rune(l[j])) || unicode.IsSymbol(rune(l[j]))) && l[j] != '.' {
				return true
			}
		}
	}
	return false
}

// From an index, find if there is a number at that index and return the whole number
func findNum(input string, index int) (string, int, int, bool) {
	start := index
	stop := index
	for {
		if start > 0 {
			//Find start index of num
			if unicode.IsDigit(rune(input[start])) {
				//Test one charater to the left
				start -= 1
			} else {
				start += 1
				break
			}
		} else {
			start = 0
			if unicode.IsDigit(rune(input[start])) {
				break
			}
			start += 1
			break
		}
	}
	for {
		if stop < len(input) {
			//find stop index of num
			if unicode.IsDigit(rune(input[stop])) {
				stop += 1
			} else {
				stop -= 1
				break
			}
		} else {
			break
		}
	}
	//fmt.Println("Start", start, " stop", stop)
	stop = min(stop, len(input)-1)
	//fmt.Println(string(input[start]), string(input[stop]))

	if unicode.IsDigit(rune(input[start])) && unicode.IsDigit(rune(input[stop])) {
		//fmt.Println("Found num", input[start:stop+1], "in input", input)
		return input[start : stop+1], start, stop, true

	}
	return "", 0, 0, false
}
