package main

import (
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/rawat9/go-utils/conv"
	"github.com/rawat9/go-utils/slice"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	gateConnections, wires := parseInput()

	sort.Slice(wires, func(i, j int) bool {
		return strings.HasPrefix(wires[i].srcA, "x") || strings.HasPrefix(wires[i].srcB, "y")
	})

	for _, wire := range wires {
		calc(wire, wires, gateConnections)
	}

	for _, wire := range remaining {
		calc(wire, wires, gateConnections)
	}

	binaryZ := binaryNumberStartingWith("z", gateConnections)
	return convertToDecimal(Reverse(binaryZ))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func binaryNumberStartingWith(s string, gateConnections map[string]string) string {
	keys := maps.Keys(gateConnections)

	startsWith := slice.Filter(slices.Collect(keys), func(item string, index int) bool {
		return strings.HasPrefix(item, s)
	})

	sort.Slice(startsWith, func(i, j int) bool {
		return conv.ToInt(startsWith[i][1:]) < conv.ToInt(startsWith[j][1:])
	})

	binary := slice.Map(startsWith, func(item string, index int) string {
		return gateConnections[item]
	})
	return strings.Join(binary, "")
}

func convertToDecimal(binary string) int {
	res, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(res)
}

var remaining []wire

// 21436594354022
func calc(w wire, wires []wire, gateConnections map[string]string) {
	valA, valAOk := gateConnections[w.srcA]
	valB, valBOk := gateConnections[w.srcB]

	if valAOk && valBOk {
		switch w.gate {
		case "AND":
			if valA == "1" && valB == "1" {
				gateConnections[w.dest] = "1"
			} else {
				gateConnections[w.dest] = "0"
			}
		case "OR":
			if valA == "1" || valB == "1" {
				gateConnections[w.dest] = "1"
			} else {
				gateConnections[w.dest] = "0"
			}
		case "XOR":
			if valA != valB {
				gateConnections[w.dest] = "1"
			} else {
				gateConnections[w.dest] = "0"
			}
		}
	} else {
		for _, item := range wires {
			if item.dest == w.srcA {
				remaining = append(remaining, w)
				calc(item, wires, gateConnections)
			}
			if item.dest == w.srcB {
				remaining = append(remaining, w)
				calc(item, wires, gateConnections)
			}
		}
	}
}

func part2() int {
	return 0
}

type wire struct {
	srcA string
	srcB string
	dest string
	gate string
}

func parseInput() (map[string]string, []wire) {
	sections := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	gateConnections := make(map[string]string)

	for _, connection := range strings.Split(sections[0], "\n") {
		conn := strings.Split(connection, ": ")
		gateConnections[conn[0]] = conn[1]
	}

	var wires []wire

	initialValues := strings.Split(sections[1], "\n")

	for _, value := range initialValues {
		r := strings.Split(value, " -> ")
		l := strings.Split(r[0], " ")

		wires = append(wires, wire{dest: r[1], srcA: l[0], gate: l[1], srcB: l[2]})
	}

	return gateConnections, wires
}
