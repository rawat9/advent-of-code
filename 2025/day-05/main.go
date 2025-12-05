package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/rawat9/go-utils/conv"
	"github.com/rawat9/go-utils/slice"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {
	ranges, availableIDs := parseInput()
	fresh := 0

	for _, id := range availableIDs {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				fresh++
				break
			}
		}
	}

	return fresh
}

func part2() int {
	ranges, _ := parseInput()
	fresh := 0

	// sort ranges by start
	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.start, b.start)
	})

	r := []Range{ranges[0]}
	for _, curr := range ranges[1:] {
		last := r[len(r)-1]

		if last.end >= curr.start {
			last.end = max(last.end, curr.end)
			r[len(r)-1] = last
		} else {
			r = append(r, curr)
		}
	}

	for _, t := range r {
		fresh += (t.end - t.start) + 1
	}
	return fresh
}

type Range struct {
	start int
	end   int
}

func parseInput() ([]Range, []int) {
	in := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	IDRanges := slice.Map(strings.Split(strings.TrimRight(in[0], "\n"), "\n"), func(item string, index int) Range {
		r := strings.Split(item, "-")
		return Range{start: conv.ToInt(r[0]), end: conv.ToInt(r[1])}
	})
	availableIDs := slice.Map(strings.Split(strings.TrimRight(in[1], "\n"), "\n"), func(item string, index int) int {
		return conv.ToInt(item)
	})

	return IDRanges, availableIDs
}
