package main

import (
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		panic("File at path ./input could not be read.")
	}
	input := string(bytes)

	resultPart1 := part1(input)
	println("answer part 1:", resultPart1)

	resultPart2 := part2(input)
	println("answer part 2:", resultPart2)
}

func part1(input string) int {
	niceStrings := 0
	for _, text := range strings.Split(input, "\n") {
		if onlyNicePairs(text) && twiceSameLetter(text) && atLeastThreeVowels(text) {
			niceStrings += 1
		}
	}
	return niceStrings
}

func part2(input string) int {
	niceStrings := 0
	for _, text := range strings.Split(input, "\n") {
		if twiceSameLetterWithoutOverlapping(text) && repeatingLetterWithLetterBetween(text) {
			niceStrings += 1
		}
	}
	return niceStrings
}

func atLeastThreeVowels(text string) bool {
	niceVowels := 0
	for _, char := range text {
		if strings.ContainsRune("aeiou", char) {
			niceVowels += 1
		}
	}
	return niceVowels >= 3
}

func twiceSameLetter(text string) bool {
	for index := 0; index < len(text)-1; index++ {
		leftChar := text[index]
		rightChar := text[index+1]
		if leftChar == rightChar {
			return true
		}
	}
	return false
}

func onlyNicePairs(text string) bool {
	naughtyPairs := []string{
		"ab", "cd", "pq", "xy",
	}
	for _, naughtyPair := range naughtyPairs {
		if strings.Contains(text, naughtyPair) {
			return false
		}
	}
	return true
}

func twiceSameLetterWithoutOverlapping(text string) bool {
	rightBound := len(text) - 1
	for outerIndex := 0; outerIndex < rightBound; outerIndex++ {
		pair := text[outerIndex : outerIndex+2]
		// index + 2, so that it is not overlapping
		// Min integer of index+2 or text length-1 to stay in bounds
		for innerIndex := minInt(outerIndex+2, rightBound); innerIndex < rightBound; innerIndex++ {
			innerPair := text[innerIndex : innerIndex+2]
			if pair == innerPair {
				return true
			}
		}

	}
	return false
}

func repeatingLetterWithLetterBetween(text string) bool {
	rightBound := len(text) - 2
	for index := 0; index < rightBound; index++ {
		leftChar := text[index]
		rightChar := text[index+2]
		if leftChar == rightChar {
			return true
		}

	}
	return false
}

func minInt(left int, right int) int {
	if left < right {
		return left
	}
	return right
}
