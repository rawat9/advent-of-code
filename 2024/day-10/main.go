package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/conv"
	"github.com/rawat9/go-utils/slice"
	"log"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

var directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var serialize = func(x int, y int) string {
	return conv.ToStr(x) + "-" + conv.ToStr(y)
}

func part1() int {
	matrix := parseInput()
	starts := findStartPositions(matrix)

	score := 0
	for _, start := range starts {
		_, b := traversePath(matrix, start)
		score += b
	}

	return score
}

// BFS
func traversePath(matrix [][]int, start [2]int) (int, int) {
	vis := buildVis(len(matrix), len(matrix[0]))
	path := make(map[string][][2]int)

	if start[0] == -1 || start[1] == -1 {
		log.Fatal("Could not find start location at", start)
	}

	var visited [][2]int

	queue := [][2]int{{start[0], start[1]}}
	distinctPath := make(map[string]int)
	score := 0

	for len(queue) != 0 {
		value := queue[0]
		queue = queue[1:]

		vis[value[0]][value[1]] = "#"

		for _, dir := range directions {
			dx, dy := dir[0], dir[1]

			currentHeight := matrix[value[0]][value[1]]

			x := value[0] + dx
			y := value[1] + dy

			if !(y <= len(matrix[0])-1 && y >= 0 && x <= len(matrix)-1 && x >= 0) {
				continue
			}
			newHeight := matrix[x][y]

			if newHeight != currentHeight+1 {
				continue
			}

			if newHeight == 9 {
				distinctPath[serialize(x, y)]++
				path[serialize(x, y)] = visited
			} else {
				if !slices.Contains(queue, [2]int{x, y}) {
					queue = append(queue, [2]int{x, y})
					visited = append(visited, [2]int{x, y})
				}
			}
		}
	}

	fmt.Println(distinctPath)
	for _, v := range distinctPath {
		score += v
	}

	return len(path), score
}

func parseInput() [][]int {
	var matrix [][]int

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		row := slice.Map(strings.Split(line, ""), func(item string, index int) int {
			return conv.ToInt(item)
		})
		matrix = append(matrix, row)
	}

	return matrix
}

func findStartPositions(matrix [][]int) [][2]int {
	m := len(matrix)
	n := len(matrix[0])

	var positions [][2]int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				positions = append(positions, [2]int{i, j})
			}
		}
	}

	return positions
}

func part2() int {
	return 0
}

func buildVis(m int, n int) [][]string {
	matrix := make([][]string, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]string, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = "0"
		}
	}
	return matrix
}
