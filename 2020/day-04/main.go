package main

import (
	_ "embed"
	"fmt"
	"github.com/rawat9/go-utils/conv"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1())
}

func part1() int {
	passports := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	valid := 0

	for _, fields := range passports {
		ps := make(map[string]string, 8)
		for _, field := range strings.FieldsFunc(fields, func(r rune) bool {
			return r == ' ' || r == '\n'
		}) {
			f := strings.Split(field, ":")
			ps[f[0]] = f[1]
		}

		if isValid(ps) {
			valid += 1
		}
	}
	return valid
}

func isValid(passport map[string]string) bool {
	_, cidOk := passport["cid"]

	byr := passport["byr"]
	if (len(byr) != 4) || (conv.ToInt(byr) < 1920 || conv.ToInt(byr) > 2002) {
		return false
	}

	iyr := passport["iyr"]
	if (len(iyr) != 4) || (conv.ToInt(iyr) < 2010 || conv.ToInt(iyr) > 2020) {
		return false
	}

	eyr := passport["eyr"]
	if (len(eyr) != 4) || (conv.ToInt(eyr) < 2020 || conv.ToInt(eyr) > 2030) {
		return false
	}

	hgt := passport["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		h := strings.Split(hgt, "cm")[0]
		if conv.ToInt(h) < 150 || conv.ToInt(h) > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		i := strings.Split(hgt, "in")[0]
		if conv.ToInt(i) < 59 || conv.ToInt(i) > 76 {
			return false
		}
	} else {
		return false
	}

	hcl := passport["hcl"]
	if strings.HasPrefix(hcl, "#") {
		color := hcl[1:]
		matched, _ := regexp.MatchString("[0-9a-f]", color)

		if len(color) != 6 || !matched {
			return false
		}
	} else {
		return false
	}

	ecl := passport["ecl"]
	eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	if !slices.Contains(eyeColors, ecl) {
		return false
	}

	pid := passport["pid"]
	if len(pid) != 9 {
		return false
	}

	return len(passport) == 8 || (len(passport) == 7 && !cidOk)
}
