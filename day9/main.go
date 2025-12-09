package main

import (
	"bufio"
	"fmt"
	"os"
)

type point2d struct {
	x, y int
}

type input []point2d

func getInput() input {
	var input input

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var p point2d
		n, err := fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		if err != nil {
			panic(err)
		}
		if n != 2 {
			panic(fmt.Errorf("expected 2 numbers, got %d", n))
		}

		input = append(input, p)
	}

	return input
}

func dist(n1, n2 int) int {
	d := n2 - n1
	if d < 0 {
		d = -d
	}
	return d + 1
}

func rectangleArea(p1, p2 point2d) int {
	dx := dist(p2.x, p1.x)
	dy := dist(p2.y, p1.y)
	area := dx * dy
	return area
}

func part1(input input) int {
	max := 0

	for i, p1 := range input {
		for _, p2 := range input[i+1:] {
			area := rectangleArea(p1, p2)
			if area > max {
				max = area
			}
		}
	}

	return max
}

func part2(input input) int {
	sum := 0

	return sum
}

func main() {
	input := getInput()
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}
