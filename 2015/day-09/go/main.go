package main

import (
	"aoc-2015-day-9/graph"
	"os"
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
	g := graph.NewGraph(input)
	minDistance, _ := graph.WalkAllDirections(g)
	return minDistance
}

func part2(input string) int {
	g := graph.NewGraph(input)
	_, maxDistance := graph.WalkAllDirections(g)
	return maxDistance
}
