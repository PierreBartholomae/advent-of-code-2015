package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

const inputPart1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const inputPart2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestPart1Parts(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		{"xmul(2,4)", 8},
		{"%&mul[3,7]!@^do_not_mul(5,5)", 25},
		{"+mul(32,64]then(mul(11,8)", 88},
		{"mul(8,5))", 40},
		{"xmul(170,386why()", 0},
	}
	for _, test := range testCases {
		result := part1(test.in)
		assert.Equal(t, test.out, result, fmt.Sprintf("they should be equal for line %s", test.in))
	}
}

func TestPart1Full(t *testing.T) {
	expected := 161
	actual := part1(inputPart1)
	assert.Equal(t, expected, actual)
}

func TestPart2Parts(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		{"xmul(170,386why()", 0},
		{"xmul(886,864/[]how()", 0},
		{"xmul(695from()", 0},
		{"xdon't()mul(695from()mul(11,8)", 0},
		{"xdon't()mul(11,8)do()mul(5,5)", 25},
	}
	for _, test := range testCases {
		result := part2(test.in)
		assert.Equal(t, test.out, result, fmt.Sprintf("they should be equal for line %s", test.in))
	}
}

func TestPart2Full(t *testing.T) {
	expected := 48
	actual := part2(inputPart2)
	assert.Equal(t, expected, actual)
}
