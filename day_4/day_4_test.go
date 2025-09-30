package main

import (
	"os"
	"strings"
	"testing"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils/parsers"
)

func setup() [][]rune {
	wd, _ := os.Getwd()
	if strings.Contains(wd, "day_4") {
		os.Chdir("..")
	}
	input := utils.FetchTestInput(4)
	parser := parsers.MatrixRuneParser{BaseParser: parsers.BaseParser{Input: input}}
	ws := parser.GetStrings()

	return ws
}

func TestPartOne(t *testing.T) {
	ws := setup()
	total := PartOne(ws)
	expected := 18

	if total != expected {
		t.Errorf("PartOne() = %d; want %d", total, expected)
	}
}

func TestPartTwo(t *testing.T) {
	ws := setup()
	total := PartTwo(ws)
	expected := 9

	if total != expected {
		t.Errorf("PartTwo() = %d; want %d", total, expected)
	}
}
