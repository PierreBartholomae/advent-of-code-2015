package graph

import (
	"math"
	"strconv"
	"strings"
)

type Graph map[string]map[string]int

func NewGraph(input string) Graph {
	graph := Graph{}
	for _, line := range strings.Split(input, "\n") {
		// parse line
		parts := strings.Split(line, " ")
		from := parts[0]
		to := parts[2]
		weight, err := strconv.Atoi(parts[4])
		if err != nil {
			panic("Weight could not be parsed")
		}

		// create map for neighbors, if not existing
		if graph[from] == nil {
			graph[from] = make(map[string]int)
		}
		if graph[to] == nil {
			graph[to] = make(map[string]int)
		}

		// set weight for both directions
		graph[from][to] = weight
		graph[to][from] = weight
	}
	return graph
}

func WalkAllDirections(graph Graph) (min, max int) {
	minTotal := math.MaxInt32
	maxTotal := 0
	for node := range graph {
		minDistance, maxDistance := WalkBreathFirst(node, graph, map[string]bool{node: true})
		minTotal = Min(minTotal, minDistance)
		maxTotal = Max(maxTotal, maxDistance)
	}
	return minTotal, maxTotal
}

func WalkBreathFirst(from string, graph Graph, visited map[string]bool) (min, max int) {
	if len(visited) == len(graph) {
		return 0, 0
	}
	minDistance := math.MaxInt32
	maxDistance := 0
	for to := range graph {
		if visited[to] {
			continue
		}
		visited[to] = true

		weight := graph[from][to]

		minRecursive, maxRecursive := WalkBreathFirst(to, graph, visited)
		minDistance = Min(minDistance, minRecursive+weight)
		maxDistance = Max(maxDistance, maxRecursive+weight)

		delete(visited, to)
	}

	return minDistance, maxDistance
}

func Min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func Max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
