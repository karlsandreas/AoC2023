package main

import (
	"fmt"
	"strconv"
	"strings"
)

type card struct {
	cardNum int
	wins    []string
	cards   []string
}

type cards []card

func (c cards) Len() int           { return len(c) }
func (c cards) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c cards) Less(i, j int) bool { return c[i].cardNum < c[j].cardNum }

func d4p1(lines []string) {
	sum := 0
	for _, l := range lines {
		cards := strings.Split(l, ":")[1]
		x := strings.Split(cards, "|")
		w := strings.Split(strings.TrimSpace(x[0]), " ")
		c := strings.Split(strings.TrimSpace(x[1]), " ")

		twoPower := -1
		found := false
		for _, j := range w {
			for _, k := range c {
				if j == k && (j != "" || k != "") {
					fmt.Println(j, k)
					twoPower += 1
					found = true
				}
			}
		}
		if found {
			sum += 1 << twoPower
			//fmt.Println(sum)
		}

	}
	fmt.Println(sum)
}

func d4p2(lines []string) {

	org := *linesToCards(lines)

	cs := *linesToCards(lines)

	for i := 0; i < len(cs); i++ {
		c := cs[i]
		matches := matchingCards(&cs[i])
		cs = append(cs, org[c.cardNum:c.cardNum+matches]...)
	}

	fmt.Println(len(cs))
}

func addCards(current int, n int, counts []int) *[]int {
	for i := current + 1; i <= n+current; i++ {
		counts[i] += 1
	}

	return &counts

}

func linesToCards(lines []string) *cards {
	var cs cards
	for _, l := range lines {
		c := card{}
		cards := strings.Split(l, ":")
		y := strings.Split(strings.TrimSpace(cards[0]), " ")
		num, err := strconv.Atoi(y[len(y)-1])
		if err != nil {
			fmt.Println(err)
			return &cs
		}
		c.cardNum = num
		x := strings.Split(cards[1], "|")

		c.wins = strings.Split(strings.TrimSpace(x[0]), " ")
		c.cards = strings.Split(strings.TrimSpace(x[1]), " ")

		c.wins = removeEmptyStrings(c.wins)
		c.cards = removeEmptyStrings(c.cards)
		cs = append(cs, c)
	}

	return &cs

}

func matchingCards(c *card) int {
	matches := 0
	for _, j := range c.wins {
		for _, k := range c.cards {
			if j == k && (j != "" || k != "") {
				matches += 1
			}
		}
	}

	return matches
}

func removeEmptyStrings(slice []string) []string {
	var result []string
	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
