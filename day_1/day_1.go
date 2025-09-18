package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils"
)

func main() {
	input := utils.FetchInput(2024, 1)
	parser := advent_of_gode_utils.ColumnIntParser{Input: input}
	columns := parser.GetInts("   ")
	sort.Ints(columns[0])
	sort.Ints(columns[1])

	fmt.Printf("Part One: %d\n", PartOne(columns))
	fmt.Printf("Part Two: %d\n", PartTwo(columns))
}

func PartOne(columns [][]int) int {
	total := 0

	for i := 0; i < len(columns[0]); i++ {
		total += int(math.Abs(float64(columns[0][i] - columns[1][i])))
	}

	return total
}

func PartTwo(columns [][]int) int {
	total := 0

	mem := make(map[int]int)
	cursors := [2]int{0, 0}

	for cursors[0] < len(columns[0]) {
		left := columns[0][cursors[0]]
		cache, ok := mem[left]
		if ok {
			total += cache
		} else {
			mem[left] = 0
		Loop:
			for cursors[1] < len(columns[1]) {
				right := columns[1][cursors[1]]
				switch {
				case right == left:
					mem[left] += left
					cursors[1]++
				case right < left:
					cursors[1]++
				default:
					break Loop
				}
			}
			total += mem[left]
		}
		cursors[0]++
	}

	return total
}
