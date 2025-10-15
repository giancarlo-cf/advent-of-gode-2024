package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils/data_structures"
	"github.com/giancarlo-cf/advent-of-gode-utils/parsers"
)

func main() {
	input := utils.FetchInput(2024, 5)

	fmt.Printf("Part One: %d\n", PartOne(input))
	fmt.Printf("Part Two: %d\n", PartTwo(input))
}

func process(input string) (rulesMap map[int]data_structures.Set[int], updates [][]int) {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	chunks := strings.Split(input, "\n\n")

	rulesParser := parsers.RowIntParser{BaseParser: parsers.BaseParser{Input: chunks[0]}}
	rules := rulesParser.GetInts("|")

	rulesMap = make(map[int]data_structures.Set[int])

	for _, rule := range rules {
		key := rule[0]
		value := rule[1]

		if _, ok := rulesMap[key]; !ok {
			rulesMap[key] = *data_structures.NewSet[int](make(map[int]struct{}))
		}

		set := rulesMap[key]
		set.Add(value)
	}

	updatesParser := parsers.RowIntParser{BaseParser: parsers.BaseParser{Input: chunks[1]}}
	updates = updatesParser.GetInts(",")

	return
}

func check(rulesMap map[int]data_structures.Set[int], slice []int) bool {
	target := slice[0]
	set := data_structures.NewSet(make(map[int]struct{}))
	for _, n := range slice[1:] {
		set.Add(n)
	}
	targetSet := rulesMap[target]

	return set.IsSubset(&targetSet)
}

func PartOne(input string) int {
	total := 0

	rulesMap, updates := process(input)

	for _, update := range updates {
		length := len(update)

		var wg sync.WaitGroup
		var failed atomic.Bool

		for i := 0; i < length-1; i++ {
			wg.Add(1)
			go func(slice []int) {
				defer wg.Done()
				if failed.Load() {
					return
				}
				if !check(rulesMap, slice) {
					failed.Store(true)
				}
			}(update[i:])
		}

		wg.Wait()
		if !failed.Load() {
			total += update[length/2]
		}
	}
	return total
}

func PartTwo(input string) int {
	total := 0

	rulesMap, updates := process(input)

	for _, update := range updates {
		length := len(update)

		for i := 0; i < length-1; i++ {
			offset := 1
			for !check(rulesMap, update[i:]) {
				update[i], update[i+offset] = update[i+offset], update[i]
				offset++
			}
		}

		total += update[length/2]
	}
	return total - PartOne(input)
}
