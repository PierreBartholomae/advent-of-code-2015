package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

const input = "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"

func TestPart1Parts(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		{"7 6 4 2 1", 1},
		{"1 2 7 8 9", 0},
		{"9 7 6 2 1", 0},
		{"1 3 2 4 5", 0},
		{"8 6 4 4 1", 0},
		{"1 3 6 7 9", 1},
		{"8 11 13 14 15 18 17", 0},
	}
	for _, test := range testCases {
		result := part1(test.in)
		assert.Equal(t, test.out, result, fmt.Sprintf("they should be equal for line %s", test.in))
	}
}

func TestPart1Full(t *testing.T) {
	expected := 2
	actual := part1(input)
	assert.Equal(t, expected, actual)
}

func TestPart2(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		{"7 6 4 2 1", 1},
		{"1 2 7 8 9", 0},
		{"9 7 6 2 1", 0},
		{"1 3 2 4 5", 1},
		{"8 6 4 4 1", 1},
		{"1 3 6 7 9", 1},
	}
	for _, test := range testCases {
		result := part2(test.in)
		assert.Equal(t, test.out, result, fmt.Sprintf("they should be equal for line %s", test.in))
	}
}

func TestPart2Full(t *testing.T) {
	expected := 4
	actual := part2(input)
	assert.Equal(t, expected, actual)
}
