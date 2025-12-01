package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/rawat9/go-utils/conv"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {
	start := 50
	curr := start
	var password int

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		if strings.HasPrefix(line, "L") {
			distance := conv.ToInt(strings.TrimPrefix(line, "L")) % 100
			curr -= distance
			if curr < 0 {
				curr += 100
			}
		} else if strings.HasPrefix(line, "R") {
			distance := conv.ToInt(strings.TrimPrefix(line, "R")) % 100
			curr = curr + distance
			if curr >= 100 {
				curr -= 100
			}
		}

		if curr == 0 {
			password++
		}
	}

	return password
}

func part2() int {
	start := 50
	curr := start
	var password int

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		if strings.HasPrefix(line, "L") {
			distance := conv.ToInt(strings.TrimPrefix(line, "L")) % 100
			password += conv.ToInt(strings.TrimPrefix(line, "L")) / 100
			curr -= distance
			if curr < 0 {
				if curr+distance != 0 {
					password++
				}
				curr += 100
			}
		} else if strings.HasPrefix(line, "R") {
			distance := conv.ToInt(strings.TrimPrefix(line, "R")) % 100
			password += conv.ToInt(strings.TrimPrefix(line, "R")) / 100
			curr += distance
			if curr >= 100 {
				if curr != 100 {
					password++
				}
				curr -= 100
			}
		}

		if curr == 0 {
			password++
		}
	}

	return password
}
