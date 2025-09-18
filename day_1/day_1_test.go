package main

import (
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils"
)

func setup() [][]int {
	wd, _ := os.Getwd()
	if strings.Contains(wd, "day_1") {
		os.Chdir("..")
	}
	input := utils.FetchTestInput(1)
	parser := advent_of_gode_utils.ColumnIntParser{BaseParser: advent_of_gode_utils.BaseParser{Input: input}}
	columns := parser.GetInts("   ")
	sort.Ints(columns[0])
	sort.Ints(columns[1])

	return columns
}

func TestPartOne(t *testing.T) {
	columns := setup()
	total := PartOne(columns)
	expected := 11

	if total != expected {
		t.Errorf("PartOne() = %d; want %d", total, expected)
	}
}

func TestPartTwo(t *testing.T) {
	columns := setup()
	total := PartTwo(columns)
	expected := 31

	if total != expected {
		t.Errorf("PartTwo() = %d; want %d", total, expected)
	}
}
