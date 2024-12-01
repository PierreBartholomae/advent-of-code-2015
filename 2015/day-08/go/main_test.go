package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	bytes, err := os.ReadFile("test")
	if err != nil {
		panic("File at path ./test could not be read.")
	}

	expected := 12
	input := string(bytes)
	result := part1(input)
	assert.Equal(t, expected, result, "they should be equal")

}
