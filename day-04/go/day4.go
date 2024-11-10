package main

import (
	"crypto/md5"
	"encoding/hex"
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

func part1(secret string) int {
	return findChecksumMatchingPrefix(secret, "00000")
}

func part2(secret string) int {
	return findChecksumMatchingPrefix(secret, "000000")
}

func findChecksumMatchingPrefix(secret string, prefix string) int {
	h := md5.New()

	for index := 0; index < math.MaxInt; index++ {
		// Reset the previous checksum
		h.Reset()

		// create a new hash with the current index
		hashInput := secret + strconv.Itoa(index)

		// create the hash
		h.Write([]byte(hashInput))

		// get the checksum for the hash
		checksum := h.Sum(nil)
		checksumString := hex.EncodeToString(checksum[:])

		// check, if first five characters of the checksum are zeroes
		if strings.HasPrefix(checksumString, prefix) {
			fmt.Printf("Checksum %s is a match! \n", checksumString)
			return index
		}
	}

	panic(fmt.Sprintf("No index found leading to a checksum starting prefix %s.", prefix))
}
