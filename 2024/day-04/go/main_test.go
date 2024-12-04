package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1Parts(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		// horizontal forward
		{"MMMXMASMMM\n", 1},
		// horizontal backwards
		{"MMMSAMXMMM\n", 1},
		// vertical forward
		{"MXMMMMMMMM\nMMMMMMMMMM\nMAMMMMMMMM\nMSMMMMMMMM\n", 1},
		// vertical backwards
		{"MSMMMMMMMM\nMAMMMMMMMM\nMMMMMMMMMM\nMXMMMMMMMM\n", 1},
		// diagonal left-to-right forward
		{"MXMMMMMMMM\nMMMMMMMMMM\nMMMAMMMMMM\nMMMMSMMMMM\n", 1},
		// diagonal left-to-right backwards
		{"MSMMMMMMMM\nMMAMMMMMMM\nMMMMMMMMMM\nMMMMXMMMMM\n", 1},
	}
	for _, test := range testCases {
		result := part1(test.in)
		assert.Equal(t, test.out, result, fmt.Sprintf("they should be equal for line %s", test.in))
	}
}

func TestPart1Full(t *testing.T) {
	expected := 18
	actual := part1(input)
	assert.Equal(t, expected, actual)
}

func TestPart2Full(t *testing.T) {
	expected := 9
	actual := part2(input)
	assert.Equal(t, expected, actual)
}
