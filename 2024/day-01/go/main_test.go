package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPart1(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		{in: "3   4\n4   3\n2   5\n1   3\n3   9\n3   3", out: 11},
	}
	for _, test := range testCases {
		result := part1(test.in)
		assert.Equal(t, test.out, result, "they should be equal")
	}
}

func TestPart2(t *testing.T) {
	var testCases = []struct {
		in  string
		out int
	}{
		{in: "3   4\n4   3\n2   5\n1   3\n3   9\n3   3", out: 31},
	}
	for _, test := range testCases {
		result := part2(test.in)
		assert.Equal(t, test.out, result, "they should be equal")
	}
}
