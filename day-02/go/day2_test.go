package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestPart1(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}
	for _, test := range input {
		present := Present{NewDimensions(test.input)}
		result := present.Wrapping()
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestPart2(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}
	for _, test := range input {
		present := Present{NewDimensions(test.input)}
		result := present.Ribbon()
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
