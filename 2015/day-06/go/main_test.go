package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{"turn on 0,0 through 2,2", 9},
		{"turn on 10,10 through 12,12", 9},
		{"turn on 0,0 through 999,999", 1_000_000},
		{"toggle 0,0 through 999,0", 1_000},
		{"turn on 0,0 through 999,0\ntoggle 0,0 through 999,0", 0},
		{"turn off 499,499 through 500,500", -4},
		{"turn on 0,0 through 999,999\ntoggle 0,0 through 999,0\nturn off 499,499 through 500,500", 998_996},
	}
	for _, test := range input {
		result := part1(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestPart2(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{"turn on 0,0 through 0,0", 1},
		{"toggle 0,0 through 999,999", 2_000_000},
	}
	for _, test := range input {
		result := part2(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
