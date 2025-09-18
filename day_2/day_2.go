package main

import (
	"fmt"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils"
)

func main() {
	input := utils.FetchInput(2024, 2)
	parser := advent_of_gode_utils.RowIntParser{BaseParser: advent_of_gode_utils.BaseParser{Input: input}}
	rows := parser.GetInts(" ")

	fmt.Printf("Part One: %d\n", PartOne(rows))
	fmt.Printf("Part Two: %d\n", PartTwo(rows))
}

func countSafeReportsWithCandidates(rows [][]int) (int, [][][]int) {
	total := 0
	var candidates [][][]int

Loop:
	for _, row := range rows {
		diff := row[0] - row[1]
		factor := 1
		switch {
		case diff < 0:
			factor = -1
		case diff > 0:
			factor = 1
		}
		for i := 0; i < len(row)-1; i++ {
			diff = (row[i] - row[i+1]) * factor
			if diff < 1 || diff > 3 {
				var subCandidates [][]int

				if i > 0 {
					var firstCandidate []int
					firstCandidate = append(firstCandidate, row[:i-1]...)
					firstCandidate = append(firstCandidate, row[i:]...)
					subCandidates = append(subCandidates, firstCandidate)
				}

				var secondCandidate []int
				secondCandidate = append(secondCandidate, row[:i]...)
				secondCandidate = append(secondCandidate, row[i+1:]...)
				subCandidates = append(subCandidates, secondCandidate)

				var thirdCandidate []int
				thirdCandidate = append(thirdCandidate, row[:i+1]...)
				if i+2 < len(row) {
					thirdCandidate = append(thirdCandidate, row[i+2:]...)
				}
				subCandidates = append(subCandidates, thirdCandidate)

				candidates = append(candidates, subCandidates)
				continue Loop
			}
		}
		total++
	}

	return total, candidates
}

func PartOne(rows [][]int) int {
	total, _ := countSafeReportsWithCandidates(rows)
	return total
}

func PartTwo(rows [][]int) int {
	total, candidates := countSafeReportsWithCandidates(rows)

	for _, subCandidates := range candidates {
		if subtotal, _ := countSafeReportsWithCandidates(subCandidates); subtotal > 0 {
			total++
		}
	}

	return total
}
