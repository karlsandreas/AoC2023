package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type cubeCount struct {
	red   int
	blue  int
	green int
}

type game struct {
	id      int
	cubeRec []cubeCount
	valid   bool
	power   int
}

func d2() {
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

	sumValid := 0
	sumPowers := 0
	games := []game{}

	for _, line := range lines {
		gameEntry := game{}
		x := strings.Split(line, ":")
		fmt.Println(x[0])
		id, err := strconv.Atoi(multiDigitInStr(x[0]))
		if err != nil {
			fmt.Println(err)
			return
		}
		gameEntry.id = id //Last element in the first split
		rounds := strings.Split(x[1], ";")
		//Count every round
		for _, r := range rounds {
			maxCubes := cubeCount{
				red:   12,
				green: 13,
				blue:  14}

			colorCounts := strings.Split(r, ",")
			//Count the value for every cube
			cube := cubeCount{}
			for _, y := range colorCounts {
				count, color, err := findColorCount(y)
				if err != nil {
					fmt.Println(err)
					return
				}

				switch color {
				case "red":
					cube.red = count
				case "blue":
					cube.blue = count
				case "green":
					cube.green = count
				}
			}
			validCube := validRound(maxCubes, cube)
			gameEntry.valid = validCube

			gameEntry.cubeRec = append(gameEntry.cubeRec, cube)

		}
		if gameEntry.valid {
			sumValid += gameEntry.id
		}
		gameEntry.power = calculateMinPower(gameEntry.cubeRec)
		fmt.Println("Power is", gameEntry.power)
		sumPowers += gameEntry.power

		games = append(games, gameEntry)
		//printGameRec(gameEntry)
	}
	fmt.Println("Sum valid is ", sumValid)
	fmt.Println("Sum power is:", sumPowers)
}

func multiDigitInStr(input string) string {
	input = strings.TrimSpace(input)

	var runes []rune

	for _, c := range input {
		if unicode.IsDigit(c) {
			runes = append(runes, c)
		}
	}
	digits := ""

	for _, r := range runes {
		digits = digits + string(r)
	}

	return digits

}

func findColorCount(input string) (int, string, error) {
	keyVal := strings.Split(strings.TrimSpace(input), " ")

	val, err := strconv.Atoi(keyVal[0])
	if err != nil {
		fmt.Println("Failed to convert ", keyVal[0], "to int:", err)
		return 0, "", err
	}

	return val, keyVal[1], nil
}

func validRound(max cubeCount, input cubeCount) bool {
	if input.red > max.red ||
		input.blue > max.blue ||
		input.green > max.green {
		return false
	}

	return true
}

func calculateMinPower(cubes []cubeCount) int {
	minCount := cubeCount{
		red:   1,
		blue:  1,
		green: 1}

	for _, c := range cubes {
		if c.red > minCount.red {
			minCount.red = c.red
		}
		if c.blue > minCount.blue {
			minCount.blue = c.blue
		}
		if c.green > minCount.green {
			minCount.green = c.green
		}
	}
	printCubeCount(minCount)

	return minCount.red * minCount.blue * minCount.green
}

func printGameRec(g game) {
	fmt.Println("ID:", g.id)
	fmt.Println("Valid:", g.valid)
	printCubeCounts(g.cubeRec)

}

func printCubeCounts(c []cubeCount) {
	for _, g := range c {
		printCubeCount(g)

	}
}

func printCubeCount(c cubeCount) {
	fmt.Println("Red:", c.red, "Blue:", c.blue, "Green:", c.green)
}
