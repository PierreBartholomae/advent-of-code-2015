package main

import (
	"os"
	"strings"
)

type Pos struct {
	r, c int
}

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		panic("File at path ./input could not be read.")
	}
	input := string(bytes)

	resultPart1 := part1(input)
	println("answer part 1:", resultPart1)

	resultPart2 := part2(input)
	println("answer part 2:", resultPart2)
}

func part1(input string) int {
	result := 0
	moves := []Pos{
		// horizontal
		{0, 1}, {0, -1},
		// vertical
		{1, 0}, {-1, 0},
		// diagonal
		{-1, -1}, {-1, 1}, {1, 1}, {1, -1},
	}
	lines := strings.Split(input, "\n")
	rows := len(lines)
	columns := len(lines[0])
	// fmt.Printf("grid rows %d columns %d\n", rows, columns)
	for r := range rows {
		if lines[r] == "" {
			continue
		}
		for c := range columns {
			if lines[r][c] == 'X' {
				for _, move := range moves {
					if r+move.r*3 >= 0 && r+move.r*3 < rows && c+move.c*3 >= 0 && c+move.c*3 < columns {
						// fmt.Printf("bp: r %d, c %d, move.r %d, move.c %d\n", r, c, move.r, move.c)
						if lines[r+move.r][c+move.c] == 'M' && lines[r+move.r*2][c+move.c*2] == 'A' && lines[r+move.r*3][c+move.c*3] == 'S' {
							// fmt.Printf("found: r %d, c %d, move.r %d, move.c %d\n", r, c, move.r, move.c)
							result++
						}
					}
				}
			}
		}
	}
	return result
}

func part2(input string) int {
	result := 0
	lines := strings.Split(input, "\n")
	rows := len(lines)
	columns := len(lines[0])
	// fmt.Printf("grid rows %d columns %d\n", rows, columns)
	for r := range rows {
		if lines[r] == "" {
			continue
		}
		for c := range columns {
			if lines[r][c] == 'A' {
				if r-1 >= 0 && r+1 < rows && c-1 >= 0 && c+1 < columns {
					// fmt.Printf("bp: r %d, c %d, move.r %d, move.c %d\n", r, c, move.r, move.c)
					if (('M' == lines[r-1][c-1] && 'S' == lines[r+1][c+1]) || ('S' == lines[r-1][c-1] && 'M' == lines[r+1][c+1])) &&
						(('M' == lines[r-1][c+1] && 'S' == lines[r+1][c-1]) || ('S' == lines[r-1][c+1] && 'M' == lines[r+1][c-1])) {
						// fmt.Printf("found: r %d, c %d\n", r+1, c+1)
						result++
					}
				}
			}
		}
	}
	return result
}
