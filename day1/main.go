package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getRotations() []int {
	var scanner = bufio.NewScanner(os.Stdin)
	var rotations []int
	for scanner.Scan() {
		var line = scanner.Text()
		var isRight bool = line[0] == 'R'
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		var rotation int
		if isRight {
			rotation = distance
		} else {
			rotation = -distance
		}
		rotations = append(rotations, rotation)
	}

	return rotations
}

func part1(rotations []int) int {
	var dial = 50
	var count = 0
	for _, rotation := range rotations {
		dial += rotation
		if dial%100 == 0 {
			count++
		}
	}

	return count
}

func part2(rotations []int) int {
	var dial = 50
	var count = 0
	for _, rotation := range rotations {
		dial += rotation
		if dial > 0 {
			count += dial / 100
			dial %= 100
		} else {
			count += -dial/100 + 1
			if dial == rotation {
				count -= 1
			}
			dial = (dial%100 + 100) % 100
		}
	}

	return count
}

func main() {
	rotations := getRotations()
	fmt.Println("part 1:", part1(rotations))
	fmt.Println("part 2:", part2(rotations))
}
