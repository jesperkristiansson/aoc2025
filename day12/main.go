package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type shape [][]bool
type problem struct {
	width, height int
	shapes        []int
}

type input struct {
	shapes   []shape
	problems []problem
}

var newLine = []byte("\r\n")
var doubleNewLine = []byte("\r\n\r\n")

func getShapes(shapesBytes [][]byte) []shape {
	shapes := make([]shape, len(shapesBytes))

	for shapeI, shapeBytes := range shapesBytes {
		spl := bytes.Split(shapeBytes, newLine)

		//Ignore the first line because it only contains "X:"
		spl = spl[1:]
		shape := make(shape, len(spl))
		for i := range spl {
			shape[i] = make([]bool, len(spl[i]))
			for j := range spl[i] {
				switch spl[i][j] {
				case '#':
					shape[i][j] = true
				case '.':
					shape[i][j] = false
				default:
					panic("Found not valid character")
				}
			}
		}
		shapes[shapeI] = shape
	}

	return shapes
}

func getProblems(problemBytes []byte) []problem {
	spl := bytes.Split(problemBytes, newLine)
	problems := make([]problem, 0, len(spl))

	for _, problemLine := range spl {
		if len(problemLine) == 0 {
			continue
		}
		line := string(problemLine)

		var problem problem
		lineSpl := strings.SplitN(line, ":", 2)
		dimensions := strings.SplitN(lineSpl[0], "x", 2)
		width, err := strconv.Atoi(dimensions[0])

		if err != nil {
			panic(fmt.Errorf("not a valid width: %s", dimensions[0]))
		}
		problem.width = width

		height, err := strconv.Atoi(dimensions[1])
		if err != nil {
			panic(fmt.Errorf("not a valid height: %s", dimensions[0]))
		}
		problem.height = height

		intStrs := strings.Fields(lineSpl[1])
		problem.shapes = make([]int, len(intStrs))
		for i, intStr := range intStrs {
			n, err := strconv.Atoi(intStr)
			if err != nil {
				panic("Invalid int")
			}
			problem.shapes[i] = n
		}

		problems = append(problems, problem)
	}

	return problems
}

func getInput() input {
	inp, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	spl := bytes.Split(inp, doubleNewLine)

	shapesBytes := spl[:len(spl)-1]
	problemsBytes := spl[len(spl)-1]

	input := input{shapes: getShapes(shapesBytes), problems: getProblems(problemsBytes)}

	return input
}

func rotateShape(s shape) shape {
	height := len(s)
	width := len(s[0])

	rotatedS := make(shape, width)
	for i := range s {
		rotatedS[i] = make([]bool, height)
	}

	for originalRow := range height {
		rotatedCol := height - originalRow - 1
		for originalCol := range width {
			rotatedRow := originalCol
			rotatedS[rotatedRow][rotatedCol] = s[originalRow][originalCol]
		}
	}

	return rotatedS
}

func flipShape(s shape) shape {
	height := len(s)
	width := len(s[0])

	flippedS := make(shape, width)
	for i := range s {
		flippedS[i] = make([]bool, height)
	}

	for originalRow := range height {
		flippedRow := height - originalRow - 1
		copy(flippedS[flippedRow], s[originalRow])
	}

	return flippedS
}

func shapesEqual(s1, s2 shape) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	for row := range s1 {
		if len(s1[row]) != len(s2[row]) {
			return false
		}

		for col := range s1 {
			if s1[row][col] != s2[row][col] {
				return false
			}
		}
	}

	return true
}

func containsShape(ss []shape, s shape) bool {
	for _, s2 := range ss {
		if shapesEqual(s, s2) {
			return true
		}
	}
	return false
}

func allVariations(s shape) []shape {
	var uniqueVariations []shape

	currentVariation := s
	for range 2 {
		for range 4 {
			if !containsShape(uniqueVariations, currentVariation) {
				uniqueVariations = append(uniqueVariations, currentVariation)
			}
			currentVariation = rotateShape(currentVariation)
		}
		currentVariation = flipShape(currentVariation)
	}

	return uniqueVariations
}

func bruteForceRecursive(region [][]bool, shapesToUse []int, variations [][]shape, x, y int) bool {
	if x == len(region[0])-2 {
		x = 0
		y++
	}
	if y == len(region)-2 {
		for _, n := range shapesToUse {
			if n != 0 {
				return false
			}
		}
		return true
	}

	for i, n := range shapesToUse {
		if n != 0 {
			for _, variation := range variations[i] {
				variationFits := true
				for r := range variation {
					for c := range variation[r] {
						if variation[r][c] && region[y+r][x+c] {
							variationFits = false
						}
					}
				}

				if !variationFits {
					continue
				}

				shapesToUse[i]--
				for r := range variation {
					for c := range variation[r] {
						if variation[r][c] {
							region[y+r][x+c] = true
						}
					}
				}

				if bruteForceRecursive(region, shapesToUse, variations, x+1, y) {
					return true
				}

				shapesToUse[i]++
				for r := range variation {
					for c := range variation[r] {
						if variation[r][c] {
							region[y+r][x+c] = false
						}
					}
				}
			}
		}
	}

	return bruteForceRecursive(region, shapesToUse, variations, x+1, y)
}

func canSolve(problem problem, shapes []shape) bool {
	totalArea := problem.height * problem.width

	numShapes := 0
	totalShapeArea := 0
	for i, s := range shapes {
		area := 0
		for r := range s {
			for c := range s[r] {
				if s[r][c] {
					area++
				}
			}
		}
		totalShapeArea += area * problem.shapes[i]

		numShapes += problem.shapes[i]
	}

	//Early return if all shapes can't possibly fit
	if totalArea < totalShapeArea {
		return false
	}

	//Early return if all shapes fit without squeezing
	independentShapeSpots := (problem.height / 3) * (problem.width / 3)
	if independentShapeSpots >= numShapes {
		return true
	}

	variations := make([][]shape, len(shapes))
	for i, shape := range shapes {
		variations[i] = allVariations(shape)
	}

	region := make([][]bool, problem.height)
	for i := range region {
		region[i] = make([]bool, problem.width)
	}

	// Brute force approach
	// Note: this is not needed for the input data, as every problem is handled by one of the early returns
	return bruteForceRecursive(region, problem.shapes, variations, 0, 0)
}

func part1(input input) int {
	sum := 0

	for _, problem := range input.problems {
		if canSolve(problem, input.shapes) {
			sum += 1
		}
	}

	return sum
}

func main() {
	input := getInput()
	fmt.Println("part 1:", part1(input))
}
