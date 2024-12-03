package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/conv"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	pattern := `mul\(\d+,\d+\)`
	input = strings.TrimRight(input, "\n")
	re, err := regexp.Compile(pattern)

	if err != nil {
		panic(err)
	}

	c := regexp.MustCompile(`[0-9]+`)

	res := 0

	for _, v := range re.FindAllString(input, -1) {
		digits := c.FindAllString(v, -1)
		res += conv.ToInt(digits[0]) * conv.ToInt(digits[1])
	}

	return res
}

func part2() int {
	pattern := `don't\(\)|do\(\)|mul\(\d+,\d+\)`
	input = strings.TrimRight(input, "\n")
	re, err := regexp.Compile(pattern)

	if err != nil {
		panic(err)
	}

	c := regexp.MustCompile(`[0-9]+`)

	canDo := true

	res := 0
	for _, v := range re.FindAllString(input, -1) {
		if v == "don't()" {
			canDo = false
		}
		if v == "do()" {
			canDo = true
			continue
		}

		if canDo {
			digits := c.FindAllString(v, -1)
			res += conv.ToInt(digits[0]) * conv.ToInt(digits[1])
		}
	}

	return res
}
