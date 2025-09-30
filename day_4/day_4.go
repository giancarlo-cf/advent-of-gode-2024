package main

import (
	"fmt"
	"regexp"

	"github.com/giancarlo-cf/advent-of-gode-2024/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils/parsers"
)

func main() {
	input := utils.FetchInput(2024, 4)
	parser := parsers.MatrixRuneParser{BaseParser: parsers.BaseParser{Input: input}}

	ws := parser.GetStrings()

	fmt.Printf("Part One: %d\n", PartOne(ws))
	fmt.Printf("Part Two: %d\n", PartTwo(ws))
}

func PartOne(ws [][]rune) int {
	total := 0
	n := len(ws)
	m := len(ws[0])

	strings := make([]string, 0)

	for i := 0; i < n; i++ {
		row := string(ws[i])
		strings = append(strings, row)

		diagonal := make([]rune, 0)

		for id, jd := i, 0; id < n && jd < m; id, jd = id+1, jd+1 {
			diagonal = append(diagonal, ws[id][jd])
		}
		strings = append(strings, string(diagonal))

		diagonal = make([]rune, 0)

		for id, jd := i, m-1; id < n && jd >= 0; id, jd = id+1, jd-1 {
			diagonal = append(diagonal, ws[id][jd])
		}
		strings = append(strings, string(diagonal))

		if i == 0 {
			continue
		}

		diagonal = make([]rune, 0)

		for id, jd := n-i-1, 0; id >= 0 && jd < m; id, jd = id-1, jd+1 {
			diagonal = append(diagonal, ws[id][jd])
		}

		strings = append(strings, string(diagonal))

		diagonal = make([]rune, 0)

		for id, jd := n-i-1, m-1; id >= 0 && jd >= 0; id, jd = id-1, jd-1 {
			diagonal = append(diagonal, ws[id][jd])
		}
		strings = append(strings, string(diagonal))

	}

	for j := 0; j < m; j++ {
		column := make([]rune, n)

		for i := 0; i < n; i++ {
			column[i] = ws[i][j]
		}
		strings = append(strings, string(column))
	}

	re := regexp.MustCompile(`XMAS`)
	reversedRe := regexp.MustCompile(`SAMX`)

	for _, s := range strings {
		total += len(re.FindAllString(s, -1))
		total += len(reversedRe.FindAllString(s, -1))
	}

	return total
}

func PartTwo(ws [][]rune) int {
	total := 0
	n := len(ws)
	m := len(ws[0])

	diagonals := make(map[string][][2]int)

	for i := 0; i < n; i++ {
		diagonal := make([]rune, 0)
		diagonalCoordinates := make([][2]int, 0)

		for id, jd := i, 0; id < n && jd < m; id, jd = id+1, jd+1 {
			diagonal = append(diagonal, ws[id][jd])
			diagonalCoordinates = append(diagonalCoordinates, [2]int{id, jd})
		}
		diagonals[string(diagonal)] = diagonalCoordinates

		diagonal = make([]rune, 0)
		diagonalCoordinates = make([][2]int, 0)

		for id, jd := i, m-1; id < n && jd >= 0; id, jd = id+1, jd-1 {
			diagonal = append(diagonal, ws[id][jd])
			diagonalCoordinates = append(diagonalCoordinates, [2]int{id, jd})
		}
		diagonals[string(diagonal)] = diagonalCoordinates

		if i == 0 {
			continue
		}

		diagonal = make([]rune, 0)
		diagonalCoordinates = make([][2]int, 0)

		for id, jd := n-i-1, 0; id >= 0 && jd < m; id, jd = id-1, jd+1 {
			diagonal = append(diagonal, ws[id][jd])
			diagonalCoordinates = append(diagonalCoordinates, [2]int{id, jd})
		}

		diagonals[string(diagonal)] = diagonalCoordinates

		diagonal = make([]rune, 0)
		diagonalCoordinates = make([][2]int, 0)

		for id, jd := n-i-1, m-1; id >= 0 && jd >= 0; id, jd = id-1, jd-1 {
			diagonal = append(diagonal, ws[id][jd])
			diagonalCoordinates = append(diagonalCoordinates, [2]int{id, jd})
		}
		diagonals[string(diagonal)] = diagonalCoordinates
	}

	re := regexp.MustCompile(`MAS`)
	reversedRe := regexp.MustCompile(`SAM`)

	appearances := make(map[[2]int]int)

	for diagonal, coordinates := range diagonals {
		indexes := append(re.FindAllStringIndex(diagonal, -1), reversedRe.FindAllStringIndex(diagonal, -1)...)
		for _, index := range indexes {
			coordinateOfA := coordinates[index[0]+1]
			_, ok := appearances[coordinateOfA]
			if ok {
				appearances[coordinateOfA]++
			} else {
				appearances[coordinateOfA] = 1
			}
		}
	}

	for _, count := range appearances {
		if count > 1 {
			total++
		}
	}

	return total
}
