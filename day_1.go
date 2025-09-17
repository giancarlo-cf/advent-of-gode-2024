package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
	advent_of_gode_utils "github.com/giancarlo-cf/advent-of-gode-utils"
)

func main() {
	input := utils.FetchInput(2024, 1)
	parser := advent_of_gode_utils.ColumnIntParser{input}
	columns := parser.GetInts("   ")

	columnOne := columns[0]
	columnTwo := columns[1]

	sort.Ints(columnOne)
	sort.Ints(columnTwo)

	total := 0

	for i := 0; i < len(columnOne); i++ {
		total += int(math.Abs(float64(columnOne[i] - columnTwo[i])))
	}

	fmt.Printf("Part One: %d\n", total)
}
