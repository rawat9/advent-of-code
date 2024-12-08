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
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	grid := parseInput()
	antennaMap := buildAntennaMap(grid)
	var visited [][2]int

	for _, v := range antennaMap {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				m, n := v[j][0], v[j][1]
				curr := [2]int{v[i][0], v[i][1]}

				dx, dy := m-curr[0], n-curr[1]
				px1, py1 := m+dx, n+dy
				px2, py2 := curr[0]-dx, curr[1]-dy
				p1, p2 := [2]int{px1, py1}, [2]int{px2, py2}
				if px1 <= len(grid)-1 && px1 >= 0 && py1 <= len(grid[0])-1 && py1 >= 0 && !slices.Contains(visited, p1) {
					visited = append(visited, p1)
					grid[px1][py1] = "#"
				}
				if px2 <= len(grid)-1 && px2 >= 0 && py2 <= len(grid[0])-1 && py2 >= 0 && !slices.Contains(visited, p2) {
					visited = append(visited, p2)
					grid[px2][py2] = "#"
				}
			}
		}
	}

	return len(visited)
}

func part2() int {
	grid := parseInput()
	antennaMap := buildAntennaMap(grid)
	var visited [][2]int

	for _, v := range antennaMap {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				m, n := v[j][0], v[j][1]
				currX, currY := v[i][0], v[i][1]

				dx, dy := m-currX, n-currY
				for m <= len(grid)-1 && n <= len(grid[0])-1 && m >= 0 && n >= 0 {
					pos := [2]int{m, n}
					if !slices.Contains(visited, pos) {
						visited = append(visited, pos)
					}
					m += dx
					n += dy
				}

				for currX <= len(grid)-1 && currX >= 0 && currY <= len(grid[0])-1 && currY >= 0 {
					pos := [2]int{currX, currY}
					if !slices.Contains(visited, pos) {
						visited = append(visited, pos)
					}
					currX -= dx
					currY -= dy
				}
			}
		}
	}

	return len(visited)
}

func buildAntennaMap(matrix [][]string) map[string][][2]int {
	antennaMap := make(map[string][][2]int)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != "." {
				antennaMap[matrix[i][j]] = append(antennaMap[matrix[i][j]], [2]int{i, j})
			}
		}
	}

	return antennaMap
}

func parseInput() [][]string {
	var matrix [][]string

	for _, row := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		r := strings.Split(row, "")
		matrix = append(matrix, r)
	}

	return matrix
}

func pprint(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}
