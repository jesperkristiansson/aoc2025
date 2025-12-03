package main

import (
	"bufio"
	"fmt"
	"os"
)

type joltage byte

func getBanks() [][]joltage {
	var banks [][]joltage

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()

		joltageLine := make([]joltage, len(str))
		for i, ch := range str {
			joltageLine[i] = joltage(ch - '0')
		}

		banks = append(banks, joltageLine)
	}

	return banks
}

func maxJoltage(bank []joltage, n int) int {
	maxI := make([]int, n)
	for j := range n {
		var startI int
		if j == 0 {
			startI = 0
		} else {
			startI = maxI[j-1] + 1
		}
		maxI[j] = startI
		endI := len(bank) - (n - j)

		for i := startI; i <= endI; i++ {
			if bank[i] > bank[maxI[j]] {
				maxI[j] = i
			}
		}
	}

	totalJoltage := 0
	factor := 1
	for j := n - 1; j >= 0; j-- {
		totalJoltage += factor * int(bank[maxI[j]])
		factor *= 10
	}

	return totalJoltage
}

func part1(banks [][]joltage) int {
	result := 0
	for _, bank := range banks {
		result += maxJoltage(bank, 2)
	}

	return result
}

func part2(banks [][]joltage) int {
	result := 0
	for _, bank := range banks {
		result += maxJoltage(bank, 12)
	}

	return result
}

func main() {
	banks := getBanks()
	fmt.Println("part 1:", part1(banks))
	fmt.Println("part 2:", part2(banks))
}
