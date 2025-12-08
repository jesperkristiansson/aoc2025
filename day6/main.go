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

type data struct {
	numbers   [][]int
	operators []operator
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

func getData() data {
	var data data

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		stringSlice := strings.Fields(line)

		if len(stringSlice) == 0 || len(stringSlice[0]) == 0 {
			panic("Expected at least one rune in line")
		}

		r := rune(stringSlice[0][0])
		if unicode.IsDigit(r) {
			if len(data.numbers) == 0 {
				data.numbers = make([][]int, len(stringSlice))
			}
			for i, str := range stringSlice {
				n, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				data.numbers[i] = append(data.numbers[i], n)
			}
		} else {
			data.operators = make([]operator, len(stringSlice))
			for i, str := range stringSlice {
				op, err := toOperator(str)
				if err != nil {
					panic(err)
				}
				data.operators[i] = op
			}

			break
		}
	}

	return data
}

func doOperation(numbers []int, op operator) int {
	switch op {
	case operatorAdd:
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		return sum
	case operatorMul:
		product := 1
		for _, n := range numbers {
			product *= n
		}
		return product
	default:
		panic(fmt.Errorf("operator not recognizable '%d'", op))
	}
}

func part1(data data) int {
	sum := 0
	for i := range data.operators {
		sum += doOperation(data.numbers[i], data.operators[i])
	}
	return sum
}

func part2(data data) int {
	return 0
}

func main() {
	data := getData()
	fmt.Println(data.numbers, data.operators)
	fmt.Println("part 1:", part1(data))
	fmt.Println("part 2:", part2(data))
}
