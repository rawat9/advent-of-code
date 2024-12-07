package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/conv"
	S "github.com/rawat9/go-utils/slice"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	exprMap := parseInput()
	count := 0

	for k, v := range exprMap {
		var s []int
		calc(v, -1, &s)
		if slices.Contains(s, k) {
			count += k
		}
	}

	return count
}

func calc(l []int, start int, s *[]int) {
	if len(l) == 0 {
		if !slices.Contains(*s, start) {
			*s = append(*s, start)
		}
	} else {
		if start == -1 {
			calc(l[1:], l[0], s)
		} else {
			calc(l[1:], start+l[0], s)
			calc(l[1:], start*l[0], s)
		}
	}
}

func calc2(l []int, start int, p *[]int, num int) {
	if start != -1 && start > num {
		return
	}
	if len(l) == 0 {
		if !slices.Contains(*p, start) {
			*p = append(*p, start)
		}
	} else {
		if start == -1 {
			calc2(l[1:], l[0], p, num)
		} else {
			calc2(l[1:], start+l[0], p, num)
			calc2(l[1:], start*l[0], p, num)
			calc2(l[1:], concat(start, l[0]), p, num)
		}
	}
}

func concat(m int, n int) int {
	return conv.ToInt(conv.ToStr(m) + conv.ToStr(n))
}

func part2() int {
	exprMap := parseInput()
	count := 0

	for k, v := range exprMap {
		var p []int
		calc2(v, -1, &p, k)
		if slices.Contains(p, k) {
			count += k
		}
	}

	return count
}

func parseInput() map[int][]int {
	m := make(map[int][]int)

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		slice := strings.Split(line, ": ")
		nums := S.Map(strings.Split(slice[1], " "), func(item string, index int) int {
			return conv.ToInt(item)
		})
		m[conv.ToInt(slice[0])] = nums
	}

	return m
}
