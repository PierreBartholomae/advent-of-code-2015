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
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		nums := extractNums(line)
		sum += isSafe(nums)
	}
	return sum
}

func part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		nums := extractNums(line)
		res := isSafe(nums)
		if res == 1 {
			sum += 1
		} else {
			for i := range len(nums) {
				var numsCopy []int
				numsCopy = append(numsCopy, nums[:i]...)
				numsCopy = append(numsCopy, nums[i+1:]...)
				res := isSafe(numsCopy)
				if res == 1 {
					sum += 1
					break
				}
			}
		}
	}
	return sum
}

func extractNums(line string) []int {
	var nums []int
	for _, numStr := range strings.Fields(line) {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func isSafe(nums []int) int {
	safe := true
	onlyInc := true
	onlyDec := true
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff > 0 {
			onlyDec = false
		}
		if diff < 0 {
			onlyInc = false
		}
		if Abs(diff) < 1 || Abs(diff) > 3 {
			safe = false
		}
	}
	if safe && (onlyInc || onlyDec) {
		return 1
	}
	return 0
}

func Abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
