package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/array"
	"math"
	"slices"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1, part2 := solve()
	fmt.Println(part1, part2)
}

func solve() (int, int) {
	maxSeatID := -1

	var seats []int
	for _, v := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		seatId := calculateSeatID(v)
		maxSeatID = max(seatId, maxSeatID)
		seats = append(seats, seatId)
	}

	sort.Ints(seats)
	mySeat := getMySeatID(array.Uniq(seats))
	return maxSeatID, mySeat
}

func getMySeatID(seats []int) int {
	first, last := slices.Min(seats), slices.Max(seats)
	r := pyrange(first, last, 1)

	var mySeatId int

	for _, value := range r {
		if !slices.Contains(seats, value) {
			mySeatId = value
		}
	}

	return mySeatId
}

func pyrange(start, end, step int) []int {
	rtn := make([]int, 0, (end-start)/step)
	for i := start; i < end; i += step {
		rtn = append(rtn, i)
	}
	return rtn
}

func calculateSeatID(boardingPass string) int {
	rows := [2]int{0, 127}
	cols := [2]int{0, 7}

	for _, char := range boardingPass {
		rowDiff := rows[1] - rows[0]
		colDiff := cols[1] - cols[0]

		if string(char) == "F" {
			rows[1] = int(math.Floor(float64(rowDiff)/2)) + rows[0]

			// lower_half := 0
			//   [0, 127]
			// F [0,  63]
			// B [32, 63] ==> 63 / 2 = 31.5 + round up
			// F [32, 47] ==> 63 - 32 = 31 / 2 = 15.5 + round down
			// B [40, 47] ==> 47 - 32 = 15 / 2 = 7.5 + round up
			// B [44, 47] ==> 47 - 40 = 7 / 2 = 3.5 + round up
			// F [44, 45] ==> 47 - 44 = 3 / 2 = 1.5 + round down
		}
		if string(char) == "B" {
			rows[0] += int(math.Ceil(float64(rowDiff) / 2))
		}

		if string(char) == "L" {
			cols[1] = int(math.Floor(float64(colDiff)/2)) + cols[0]
		}

		if string(char) == "R" {
			cols[0] += int(math.Ceil(float64(colDiff) / 2))
		}
	}

	if rows[0] != rows[1] {
		fmt.Println("Unequal values", rows)
	}

	if cols[0] != cols[1] {
		fmt.Println("Unequal values", cols)
	}

	seatId := rows[0]*8 + cols[0]
	return seatId
}
