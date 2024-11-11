package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPart1(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}
	for _, test := range input {
		result := part1(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
