package main

import (
	"bufio"
	"fmt"
	"os"
)

type diagram []string

func getDiagram() diagram {
	var diagram diagram

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		diagram = append(diagram, line)
	}

	return diagram
}

func findStart(diagram diagram) (int, int) {
	for r, row := range diagram {
		for c, symbol := range row {
			if symbol == 'S' {
				return r, c
			}
		}
	}

	panic("Found no start in diagram")
}

func isInsideDiagram(diagram diagram, r, c int) bool {
	return r >= 0 && r < len(diagram) && c >= 0 && c < len(diagram[0])
}

func searchRecursive(diagram diagram, visited [][]bool, r, c int) int {
	if !isInsideDiagram(diagram, r, c) {
		return 0
	}
	if visited[r][c] {
		return 0
	}
	visited[r][c] = true

	sum := 0
	if diagram[r][c] == '^' {
		sum++
		// Search diagonals down
		sum += searchRecursive(diagram, visited, r+1, c-1)
		sum += searchRecursive(diagram, visited, r+1, c+1)
	} else {
		// Search straight down
		sum += searchRecursive(diagram, visited, r+1, c)
	}

	return sum
}

func part1(diagram diagram) int {
	startR, startC := findStart(diagram)

	visited := make([][]bool, len(diagram))
	for r := range diagram {
		visited[r] = make([]bool, len(diagram[r]))
	}

	return searchRecursive(diagram, visited, startR, startC)
}

func searchRecursive2(diagram diagram, timelines [][]int, r, c int) int {
	if !isInsideDiagram(diagram, r, c) {
		return 1
	}
	if timelines[r][c] != 0 {
		return timelines[r][c]
	}

	sum := 0
	if diagram[r][c] == '^' {
		// Search diagonals down
		sum += searchRecursive2(diagram, timelines, r+1, c-1)
		sum += searchRecursive2(diagram, timelines, r+1, c+1)
	} else {
		// Search straight down
		sum += searchRecursive2(diagram, timelines, r+1, c)
	}

	timelines[r][c] = sum
	return sum
}

func part2(diagram diagram) int {
	startR, startC := findStart(diagram)

	timelines := make([][]int, len(diagram))
	for r := range diagram {
		timelines[r] = make([]int, len(diagram[r]))
	}

	return searchRecursive2(diagram, timelines, startR, startC)
}

func main() {
	diagram := getDiagram()
	fmt.Println("part 1:", part1(diagram))
	fmt.Println("part 2:", part2(diagram))
}
