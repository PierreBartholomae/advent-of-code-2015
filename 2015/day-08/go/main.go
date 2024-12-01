package main

import (
	"bufio"
	"fmt"
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
}

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	total := 0
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		line := scanner.Bytes()
		dLine := string(line)
		println(dLine)
		stringLiteralsLen := len(line)
		inMemoryLen := 0
		for i := 1; i < len(line)-1; i++ {
			currChar := line[i]
			switch {
			case currChar == 92:
				nextChar := line[i+1]
				if nextChar == 92 {
					// If \\, then only count \ as one
					i++
					inMemoryLen += 1
				} else if nextChar == 34 {
					// If \", then only count " as one
					i++
					inMemoryLen += 1
				} else if nextChar == 120 {
					// If \x47, then get the next two bytes and decode hex value
					i += 3
					inMemoryLen += 1
				}
			case (currChar >= 48 && currChar <= 57) || (currChar >= 97 && currChar <= 122):
				// If alphanumerical
				inMemoryLen += 1
			default:
				panic(fmt.Sprintf("Show me this case!, %s", string(currChar)))
			}
		}
		total += stringLiteralsLen - inMemoryLen
	}
	return total
}
