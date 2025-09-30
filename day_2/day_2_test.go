package main

import (
	"os"
	"strings"
	"testing"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils/parsers"
)

func setup() [][]int {
	wd, _ := os.Getwd()
	if strings.Contains(wd, "day_2") {
		os.Chdir("..")
	}
	input := utils.FetchTestInput(2)
	parser := parsers.RowIntParser{BaseParser: parsers.BaseParser{Input: input}}
	rows := parser.GetInts(" ")

	return rows
}

func TestPartOne(t *testing.T) {
	rows := setup()
	total := PartOne(rows)
	expected := 2

	if total != expected {
		t.Errorf("PartOne() = %d; want %d", total, expected)
	}
}

func TestPartTwo(t *testing.T) {
	rows := setup()
	total := PartTwo(rows)
	expected := 4

	if total != expected {
		t.Errorf("PartTwo() = %d; want %d", total, expected)
	}
}
