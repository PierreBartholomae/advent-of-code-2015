package main

import (
	"os"
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
	result := 0
	for _, massStr := range strings.Split(input, "\n") {
		mass, err := strconv.Atoi(massStr)
		if err == nil {
			requiredFuel := calculateFuel(mass)
			result += requiredFuel
		}
	}
	return result
}

func part2(input string) int {
	result := 0
	for _, massStr := range strings.Split(input, "\n") {
		mass, err := strconv.Atoi(massStr)
		if err == nil {
			requiredFuel := calculateFuelRecursive(mass)
			result += requiredFuel
		}
	}
	return result
}

func calculateFuel(mass int) int {
	return int(mass/3) - 2
}

func calculateFuelRecursive(mass int) int {
	calculatedFuel := 0
	fuel := int(mass/3) - 2
	calculatedFuel += fuel
	if fuel > 0 {
		fuelMass := calculateFuelRecursive(fuel)
		calculatedFuel += fuelMass
		return calculatedFuel
	} else {
		return 0
	}
}
