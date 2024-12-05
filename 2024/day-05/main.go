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
	rules, updates := parseInput()

	rulesMap := buildRulesMap(rules)

	var (
		correct, incorrect int
	)

	for _, update := range updates {
		u := strings.Split(update, ",")
		result := processUpdate(rulesMap, u)
		if result {
			correct += conv.ToInt(u[len(u)/2])
		} else {
			correctOrdering(rulesMap, u)
			incorrect += conv.ToInt(u[len(u)/2])
		}
	}
	fmt.Println(correct, incorrect)
}

func processUpdate(rulesMap map[string][]string, update []string) bool {
	for i := 0; i < len(update)-1; i++ {
		val := rulesMap[update[i]]
		nextInOrder := update[i+1:]
		for _, v := range nextInOrder {
			ok := slices.Contains(val, v)

			if !ok {
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

func correctOrdering(rulesMap map[string][]string, update []string) {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			val := rulesMap[update[i]]
			ok := slices.Contains(val, update[j])

			if !ok {
				// swap
				update[i], update[j] = update[j], update[i]
			}
		}
	}
}

func parseInput() ([]string, []string) {
	var pageOrderingRules []string
	var updates []string

	line := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	pageOrderingRules = strings.Split(line[0], "\n")
	updates = strings.Split(line[1], "\n")

	return pageOrderingRules, updates
}
