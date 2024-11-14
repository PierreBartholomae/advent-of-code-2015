package main

import (
	"fmt"
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

type Position struct {
	x int
	y int
}

func NewPosition(str string) Position {
	parts := strings.Split(str, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Sprintf("x value could not be parsed %v", err))
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("y value could not be parsed %v", err))
	}
	return Position{x, y}
}

type Instruction int

const (
	On Instruction = iota
	Off
	Toggle
)

var instructions = map[Instruction]string{
	On:     "on",
	Off:    "off",
	Toggle: "toggle",
}

func NewInstruction(str string) Instruction {
	switch str {
	case "on":
		return On
	case "off":
		return Off
	case "toggle":
		return Toggle
	default:
		panic(fmt.Sprintf("The value %s is no valid instruction", str))
	}
}

type Action struct {
	startPos Position
	endPos   Position
	instr    Instruction
}

func NewAction(line string) Action {
	parts := strings.Split(line, " ")
	if len(parts) == 5 {
		parts = parts[1:]
	}
	instr := NewInstruction(parts[0])
	startPos := NewPosition(parts[1])
	endPos := NewPosition(parts[3])
	return Action{startPos, endPos, instr}
}

func part1(input string) int {
	lights := [1000][1000]bool{}
	for _, line := range strings.Split(input, "\n") {
		action := NewAction(line)
		rows := action.startPos.x + (action.endPos.x - action.startPos.x) + 1
		cols := action.startPos.y + (action.endPos.y - action.startPos.y) + 1
		for rowIndex := action.startPos.x; rowIndex < rows; rowIndex++ {
			for colIndex := action.startPos.y; colIndex < cols; colIndex++ {
				switch action.instr {
				case On:
					lights[rowIndex][colIndex] = true
				case Off:
					lights[rowIndex][colIndex] = false
				case Toggle:
					if lights[rowIndex][colIndex] {
						lights[rowIndex][colIndex] = false

					} else {
						lights[rowIndex][colIndex] = true
					}
				default:
					panic("This should not happen")
				}
			}
		}

	}

	lightsOn := 0
	for rowIndex := 0; rowIndex < 1000; rowIndex++ {
		for colIndex := 0; colIndex < 1000; colIndex++ {
			if lights[rowIndex][colIndex] {
				lightsOn += 1
			}
		}
	}

	return lightsOn
}

func part2(input string) int {
	lights := [1000][1000]int{}
	for _, line := range strings.Split(input, "\n") {
		action := NewAction(line)
		rows := action.startPos.x + (action.endPos.x - action.startPos.x) + 1
		cols := action.startPos.y + (action.endPos.y - action.startPos.y) + 1
		for rowIndex := action.startPos.x; rowIndex < rows; rowIndex++ {
			for colIndex := action.startPos.y; colIndex < cols; colIndex++ {
				switch action.instr {
				case On:
					lights[rowIndex][colIndex] += 1
				case Off:
					if lights[rowIndex][colIndex] > 0 {
						lights[rowIndex][colIndex] -= 1
					}
				case Toggle:
					lights[rowIndex][colIndex] += 2
				default:
					panic("This should not happen")
				}
			}
		}

	}

	lightsOn := 0
	for rowIndex := 0; rowIndex < 1000; rowIndex++ {
		for colIndex := 0; colIndex < 1000; colIndex++ {
			lightsOn += lights[rowIndex][colIndex]
		}
	}

	return lightsOn
}
