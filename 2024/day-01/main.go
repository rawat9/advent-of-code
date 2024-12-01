package main

import (
	_ "embed"
	"github.com/rawat9/go-utils/conv"
	"math"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	println(part1())
	println(part2())
}

func part1() int {
	l1, l2 := parseInput()

	sort.Ints(l1)
	sort.Ints(l2)

	distance := 0.0

	for i := 0; i < len(l1); i++ {
		distance += math.Abs(float64(l1[i] - l2[i]))
	}
	return int(distance)
}

func part2() int {
	l1, l2 := parseInput()

	freqMap := make(map[int]int)

	// build frequency map
	for _, v := range l2 {
		freqMap[v]++
	}

	total := 0
	for _, v := range l1 {
		total += v * freqMap[v]
	}
	return total
}

func parseInput() ([]int, []int) {
	var l1, l2 []int
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		s := strings.Split(line, "   ")

		l1 = append(l1, conv.ToInt(s[0]))
		l2 = append(l2, conv.ToInt(s[1]))
	}
	return l1, l2
}
