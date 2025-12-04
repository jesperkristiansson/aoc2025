package main

import (
	"bufio"
	"fmt"
	"os"
)

const paperRune = '@'
const maxAdjacentPapers = 3

func getGrid() [][]bool {
	var grid [][]bool

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		gridLine := make([]bool, 0, len(line))
		for _, r := range line {
			gridLine = append(gridLine, r == paperRune)
		}

		grid = append(grid, gridLine)
	}

	return grid
}

// Assumes that the grid is not empty and that each row is the same length
func isInsideGrid(grid [][]bool, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func isPaper(grid [][]bool, r, c int) bool {
	if !isInsideGrid(grid, r, c) {
		return false
	}
	return grid[r][c]
}

func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func countAdjacentPapers(grid [][]bool, r, c int) int {
	papers := 0

	papers += boolToInt(isPaper(grid, r-1, c-1))
	papers += boolToInt(isPaper(grid, r-1, c))
	papers += boolToInt(isPaper(grid, r-1, c+1))
	papers += boolToInt(isPaper(grid, r, c-1))
	papers += boolToInt(isPaper(grid, r, c+1))
	papers += boolToInt(isPaper(grid, r+1, c-1))
	papers += boolToInt(isPaper(grid, r+1, c))
	papers += boolToInt(isPaper(grid, r+1, c+1))

	return papers
}

func part1(grid [][]bool) int {
	result := 0
	for r := range grid {
		for c := range grid[r] {
			if isPaper(grid, r, c) && countAdjacentPapers(grid, r, c) <= maxAdjacentPapers {
				result++
			}
		}
	}
	return result
}

// Removes as many rolls of paper as possible and returns the number removed
func removePaper(grid [][]bool) int {
	result := 0
	for r := range grid {
		for c := range grid[r] {
			if isPaper(grid, r, c) && countAdjacentPapers(grid, r, c) <= maxAdjacentPapers {
				result++
				grid[r][c] = false
			}
		}
	}
	return result
}

func part2(grid [][]bool) int {
	result := 0
	for {
		papersRemoved := removePaper(grid)
		if papersRemoved == 0 {
			break
		}
		result += papersRemoved
	}
	return result
}

func main() {
	grid := getGrid()
	fmt.Println("part 1:", part1(grid))
	fmt.Println("part 2:", part2(grid))
}
