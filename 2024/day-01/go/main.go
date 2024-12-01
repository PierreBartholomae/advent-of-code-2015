package main

import (
	"os"
	"sort"
	"strconv"
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
	var lhsLocationIds []int
	var rhsLocationIds []int

	for _, line := range strings.Split(input, "\n") {
		if lhsStr, rhsStr, found := strings.Cut(line, "   "); found {
			if lhs, err := strconv.Atoi(lhsStr); err == nil {
				lhsLocationIds = append(lhsLocationIds, lhs)
			}
			if rhs, err := strconv.Atoi(rhsStr); err == nil {
				rhsLocationIds = append(rhsLocationIds, rhs)
			}
		}
	}

	sort.Ints(lhsLocationIds)
	sort.Ints(rhsLocationIds)

	result := 0
	for i := 0; i < len(lhsLocationIds); i++ {
		if lhsLocationIds[i] > rhsLocationIds[i] {
			result += lhsLocationIds[i] - rhsLocationIds[i]
		} else {
			result += rhsLocationIds[i] - lhsLocationIds[i]
		}
	}
	return result
}

func part2(input string) int {
	var lhsLocationIds []int
	rhsLocationIdFrequency := map[int]int{}

	for _, line := range strings.Split(input, "\n") {
		if lhsStr, rhsStr, found := strings.Cut(line, "   "); found {
			if lhs, err := strconv.Atoi(lhsStr); err == nil {
				lhsLocationIds = append(lhsLocationIds, lhs)
			}
			if rhs, err := strconv.Atoi(rhsStr); err == nil {
				rhsLocationIdFrequency[rhs] += 1
			}
		}
	}

	result := 0
	for _, lhs := range lhsLocationIds {
		rhsFrequency := rhsLocationIdFrequency[lhs]
		result += lhs * rhsFrequency
	}
	return result
}
