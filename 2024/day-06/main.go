package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/bndr/gotabulate"
	"os"
	"os/exec"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1()
}

func part1() {
	matrix := parseInput()
	traverse(matrix)
}

func isBlocked(matrix [][]string, position [2]int) bool {
	x, y := position[0], position[1]
	return matrix[x][y] == "#"
}

var DIRECTIONS = map[string][2]int{"UP": {-1, 0}, "DOWN": {1, 0}, "RIGHT": {0, 1}, "LEFT": {0, -1}}

func move(matrix [][]string, direction string, pos [2]int, visited [][2]int) {
	dx := pos[0] + DIRECTIONS[direction][0]
	dy := pos[1] + DIRECTIONS[direction][1]
	nextPos := [2]int{dx, dy}

	if dx >= len(matrix) || dy >= len(matrix[0]) {
		fmt.Println(visited)
		fmt.Println(len(visited))
		//pprint(matrix)
		return
	}

	//	fmt.Println("LAST", dx, dy, pos)
	//	fmt.Println(len(visited))
	//	pprint(matrix)
	//	return
	//}

	if isBlocked(matrix, nextPos) {
		switch direction {
		case "UP":
			move(matrix, "RIGHT", pos, visited)
		case "DOWN":
			move(matrix, "LEFT", pos, visited)
		case "RIGHT":
			move(matrix, "DOWN", pos, visited)
		case "LEFT":
			move(matrix, "UP", pos, visited)
		}
	}
	if !slices.Contains(visited, nextPos) {
		visited = append(visited, nextPos)
		matrix[nextPos[0]][nextPos[1]] = "X"
	}
	move(matrix, direction, nextPos, visited)
}

func traverse(matrix [][]string) {
	startX, startY := findStartPosition(matrix)

	move(matrix, "UP", [2]int{startX, startY}, [][2]int{})
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

func part2() {}

func parseInput() [][]string {
	var matrix [][]string

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		slice := strings.Split(line, "")
		matrix = append(matrix, slice)
	}

	return matrix
}

func pprint(matrix [][]string) {
	var m []string
	for _, thing := range matrix {
		fmt.Println(strings.Join(thing, ""))
	}
	table := gotabulate.Create(m)
	fmt.Println(table.Render("simple"))
}

//func uniq[T int, Slice ~[2]T](slice Slice) Slice {
//	slices.Sort(slice)
//	return slices.Compact(slice)
//}

func clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}

func flatPrint(matrix [][]string) {
	var m []string

	for _, thing := range matrix {
		//m = append(m, strings.Join(thing, ""))
		s, _ := json.MarshalIndent(thing, "", " ")
		m = append(m, string(s))
	}
	fmt.Println(m)
}
