package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func toOperator(ch byte) (operator, error) {
	switch ch {
	case '+':
		return operatorAdd, nil
	case '*':
		return operatorMul, nil
	default:
		return operator(0), fmt.Errorf("'%s' is not an operator", ch)
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
		if len(str) != 1 {
			panic("Expected operator to only be single character")
		}
		op, err := toOperator(str[0])
		if err != nil {
			panic(err)
		}
		operations[i].op = op
	}

	return operations
}

func getPart2Data(lines []string) []operation {
	var operations []operation

	operatorLine := lines[len(lines)-1]
	numberLines := lines[:len(lines)-1]

	outI := 0
	for i := len(operatorLine) - 1; i >= 0; {
		var operation operation

		for {
			factor := 1
			num := 0
			for j := len(numberLines) - 1; j >= 0; j-- {
				numberLine := numberLines[j]
				ch := numberLine[i]
				if unicode.IsDigit(rune(ch)) {
					num += factor * int(ch-'0')
					factor *= 10
				}
			}

			operation.numbers = append(operation.numbers, num)

			if operatorLine[i] != ' ' {
				op, err := toOperator(operatorLine[i])
				if err != nil {
					panic(err)
				}

				operation.op = op

				outI++
				i -= 2
				break
			} else {
				i -= 1
			}
		}

		operations = append(operations, operation)
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

func doOperations(operations []operation) int {
	sum := 0
	for _, operation := range operations {
		sum += doOperation(operation)
	}
	return sum
}

func main() {
	lines := getLines()
	part1Operations := getPart1Data(lines)
	part2Operations := getPart2Data(lines)
	fmt.Println("part 1:", doOperations(part1Operations))
	fmt.Println("part 2:", doOperations(part2Operations))
}
