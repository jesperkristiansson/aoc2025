package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
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

func det(p1, p2 point2d) int {
	return p1.x*p2.y - p2.x*p1.y
}

func orientation(points input) int {
	sum := 0
	for i := 0; i < len(points)-1; i++ {
		sum += det(points[i], points[i+1])
	}
	sum += det(points[len(points)-1], points[0])

	if sum == 0 {
		panic("Ill-formed polygon")
	}

	return sum
}

type xLine struct {
	x            int
	startY, endY int
}

type yLine struct {
	y            int
	startX, endX int
}

func getLines(points input) ([]xLine, []yLine) {
	var xLines []xLine
	var yLines []yLine

	for i := 0; i < len(points); i++ {
		nextI := i + 1
		if nextI == len(points) {
			nextI = 0
		}
		p1 := points[i]
		p2 := points[nextI]
		switch {
		case p1.x == p2.x:
			lowY := min(p1.y, p2.y)
			highY := max(p1.y, p2.y)
			line := xLine{x: p1.x, startY: lowY, endY: highY}
			xLines = append(xLines, line)
		case p1.y == p2.y:
			lowX := min(p1.x, p2.x)
			highX := max(p1.x, p2.x)
			line := yLine{y: p1.y, startX: lowX, endX: highX}
			yLines = append(yLines, line)
		default:
			panic("points do not share x- or y-coordinates")
		}
	}

	return xLines, yLines
}

func part2(points input) int {
	// Not needed for the input data
	orientation := orientation(points)
	isRight := orientation > 0
	_ = isRight

	xLines, yLines := getLines(points)

	slices.SortFunc(xLines, func(a, b xLine) int { return cmp.Compare(a.x, b.x) })
	slices.SortFunc(yLines, func(a, b yLine) int { return cmp.Compare(a.y, b.y) })

	maxArea := 0

	for i, p1 := range points {
	loop:
		for _, p2 := range points[i+1:] {
			// TODO: check orientation

			topLeft := point2d{x: min(p1.x, p2.x), y: min(p1.y, p2.y)}
			bottomRight := point2d{x: max(p1.x, p2.x), y: max(p1.y, p2.y)}

			lowXIdx, _ := slices.BinarySearchFunc(xLines, topLeft.x+1, func(a xLine, x int) int { return cmp.Compare(a.x, x) })
			highXIdx, _ := slices.BinarySearchFunc(xLines, bottomRight.x, func(a xLine, x int) int { return cmp.Compare(a.x, x) })
			for xIdx := lowXIdx; xIdx < highXIdx; xIdx++ {
				line := xLines[xIdx]
				if !(line.startY >= bottomRight.y || line.endY <= topLeft.y) {
					continue loop
				}
			}

			lowYIdx, _ := slices.BinarySearchFunc(yLines, topLeft.y+1, func(a yLine, y int) int { return cmp.Compare(a.y, y) })
			highYIdx, _ := slices.BinarySearchFunc(yLines, bottomRight.y, func(a yLine, y int) int { return cmp.Compare(a.y, y) })
			for yIdx := lowYIdx; yIdx < highYIdx; yIdx++ {
				line := yLines[yIdx]
				if !(line.startX >= bottomRight.x || line.endX <= topLeft.x) {
					continue loop
				}
			}

			area := rectangleArea(p1, p2)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	input := getInput()
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}
