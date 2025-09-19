package main

import (
	"os"
	"strings"
	"testing"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
)

func setup() string {
	wd, _ := os.Getwd()
	if strings.Contains(wd, "day_3") {
		os.Chdir("..")
	}
	input := utils.FetchTestInput(3)
	return input
}

func TestPartOne(t *testing.T) {
	input := setup()
	total := PartOne(input)
	expected := 161

	if total != expected {
		t.Errorf("PartOne() = %d; want %d", total, expected)
	}
}

func TestPartTwo(t *testing.T) {
	setup()
	input := utils.FetchSpecialTestInput(3)
	total := PartTwo(input)
	expected := 48

	if total != expected {
		t.Errorf("PartTwo() = %d; want %d", total, expected)
	}
}
