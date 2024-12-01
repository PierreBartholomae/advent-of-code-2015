package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

func TestNewGraphFromInput(t *testing.T) {
	testCases := []struct {
		from   string
		to     string
		weight int
	}{
		{"London", "Dublin", 464},
		{"Dublin", "London", 464},
		{"London", "Belfast", 518},
		{"Belfast", "London", 518},
		{"Dublin", "Belfast", 141},
		{"Belfast", "Dublin", 141},
	}

	g := NewGraph(input)
	for _, tc := range testCases {
		_, ok := g[tc.from]
		assert.Equal(t, true, ok)

		weight, ok := g[tc.from][tc.to]
		assert.Equal(t, true, ok)

		assert.Equal(t, tc.weight, weight)
	}
}

func TestWalkBreathFirst(t *testing.T) {
	graph := NewGraph(input)
	visited := map[string]bool{}

	minDistanceLondon := WalkBreathFirst("London", graph, visited)
	assert.Equal(t, 605, minDistanceLondon)

	minDistanceDublin := WalkBreathFirst("Dublin", graph, visited)
	assert.Equal(t, 659, minDistanceDublin)

	minDistanceBelfast := WalkBreathFirst("Belfast", graph, visited)
	assert.Equal(t, 605, minDistanceBelfast)
}

func TestWalkAllDirections(t *testing.T) {
	graph := NewGraph(input)

	minDistance := WalkAllDirections(graph)
	assert.Equal(t, 605, minDistance)
}
