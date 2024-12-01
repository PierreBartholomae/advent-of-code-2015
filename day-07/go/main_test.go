package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		input    string
		expected uint16
	}{
		{"123 -> x\n456 -> y\nx AND y -> d\nx OR y -> e\nx LSHIFT 2 -> f\ny RSHIFT 2 -> g\nNOT x -> h\nNOT y -> i\nh -> a",
			uint16(65412),
		},
	}
	for _, tt := range tests {
		result := part1(tt.input)
		assert.Equal(t, tt.expected, result, "they should be equal")
	}
}
