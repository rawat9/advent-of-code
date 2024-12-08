package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part2()
}

func part1() {
	matrix := parseInput()
	move(matrix, false)
}

func part2() {
	matrix := parseInput()
	visited := move(matrix, false)
	fmt.Println(visited, len(visited))
	cycles := 0

	for _, v := range visited[2:] {
		x, y := v[0], v[1]
		matrix = reset(&matrix)
		matrix[x][y] = "O"
		if move2(matrix, true) {
			cycles++
		}
	}
	fmt.Println(cycles)
}

func reset(matrix *[][]string) [][]string {
	*matrix = parseInput()
	return *matrix
}

func isBlocked(matrix [][]string, x, y int, part2 bool) bool {
	if part2 {
		return matrix[x][y] == "#" || matrix[x][y] == "O"
	}
	return matrix[x][y] == "#"
}

func move2(matrix [][]string, part2 bool) bool {
	startX, startY := findStartPosition(matrix)
	blockedX, blockedY := findBlockedPosition(matrix)

	direction := "UP"
	matrix[startX][startY] = "X"
	visited := [][2]int{{startX, startY}}

	i, j := startX, startY
	prev := [2]int{i, j}
	//var blocked [][2]int
	blockedCount := 0

	for {
		fmt.Println(blockedCount)
		if blockedCount >= 15 {
			return true
		}
		if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
			return false
		}
		if isBlocked(matrix, i, j, part2) {
			switch direction {
			case "UP":
				direction = "RIGHT"
			case "RIGHT":
				direction = "DOWN"
			case "DOWN":
				direction = "LEFT"
			case "LEFT":
				direction = "UP"
			}
			if blockedX == i && blockedY == j {
				blockedCount++
			}
			i = prev[0]
			j = prev[1]
		} else {
			pos := [2]int{i, j}
			if !slices.Contains(visited, pos) {
				visited = append(visited, pos)
				matrix[i][j] = "X"
			}
		}

		prev = [2]int{i, j}
		i += DIRECTIONS[direction][0]
		j += DIRECTIONS[direction][1]
	}
}

func move(matrix [][]string, part2 bool) [][2]int {
	startX, startY := findStartPosition(matrix)
	blockedX, blockedY := findBlockedPosition(matrix)

	direction := "UP"
	matrix[startX][startY] = "X"
	visited := [][2]int{{startX, startY}}

	i, j := startX, startY
	prev := [2]int{i, j}
	//var blocked [][2]int
	blockedCount := 0

	for {
		if blockedCount > 10 {
		}
		if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
			return visited
		}
		if isBlocked(matrix, i, j, part2) {
			switch direction {
			case "UP":
				direction = "RIGHT"
			case "RIGHT":
				direction = "DOWN"
			case "DOWN":
				direction = "LEFT"
			case "LEFT":
				direction = "UP"
			}
			if blockedX == i && blockedY == j {
				blockedCount++
			}
			//
			//b := [2]int{blockedX, blockedY}
			//if slices.Contains(blocked, b) {
			//	fmt.Println("CYCLE", i, j)
			//	return [][2]int{{1, 2}}
			//}
			i = prev[0]
			j = prev[1]
		} else {
			pos := [2]int{i, j}
			if !slices.Contains(visited, pos) {
				visited = append(visited, pos)
				matrix[i][j] = "X"
			}
		}
		//pprint(matrix)

		prev = [2]int{i, j}
		i += DIRECTIONS[direction][0]
		j += DIRECTIONS[direction][1]
	}
}

func findStartPosition(matrix [][]string) (int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "^" {
				return i, j
			}
		}
	}

	return -1, -1
}

func findBlockedPosition(matrix [][]string) (int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "O" {
				return i, j
			}
		}
	}

	return -1, -1
}

var DIRECTIONS = map[string][2]int{"UP": {-1, 0}, "DOWN": {1, 0}, "RIGHT": {0, 1}, "LEFT": {0, -1}}

func parseInput() [][]string {
	var matrix [][]string

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		slice := strings.Split(line, "")
		matrix = append(matrix, slice)
	}

	return matrix
}

func pprint(matrix [][]string) {
	for _, thing := range matrix {
		fmt.Println(strings.Join(thing, ""))
	}
}
