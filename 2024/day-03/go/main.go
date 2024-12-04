package main

import (
	"os"
	"unicode"
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
	sum := 0
	n := len(input)
	for i := 0; i < n; i++ {
		if input[i] == 'm' {
			if input[i+1] == 'u' && input[i+2] == 'l' && input[i+3] == '(' {
				i += 4
				lhs := getNumber(input, &i)
				if input[i] == ',' {
					i += 1
					rhs := getNumber(input, &i)
					if input[i] == ')' {
						sum += lhs * rhs
					}
				}
			}
		}
	}
	return sum
}

func part2(input string) int {
	sum := 0
	n := len(input)
	enabled := true
	for i := 0; i < n; i++ {
		if input[i] == 'm' {
			if input[i+1] == 'u' && input[i+2] == 'l' && input[i+3] == '(' {
				i += 4
				lhs := getNumber(input, &i)
				if input[i] == ',' {
					i += 1
					rhs := getNumber(input, &i)
					if input[i] == ')' && enabled {
						sum += lhs * rhs
					}
				}
			}
		} else if input[i] == 'd' {
			if input[i+1] == 'o' && input[i+2] == 'n' && input[i+3] == '\'' && input[i+4] == 't' {
				i += 5
				enabled = false
			} else if input[i+1] == 'o' {
				i += 2
				enabled = true
			}
		}
	}
	return sum
}

func getNumber(input string, i *int) int {
	num := 0
	for num < 1000 && unicode.IsDigit(rune(input[*i])) {
		// (10 times the previous number) + (ascii decimal number at i - ascii decimal number of 0 (which is 48))
		// Example: 695
		// 1 iteration: 10*0 + 54-48 = 6
		// 2 iteration: 10*6 + 57-48 = 69
		// 3 iteration: 10*69 + 53-48 = 695
		num = 10*num + int(input[*i]-'0')
		*i += 1
	}
	return num
}
