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
	stones := parseInput()

	cache := make(map[int]int)
	for _, v := range stones {
		cache[v]++
	}

	p1 := blink(cache, 0, 25)
	p2 := blink(cache, 0, 75)

	fmt.Println(p1, p2)
}

func blink(stones map[int]int, n int, limit int) int {
	if n == limit {
		return calc(stones)
	}

	cache := make(map[int]int)

	for stone, v := range stones {
		if isEvenDigits(stone) {
			n := len(conv.ToStr(stone))
			firstHalf, secondHalf := conv.ToStr(stone)[:n/2], conv.ToStr(stone)[n/2:]
			cache[conv.ToInt(firstHalf)] += v
			cache[conv.ToInt(secondHalf)] += v
		} else if stone == 0 {
			cache[1] += v
		} else {
			cache[stone*2024] += v
		}
	}

	return blink(cache, n+1, limit)
}

func isEvenDigits(num int) bool {
	return len(conv.ToStr(num))%2 == 0
}

func parseInput() []int {
	return slice.Map(strings.Split(strings.TrimRight(input, "\n"), " "), func(it string, index int) int {
		return conv.ToInt(it)
	})
}

func calc(stones map[int]int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	return sum
}
