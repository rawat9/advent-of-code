package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/conv"
	"github.com/rawat9/go-utils/slice"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	blocks := buildBlocks()

	i := 0
	j := len(blocks) - 1

	for i <= j {
		if blocks[i] == -1 {
			if blocks[j] == -1 {
				for blocks[j] != -1 {
					j -= 1
				}
			}
			blocks[i], blocks[j] = blocks[j], blocks[i]
			j -= 1
		} else {
			i += 1
		}
	}

	r := slice.Filter(blocks, func(item int, index int) bool {
		return item != -1
	})
	res := 0

	for i, v := range r {
		res += i * v
	}

	return res
}

func buildBlocks() []int {
	var blocks []int
	count := 0

	for index, value := range strings.Split(strings.TrimRight(input, "\n"), "") {
		if index%2 != 0 {
			s := make([]int, conv.ToInt(value))
			for i := range s {
				s[i] = -1
			}
			blocks = append(blocks, s...)
		} else {
			s := make([]int, conv.ToInt(value))
			for i := range s {
				s[i] = count
			}
			blocks = append(blocks, s...)
			count++
		}
	}

	return blocks
}

func part2() int {
	return 0
}
