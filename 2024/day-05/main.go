package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/conv"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	//fmt.Println(part1())
	fmt.Println(part2())
}

// find correct updates
func part1() int {
	rules, updates := parseInput()

	rulesMap := buildRulesMap(rules)
	count := 0
	sum := 0

	for _, update := range updates {
		u := strings.Split(update, ",")
		result := processUpdate(rulesMap, u)
		if result {
			sum += conv.ToInt(u[len(u)/2])
			count++
		}
	}
	return sum
}

func processUpdate(rulesMap map[string][]string, update []string) bool {
	for i := 0; i < len(update)-1; i++ {
		val := rulesMap[update[i]]
		nextInOrder := update[i+1:]
		for _, v := range nextInOrder {
			ok := slices.Contains(val, v)

			if !ok {
				fmt.Println(update[i], v, ok)
				return false
			}
		}
	}
	return true
}

func buildRulesMap(rules []string) map[string][]string {
	rulesMap := make(map[string][]string)

	for _, rule := range rules {
		r := strings.Split(rule, "|")
		rulesMap[r[0]] = append(rulesMap[r[0]], r[1])
	}

	return rulesMap
}

func part2() int {
	rules, updates := parseInput()

	rulesMap := buildRulesMap(rules)
	count := 0
	sum := 0

	for _, update := range updates {
		u := strings.Split(update, ",")
		result := processUpdate2(rulesMap, u)
		if !result {
			sum += conv.ToInt(u[len(u)/2])
			count++
		}
		fmt.Println("===")
	}
	return sum
}

func processUpdate2(rulesMap map[string][]string, update []string) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {

		}
		val := rulesMap[update[i]]
		nextInOrder := update[i+1:]
		for _, v := range nextInOrder {
			ok := slices.Contains(val, v)

			if !ok {
				fmt.Println(update[i], v, ok)
				return false
			}
		}
	}
	return true
	// 97,13,75,29,47
	// 97,75,13,29,47
	// 97,75,29,47,13
}

func parseInput() ([]string, []string) {
	var pageOrderingRules []string
	var updates []string

	line := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	pageOrderingRules = strings.Split(line[0], "\n")
	updates = strings.Split(line[1], "\n")

	return pageOrderingRules, updates
}
