package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestDay1Part1(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}
	for _, test := range input {
		result := part1(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestDay1Part2(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{")", 1},
		{"()())", 5},
	}
	for _, test := range input {
		result := part2(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
