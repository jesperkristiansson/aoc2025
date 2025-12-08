package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operator int

const (
	operatorAdd operator = iota
	operatorMul
)

type operation struct {
	numbers []int
	op      operator
}

func toOperator(str string) (operator, error) {
	switch str {
	case "+":
		return operatorAdd, nil
	case "*":
		return operatorMul, nil
	default:
		return operator(0), fmt.Errorf("'%s' is not an operator", str)
	}
}

func getLines() []string {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func getPart1Data(lines []string) []operation {
	var operations []operation

	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		stringSlice := strings.Fields(line)

		if len(operations) == 0 {
			operations = make([]operation, len(stringSlice))
		}

		for j, str := range stringSlice {
			n, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			operations[j].numbers = append(operations[j].numbers, n)
		}
	}

	operatorTokens := strings.Fields(lines[len(lines)-1])
	for i, str := range operatorTokens {
		op, err := toOperator(str)
		if err != nil {
			panic(err)
		}
		operations[i].op = op
	}

	return operations
}

func doOperation(operation operation) int {
	op := operation.op
	switch op {
	case operatorAdd:
		sum := 0
		for _, n := range operation.numbers {
			sum += n
		}
		return sum
	case operatorMul:
		product := 1
		for _, n := range operation.numbers {
			product *= n
		}
		return product
	default:
		panic(fmt.Errorf("operator not recognizable '%d'", op))
	}
}

func part1(operations []operation) int {
	sum := 0
	for _, operation := range operations {
		sum += doOperation(operation)
	}
	return sum
}

func part2(operations []operation) int {
	return 0
}

func main() {
	lines := getLines()
	part1Operations := getPart1Data(lines)
	fmt.Println("part 1:", part1(part1Operations))
	fmt.Println("part 2:", part2(part1Operations))
}
