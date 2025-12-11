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

func part2(am adjacencyMap) int {
	sum := 0

	return sum
}

func main() {
	input := getInput()
	am := getAdjacencyMap(input)
	fmt.Println("part 1:", part1(am))
	fmt.Println("part 2:", part2(am))
}
