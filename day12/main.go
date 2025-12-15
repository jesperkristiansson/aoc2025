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

func part1(input input) int {
	sum := 0

	return sum
}

func part2(input input) int {
	sum := 0

	return sum
}

func main() {
	input := getInput()
	fmt.Println(input)
	fmt.Println("part 1:", part1(input))
	fmt.Println("part 2:", part2(input))
}
