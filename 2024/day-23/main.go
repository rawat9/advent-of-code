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
	connections := parseInput()
	fmt.Println(connections)
	networkMap := make(map[string][]string)

	for i, connection := range connections {
		a, b := connection[0], connection[1]

		networkMap[a] = append(networkMap[a], b)

		for _, item := range connections[i+1:] {
			otherA, otherB := item[0], item[1]

			if otherA == a && !slices.Contains(networkMap[a], otherB) {
				networkMap[a] = append(networkMap[a], otherB)
			} else if otherB == a && !slices.Contains(networkMap[a], otherA) {
				networkMap[a] = append(networkMap[a], otherA)
			}
		}

		//for _, conn := range connections[i+1:] {
		//	otherA, otherB := conn[0], conn[1]
		//
		//	if otherA == a {
		//		networkMap[a] = append(networkMap[a], otherA)
		//	}
		//
		//	if otherB == a {
		//		networkMap[a] = append(networkMap[a], otherB)
		//	}
		//}
	}

	fmt.Println(networkMap)

	return 0
}

// match with kh: [ tc, qp, ub, ta]

// match with qp: [ kh, ub, td, wh ]

func part2() int {
	return 0
}

func parseInput() [][2]string {
	var connections [][2]string

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		computers := strings.Split(line, "-")
		connections = append(connections, [2]string{computers[0], computers[1]})
	}

	return connections
}
