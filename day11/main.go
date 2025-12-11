package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type computer struct {
	name    string
	outputs []string
}

type input []computer
type adjacencyMap map[string][]string

func getInput() input {
	var input input

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		spl := strings.Split(line, " ")
		if len(spl) < 2 {
			panic("Expected at least 2 tokens in input")
		}
		from := spl[0][:3]
		to := spl[1:]

		c := computer{name: from, outputs: to}
		input = append(input, c)
	}

	return input
}

func getAdjacencyMap(computers input) adjacencyMap {
	m := make(adjacencyMap, len(computers))

	for _, c := range computers {
		m[c.name] = c.outputs
	}

	return m
}

type boolMap map[string]bool

func dfs(am adjacencyMap, visited boolMap, current, target string) int {
	if current == target {
		return 1
	}

	if visited[current] {
		return 0
	}
	visited[current] = true

	sum := 0
	for _, adj := range am[current] {
		sum += dfs(am, visited, adj, target)
	}

	visited[current] = false
	return sum
}

func part1(am adjacencyMap) int {
	visited := make(boolMap, len(am))

	return dfs(am, visited, "you", "out")
}

type dfsResult struct {
	pathsWithDac  int
	pathsWithFft  int
	pathsWithBoth int
	pathsWithNone int
}
type cacheData struct {
	isRunning bool
	result    dfsResult
}

type cache map[string]cacheData

func dfs2(am adjacencyMap, cache cache, current, target string) (result dfsResult) {
	if current == target {
		return dfsResult{pathsWithNone: 1}
	}

	currentCache, ok := cache[current]
	if ok {
		// Detect if we're following a loop
		if currentCache.isRunning {
			return
		} else {
			return currentCache.result
		}
	}
	currentCache.isRunning = true
	cache[current] = currentCache

	subResultSum := dfsResult{}
	for _, adj := range am[current] {
		subResult := dfs2(am, cache, adj, target)
		subResultSum.pathsWithBoth += subResult.pathsWithBoth
		subResultSum.pathsWithDac += subResult.pathsWithDac
		subResultSum.pathsWithFft += subResult.pathsWithFft
		subResultSum.pathsWithNone += subResult.pathsWithNone
	}

	switch current {
	case "dac":
		result.pathsWithDac = subResultSum.pathsWithDac + subResultSum.pathsWithNone
		result.pathsWithBoth = subResultSum.pathsWithBoth + subResultSum.pathsWithFft
	case "fft":
		result.pathsWithFft = subResultSum.pathsWithFft + subResultSum.pathsWithNone
		result.pathsWithBoth = subResultSum.pathsWithBoth + subResultSum.pathsWithDac
	default:
		result = subResultSum
	}

	currentCache.isRunning = false
	currentCache.result = result
	cache[current] = currentCache
	return
}

func part2(am adjacencyMap) int {
	cache := make(cache, len(am))

	result := dfs2(am, cache, "svr", "out")
	return result.pathsWithBoth
}

func main() {
	input := getInput()
	am := getAdjacencyMap(input)
	fmt.Println("part 1:", part1(am))
	fmt.Println("part 2:", part2(am))
}
