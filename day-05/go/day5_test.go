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
		{"ugknbfddgicrmopn", 1},
		{"aaa", 1},
		{"jchzalrnumimnmhp", 0},
		{"haegwjzuvuyypxyu", 0},
		{"dvszwmarrgswjxmb", 0},
	}
	for _, test := range input {
		result := part1(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestAtLeastThreeVowels(t *testing.T) {
	var input = []struct {
		input    string
		expected bool
	}{
		{"aei", true},
		{"xazegov", true},
		{"aeiouaeiouaeiou", true},
	}
	for _, test := range input {
		result := atLeastThreeVowels(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestTwiceSameLetter(t *testing.T) {
	var input = []struct {
		input    string
		expected bool
	}{
		{"xx", true},
		{"abcdde", true},
		{"aabbccdd", true},
	}
	for _, test := range input {
		result := twiceSameLetter(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestOnlyNicePairs(t *testing.T) {
	var input = []struct {
		input    string
		expected bool
	}{
		{"ab", false},
		{"cd", false},
		{"pq", false},
		{"xy", false},
	}
	for _, test := range input {
		result := onlyNicePairs(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestTwiceSameLetterAppearingTwice(t *testing.T) {
	var input = []struct {
		input    string
		expected bool
	}{
		{"xyxy", true},
		{"aabcdefgaa", true},
		{"aaa", false},
	}
	for _, test := range input {
		result := twiceSameLetterWithoutOverlapping(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestRepeatingLetterWithLetterBetween(t *testing.T) {
	var input = []struct {
		input    string
		expected bool
	}{
		{"xyx", true},
		{"abcdefeghi", true},
		{"aaa", true},
	}
	for _, test := range input {
		result := repeatingLetterWithLetterBetween(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

func TestPart2(t *testing.T) {
	var input = []struct {
		input    string
		expected int
	}{
		{"qjhvhtzxzqqjkmpb", 1},
		{"xxyxx", 1},
		{"uurcxstgmygtbstg", 0},
		{"ieodomkazucvgmuy", 0},
	}
	for _, test := range input {
		result := part2(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
