package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPart1(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
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
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}
	for _, test := range input {
		result := part2(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
