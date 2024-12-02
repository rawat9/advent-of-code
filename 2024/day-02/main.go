package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/conv"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part1() + part2())
}

func part1() int {
	levels := parseInput()

	safe := 0
	for _, level := range levels {
		if isSafe(level) {
			safe += 1
		}
	}

	return safe
}

func part2() int {
	levels := parseInput()

	var defUnsafeLevels [][]int
	for _, level := range levels {
		// definitely not the unsafe
		if !isSafe(level) {
			defUnsafeLevels = append(defUnsafeLevels, level)
		}
	}

	safe := 0
	for _, level := range defUnsafeLevels {
		// remove every element and check if it is safe
		for i := range level {
			a := slices.Concat(level[:i], level[i+1:])
			if isSafe(a) {
				safe += 1
				break
			}
		}
	}

	return safe
}

func parseInput() [][]int {
	var levels [][]int

	for _, report := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		l := strings.Split(report, " ")
		var level []int
		for _, v := range l {
			level = append(level, conv.ToInt(v))
		}
		levels = append(levels, level)
	}

	return levels
}

type pair struct {
	diff       int
	increasing bool
}

func isSafe(level []int) bool {
	var increasing bool
	var res []pair

	for i := 0; i < len(level)-1; i++ {
		prev, curr := level[i], level[i+1]
		diff := curr - prev

		if diff == 0 {
			return false
		}

		// decreasing level
		if diff < 0 {
			increasing = false
		}

		// increasing level
		if diff > 0 {
			increasing = true
		}

		res = append(res, pair{diff, increasing})
	}

	// check if the level is increasing or decreasing
	for i := 0; i < len(res)-1; i++ {
		if res[i].increasing != res[i+1].increasing {
			return false
		}
	}

	// check if the difference between the levels is within the range
	for _, r := range res {
		if math.Abs(float64(r.diff)) > 3 {
			return false
		}
	}

	return true
}
