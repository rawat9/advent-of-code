package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	grid := parseInput()
	fmt.Println("Part 1:", len(getPaperRolls(grid)))
	fmt.Println("Part 2:", part2(grid, 0))
}

var directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}

func part2(grid [][]string, acc int) int {
	paperRolls := getPaperRolls(grid)

	if len(paperRolls) == 0 {
		return acc
	}

	// remove paper rolls
	for _, roll := range paperRolls {
		grid[roll[0]][roll[1]] = "."
	}

	acc += len(paperRolls)
	return part2(grid, acc)
}

func getPaperRolls(grid [][]string) [][2]int {
	m := len(grid)
	n := len(grid[0])

	var paperRolls [][2]int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == "@" {
				var count int

				for _, dir := range directions {
					x, y := i+dir[0], j+dir[1]
					if x >= 0 && x < m && y >= 0 && y < n {
						if grid[x][y] == "@" {
							count++
						}
					}
				}
				if count < 4 {
					paperRolls = append(paperRolls, [2]int{i, j})
				}
			}
		}
	}
	return paperRolls
}

func parseInput() [][]string {
	var matrix [][]string

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}
	return matrix
}
