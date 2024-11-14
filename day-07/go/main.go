package main

import (
	"bufio"
	"fmt"
	"math"
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

type WireParser struct {
	circuit map[string]string
	wires   map[string]uint16
}

func NewWireParser() *WireParser {
	return &WireParser{
		circuit: map[string]string{},
		wires:   map[string]uint16{},
	}
}

func (wp *WireParser) ParseInput(input string) {
	fileScanner := bufio.NewScanner(strings.NewReader(input))
	fileScanner.Split(bufio.ScanLines)
	for {
		scannedLine := fileScanner.Scan()
		if !scannedLine {
			break
		}

		line := fileScanner.Text()
		parts := strings.Split(line, " -> ")
		wp.circuit[parts[1]] = parts[0]
	}
}

func (wp *WireParser) TraverseCircuit(wire string) uint16 {
	// If the left hand side is a number, return it
	if v, err := strconv.Atoi(wire); err == nil {
		return uint16(v)
	}

	// if a wire is requested multiple times, return it from the cache
	if v, ok := wp.wires[wire]; ok {
		return v
	}

	instructions := wp.circuit[wire]

	words := strings.Split(instructions, " ")

	var result uint16
	switch len(words) {
	case 1:
		// 123 -> x
		result = wp.TraverseCircuit(words[0])
	case 2:
		// XOR: NOT x -> h
		lhs := wp.TraverseCircuit(words[1])
		result = math.MaxUint16 ^ lhs
	case 3:
		lhs := wp.TraverseCircuit(words[0])
		rhs := wp.TraverseCircuit(words[2])
		operation := words[1]
		switch operation {
		case "AND":
			// AND: x AND y -> d
			result = lhs & rhs
		case "OR":
			// OR: x OR y -> e
			result = lhs | rhs
		case "LSHIFT":
			// <<: x LSHIFT 2 -> f
			result = lhs << rhs
		case "RSHIFT":
			// >>: y RSHIFT 2 -> g
			result = lhs >> rhs
		default:
			panic(fmt.Sprintf("operation %s is not supported", operation))
		}
	default:
		panic(fmt.Sprintf("command with len %d not implemented", len(words)))
	}

	// update the cache
	wp.wires[wire] = result

	return result
}

func (wp *WireParser) Reset() {
	wp.wires = map[string]uint16{}
}

func part1(input string) uint16 {
	wp := NewWireParser()
	wp.ParseInput(input)

	return wp.TraverseCircuit("a")
}

func part2(input string) uint16 {
	wp := NewWireParser()
	wp.ParseInput(input)

	result := wp.TraverseCircuit("a")

	wp.Reset()
	wp.circuit["b"] = strconv.Itoa(int(result))

	return wp.TraverseCircuit("a")
}
