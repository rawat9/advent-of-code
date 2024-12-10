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
	//fmt.Println(part1())
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
	blocks := buildBlocks()
	fmt.Println("Total length", len(blocks), blocks[12:14])

	i, j := 0, 1
	m, n := len(blocks)-1, len(blocks)-1

	fmt.Println(blocks)

	for {
		if blocks[i] == -1 {
			j = i + 1
			for blocks[j] == -1 {
				j++
			}
			//fmt.Println(i, j)
			i = j
			j = i + 1
		}
		i++

		if blocks[m] == -1 || blocks[n] == -1 {
			//fmt.Println(m, n, blocks[m], blocks[n])
			m--
			n--
		}
		if blocks[m] == blocks[n] {
			for blocks[m] == blocks[n] {
				n--
			}
			fmt.Println(m, n)
			//m = n
			//n = m - 1
		}
		m--

	}

	return 0
}
