package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
)

func main() {
	input := utils.FetchInput(2024, 3)

	fmt.Printf("Part One: %d\n", PartOne(input))
	fmt.Printf("Part Two: %d\n", PartTwo(input))
}

func PartOne(input string) int {
	total := 0

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	instructions := re.FindAllString(input, -1)

	for _, instruction := range instructions {
		factors := strings.Split(instruction[4:len(instruction)-1], ",")
		factorOne, err := strconv.Atoi(factors[0])
		if err != nil {
			continue
		}
		factorTwo, err := strconv.Atoi(factors[1])
		if err != nil {
			continue
		}
		total += factorOne * factorTwo
	}

	return total
}

func PartTwo(input string) int {
	total := 0

	re := regexp.MustCompile(`(mul\(\d+,\d+\)|(do|don't)\(\))`)
	instructions := re.FindAllString(input, -1)

	do := true
	for _, instruction := range instructions {
		if instruction == "do()" {
			do = true
			continue
		}
		if instruction == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}
		factors := strings.Split(instruction[4:len(instruction)-1], ",")
		factorOne, err := strconv.Atoi(factors[0])
		if err != nil {
			continue
		}
		factorTwo, err := strconv.Atoi(factors[1])
		if err != nil {
			continue
		}
		total += factorOne * factorTwo
	}

	return total
}
