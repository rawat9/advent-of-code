package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/slice"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
}

func part1() int {
	locks, keys := parseInput()

	count := 0
	for i := range locks {
		for j := range keys {
			if fit(locks[i], keys[j]) {
				count++
			}
		}
	}
	return count
}

func fit(locks, keys []int) bool {
	for i := range locks {
		if locks[i]+keys[i] > 5 {
			return false
		}
	}
	return true
}

func parseInput() ([][]int, [][]int) {
	locksAndKeys := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	var locks, keys [][]int

	for _, lk := range locksAndKeys {
		rows := strings.Split(lk, "\n")
		top := rows[0] // lock
		//bottom := rows[len(rows)-1] // key

		if isLock(top) {
			matrix := transpose(slice.Map(rows, func(row string, index int) []string {
				return strings.Split(row, "")
			}))
			locks = append(locks, calculateHeight(matrix))
		} else {
			matrix := transpose(slice.Map(rows, func(row string, index int) []string {
				return strings.Split(row, "")
			}))
			keys = append(keys, calculateHeight(matrix))
		}
	}

	return locks, keys
}

func isLock(row string) bool {
	return every(strings.Split(row, ""), func(value string) bool { return value == "#" })
}

func calculateHeight(matrix [][]string) []int {
	return slice.Map(matrix, func(row []string, index int) int {
		c := 0
		for i := range row {
			if row[i] == "#" {
				c++
			}
		}
		return c - 1
	})
}

func every[T any](arr []T, predicate func(value T) bool) bool {
	result := true
	for i := range arr {
		result = result && predicate(arr[i])
	}
	return result
}

func transpose[T any](matrix [][]T) [][]T {
	m := len(matrix)
	n := len(matrix[0])

	newMatrix := make([][]T, n)

	for i := range newMatrix {
		newMatrix[i] = make([]T, m)

	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}

	return newMatrix
}
