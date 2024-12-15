package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/conv"
	"strings"
)

//go:embed input.txt
var input string

type xy struct {
	x int
	y int
}

func main() {
	in := parseInput()
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

func part1(input map[xy][]xy) int {
	tokens := 0

	for k, v := range input {
		vx := v[0]
		vy := v[1]

		x, y := calc(vx, vy, k)
		tokens += x*3 + y
	}
	return tokens
}

// Calculate determinant
func calc(a, b, k xy) (int, int) {
	detAB := a.x*b.y - a.y*b.x
	detBK := k.x*b.y - k.y*b.x
	detAK := a.x*k.y - a.y*k.x

	if detBK%detAB != 0 || detAK%detAB != 0 {
		return 0, 0
	}

	// calculate x and y
	x := detBK / detAB
	y := detAK / detAB

	return x, y
}

func part2(input map[xy][]xy) int {
	// add 10000000000000 to x and y
	tokens := 0
	newIn := make(map[xy][]xy)

	for k := range input {
		newIn[xy{x: k.x + 10000000000000, y: k.y + 10000000000000}] = input[k]
	}

	for k, v := range newIn {
		vx := v[0]
		vy := v[1]

		x, y := calc(vx, vy, k)
		tokens += x*3 + y
	}
	return tokens
}

func parseInput() map[xy][]xy {
	conf := make(map[xy][]xy)

	for _, group := range strings.Split(strings.TrimRight(input, "\n"), "\n\n") {
		g := strings.Split(group, "\n")

		buttonA := strings.Split(strings.Split(g[:2][0], ": ")[1:][0], ", ")
		buttonB := strings.Split(strings.Split(g[:2][1], ": ")[1:][0], ", ")
		aX, bX := strings.Split(buttonA[0], "+")[1], strings.Split(buttonB[0], "+")[1]
		aY, bY := strings.Split(buttonA[1], "+")[1], strings.Split(buttonB[1], "+")[1]

		prize := strings.Split(g[2:][0], ": ")[1:]
		p := strings.Split(prize[0], ", ")
		x, y := strings.Split(p[0], "="), strings.Split(p[1], "=")

		conf[xy{x: conv.ToInt(x[1]), y: conv.ToInt(y[1])}] = []xy{{x: conv.ToInt(aX), y: conv.ToInt(aY)}, {x: conv.ToInt(bX), y: conv.ToInt(bY)}}
	}

	return conf
}
