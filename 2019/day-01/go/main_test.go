package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

const inputPart1 = "12\n14\n1969\n100756"
const inputPart2 = "12\n1969\n100756"

func TestPart1Parts(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		{"12", 2},
		{"14", 2},
		{"1969", 654},
		{"100756", 33583},
	}
	for _, test := range testCases {
		result := part1(test.in)
		assert.Equal(t, test.out, result, fmt.Sprintf("they should be equal for line %s", test.in))
	}
}

func TestPart1Full(t *testing.T) {
	expected := 2 + 2 + 654 + 33583
	actual := part1(inputPart1)
	assert.Equal(t, expected, actual)
}

func TestPart2(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		{"12", 2},
		{"1969", 966},
		{"100756", 50346},
	}
	for _, test := range testCases {
		result := part2(test.in)
		assert.Equal(t, test.out, result, fmt.Sprintf("they should be equal for line %s", test.in))
	}
}

func TestPart2Full(t *testing.T) {
	expected := 2 + 966 + 50346
	actual := part2(inputPart2)
	assert.Equal(t, expected, actual)
}
