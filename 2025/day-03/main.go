package main

import (
	_ "embed"
	"fmt"
	"math"
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
	banks := parseInput()
	total := 0

	for _, bank := range banks {
		var res int

		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				curr := conv.ToInt(bank[i] + bank[j])
				res = int(math.Max(float64(res), float64(curr)))
			}
		}

		total += res
	}

	return total
}

func part2() int {
	banks := parseInput()
	var total int

	for _, bank := range banks {
		m := len(bank)
		i, j := 0, 1

		for len(bank) > 12 && j < len(bank) {
			if bank[i] < bank[j] {
				bank = slices.Delete(bank, i, j)
				i, j = 0, 1
			} else {
				i++
				j++
			}
		}
		if len(bank) <= m {
			bank = bank[:12]
		}

		total += conv.ToInt(strings.Join(bank, ""))
	}
	return total
}

func parseInput() [][]string {
	return slice.Map(strings.Split(strings.TrimRight(input, "\n"), "\n"), func(s string, index int) []string {
		return strings.Split(s, "")
	})
}
