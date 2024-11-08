package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic("File at path ./input could not be read.")
	}
	defer file.Close()

	resultPart1 := part1(file)
	println(resultPart1)

	file, err = os.Open("input")
	if err != nil {
		panic("File at path ./input could not be read.")
	}
	defer file.Close()
	resultPart2 := part2(file)
	println(resultPart2)
}

func part1(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	var totalWrapping = 0
	for scanner.Scan() {
		line := scanner.Text()
		present := Present{NewDimensions(line)}
		wrapping := present.Wrapping()
		totalWrapping += wrapping
	}
	return totalWrapping
}

func part2(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	var totalRibbon = 0
	for scanner.Scan() {
		line := scanner.Text()
		present := Present{NewDimensions(line)}
		ribbon := present.Ribbon()
		totalRibbon += ribbon
	}
	return totalRibbon
}

type Dimensions struct {
	l int // length
	w int // width
	h int // height
}

func NewDimensions(rawDimensions string) *Dimensions {
	var dimensions = strings.Split(rawDimensions, "x")
	l, err := strconv.Atoi(dimensions[0])
	if err != nil {
		panic(fmt.Sprintf("l input [%s] could not be parsed for input %s", dimensions[0], rawDimensions))
	}
	w, err := strconv.Atoi(dimensions[1])
	if err != nil {
		panic(fmt.Sprintf("w input [%s] could not be parsed for input %s", dimensions[1], rawDimensions))
	}
	h, err := strconv.Atoi(dimensions[2])
	if err != nil {
		panic(fmt.Sprintf("h input [%s] could not be parsed for input %s", dimensions[2], rawDimensions))
	}
	d := Dimensions{l, w, h}
	return &d
}

type Present struct {
	dimensions *Dimensions
}

func (p *Present) Wrapping() int {
	d := p.dimensions
	var surfaceArea = 2*d.l*d.w + 2*d.w*d.h + 2*d.h*d.l
	var slack = int(math.Min(math.Min(float64(d.l*d.w), float64(d.w*d.h)), float64(d.h*d.l)))
	var squareFeetOfWrapping = surfaceArea + slack
	return squareFeetOfWrapping
}

func (p *Present) Ribbon() int {
	d := p.dimensions
	var sides = []int{d.l, d.w, d.h}
	sort.Ints(sides)
	var wrapRibbonLength = sides[0] + sides[0] + sides[1] + sides[1]
	var bowRibbonLength = d.l * d.w * d.h
	var feetOfRibbon = wrapRibbonLength + bowRibbonLength
	return feetOfRibbon
}
