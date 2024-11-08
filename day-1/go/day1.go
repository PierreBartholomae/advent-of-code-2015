package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		panic("File at path ./input could not be read.")
	}
	input := string(bytes)
	
	resultPart1 := part1(input)
	println(resultPart1)

	resultPart2 := part2(input)
	println(resultPart2)
}

func part1(input string) int {
	var result = 0
	for _, char := range input {
		switch char {
		case '(':
			result++
		case ')':
			result--
		default:
			panic(fmt.Sprintf("Char %c not supported", char))
		}
	}
	return result
}

func part2(input string) int {
	var result = 0
	for index, char := range input {
		switch char {
		case '(':
			result++
		case ')':
			result--
		default:
			panic(fmt.Sprintf("Char %c not supported", char))
		}
		if result == -1 {
			return index + 1
		}
	}
	panic("Could not find the basement -1")
}
