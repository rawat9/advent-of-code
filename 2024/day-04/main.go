package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	wordSearch := parseInput()
	matrix := toMatrix(wordSearch)

	fmt.Println(part1(matrix))
	fmt.Println(part2(matrix))
}

// find XMAS
func part1(matrix [][]string) int {
	return horizontalCount(matrix) + verticalCount(matrix) + diagsCount(matrix)
}

func part2(matrix [][]string) int {
	m := len(matrix)
	n := len(matrix[0])
	count := 0

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			curr := matrix[i][j]

			if j < 1 || j > n-2 {
				return 0
			}

			// MM SS - SS MM - SM SM - MS MS

			if curr == "A" {

				// M.M
				// .A.
				// S.S
				if matrix[i-1][j-1] == "M" && matrix[i+1][j-1] == "S" && matrix[i-1][j+1] == "M" && matrix[i+1][j+1] == "S" {
					count++
				}

				// S.S
				// .A.
				// M.M
				if matrix[i-1][j-1] == "S" && matrix[i+1][j-1] == "M" && matrix[i-1][j+1] == "S" && matrix[i+1][j+1] == "M" {
					count++
				}

				// M.S
				// .A.
				// M.S
				if matrix[i-1][j-1] == "M" && matrix[i+1][j-1] == "M" && matrix[i-1][j+1] == "S" && matrix[i+1][j+1] == "S" {
					count++
				}

				// S.M
				// .A.
				// S.M
				if matrix[i-1][j-1] == "S" && matrix[i+1][j-1] == "S" && matrix[i-1][j+1] == "M" && matrix[i+1][j+1] == "M" {
					count++
				}
			}
		}
	}

	return count
}

func parseInput() []string {
	var wordSearch []string

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		wordSearch = append(wordSearch, line)
	}

	return wordSearch
}

func toMatrix(wordSearch []string) [][]string {
	var matrix [][]string

	for _, l := range wordSearch {
		slice := strings.Split(l, "")
		matrix = append(matrix, slice)
	}

	return matrix
}

func diagsCount(matrix [][]string) int {
	diags := 0

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curr := matrix[i][j]

			if curr == "X" {
				// upward
				if i >= 3 && j <= n-1-3 {
					if matrix[i-1][j+1] == "M" && matrix[i-2][j+2] == "A" && matrix[i-3][j+3] == "S" {
						diags++
					}
				}
				// downward
				if i <= n-1-3 && j <= n-1-3 {
					if matrix[i+1][j+1] == "M" && matrix[i+2][j+2] == "A" && matrix[i+3][j+3] == "S" {
						diags++
					}
				}

				// left diagonals
				// upward
				if i >= 3 && j >= 3 {
					if matrix[i-1][j-1] == "M" && matrix[i-2][j-2] == "A" && matrix[i-3][j-3] == "S" {
						diags++
					}
				}

				// downward
				if i <= n-1-3 && j >= 3 {
					if matrix[i+1][j-1] == "M" && matrix[i+2][j-2] == "A" && matrix[i+3][j-3] == "S" {
						diags++
					}
				}
			}

		}
	}
	return diags
}

func horizontalCount(matrix [][]string) int {
	m := len(matrix)
	n := len(matrix[0])

	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curr := matrix[i][j]

			if curr == "X" {
				// horizontal backward
				if j >= 3 {
					if matrix[i][j-1] == "M" && matrix[i][j-2] == "A" && matrix[i][j-3] == "S" {
						count++
					}
				}

				// horizontal forward
				if j <= n-1-3 {
					if matrix[i][j+1] == "M" && matrix[i][j+2] == "A" && matrix[i][j+3] == "S" {
						count++
					}
				}
			}
		}
	}

	return count
}

func verticalCount(matrix [][]string) int {
	m := len(matrix)
	n := len(matrix[0])

	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curr := matrix[i][j]

			if curr == "X" {
				// vertical downward
				if i <= n-1-3 {
					if matrix[i+1][j] == "M" && matrix[i+2][j] == "A" && matrix[i+3][j] == "S" {
						count++
					}
				}

				// vertical upward
				if i >= 3 {
					if matrix[i-1][j] == "M" && matrix[i-2][j] == "A" && matrix[i-3][j] == "S" {
						count++
					}
				}
			}
		}
	}

	return count
}
