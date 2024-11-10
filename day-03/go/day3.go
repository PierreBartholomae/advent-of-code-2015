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
	println("answer part 1:", resultPart1)

	resultPart2 := part2(input)
	println("answer part 2:", resultPart2)
}

// Position of a house
type Position struct {
	x int
	y int
}

/*
Input
^>v<

Output
oo
xo

0.0 -> 0.1 -> 1.1. -> 1.0. -> 0.0

Input
^v^v^v^v^v"

Output
x
x
*/
func part1(input string) int {
	var visits = map[Position]bool{
		{0, 0}: true,
	}

	var currentPos = Position{
		0, 0,
	}
	for _, char := range input {
		// calculate new position
		newPos := currentPos
		switch char {
		case '^':
			newPos.y += 1
		case 'v':
			newPos.y -= 1
		case '>':
			newPos.x += 1
		case '<':
			newPos.x -= 1
		default:
			panic(fmt.Sprintf("Character %c not allowed!", char))
		}

		// Add position to visits
		visits[newPos] = true

		// update current position for next iteration
		currentPos = newPos
	}

	return len(visits)
}

func part2(input string) int {
	var visits = map[Position]bool{
		{0, 0}: true,
	}

	var currentPosSanta = Position{
		0, 0,
	}
	var currentPosRoboSanta = Position{
		0, 0,
	}

	for index, char := range input {
		// calculate new position
		newPosSanta := currentPosSanta
		newPosRoboSanta := currentPosRoboSanta

		if index%2 == 0 {
			switch char {
			case '^':
				newPosSanta.y += 1
			case 'v':
				newPosSanta.y -= 1
			case '>':
				newPosSanta.x += 1
			case '<':
				newPosSanta.x -= 1
			default:
				panic(fmt.Sprintf("Character %c not allowed!", char))
			}
			// Add position to visits
			visits[newPosSanta] = true

			// update current position for next iteration
			currentPosSanta = newPosSanta
		} else {
			switch char {
			case '^':
				newPosRoboSanta.y += 1
			case 'v':
				newPosRoboSanta.y -= 1
			case '>':
				newPosRoboSanta.x += 1
			case '<':
				newPosRoboSanta.x -= 1
			default:
				panic(fmt.Sprintf("Character %c not allowed!", char))
			}
			// Add position to visits
			visits[newPosRoboSanta] = true

			// update current position for next iteration
			currentPosRoboSanta = newPosRoboSanta
		}
	}

	return len(visits)
}
