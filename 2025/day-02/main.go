package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/rawat9/go-utils/conv"
	"github.com/rawat9/go-utils/slice"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func solve() (int, int) {
	ranges := parseInput()
	p1 := 0
	p2 := 0

	for _, r := range ranges {
		first, last := r[0], r[1]

		for id := first; id <= last; id++ {
			if isValid(id, false) {
				p1 += id
			}
			if isValid(id, true) {
				p2 += id
			}
		}
	}
	return p1, p2
}

func isValid(id int, part2 bool) bool {
	_id := conv.ToStr(id)

	if part2 {
		double := strings.Repeat(_id, 2)
		return strings.Contains(double[1:len(double)-1], _id)
	}

	if len(_id)%2 != 0 {
		return false
	}
	mid := len(_id) / 2

	left := _id[:mid]
	right := _id[mid:]

	return left == right
}

func parseInput() [][2]int {
	return slice.Map(strings.Split(strings.TrimRight(input, "\n"), ","), func(item string, index int) [2]int {
		r := strings.Split(item, "-")
		return [2]int{conv.ToInt(r[0]), conv.ToInt(r[1])}
	})
}
