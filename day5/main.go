package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

func getRanges(scanner *bufio.Scanner) []idRange {
	var ranges []idRange

	for scanner.Scan() {
		str := scanner.Text()

		str = strings.TrimSpace(str)

		if str == "" {
			break
		}

		var idRange idRange
		num, err := fmt.Sscanf(str, "%d-%d", &idRange.start, &idRange.end)
		if err != nil {
			panic(err)
		}
		if num != 2 {
			panic("Could not scan start and end of range")
		}

		ranges = append(ranges, idRange)
	}

	return ranges
}

func getIds(scanner *bufio.Scanner) []int {
	var ids []int

	for scanner.Scan() {
		str := scanner.Text()

		id, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	return ids
}

func getInput() ([]idRange, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	return getRanges(scanner), getIds(scanner)
}

func sortRanges(ranges []idRange) {
	sortFunc := func(a, b idRange) int {
		diff1 := a.start - b.start
		if diff1 != 0 {
			return diff1
		}

		diff2 := a.end - b.end
		return diff2
	}
	slices.SortFunc(ranges, sortFunc)
}

func part1(ranges []idRange, ids []int) int {
	result := 0
	for _, id := range ids {
		for _, r := range ranges {
			if r.start > id {
				break
			}

			if r.end >= id {
				result++
				break
			}
		}
	}
	return result
}

func part2(ranges []idRange) int {
	result := 0
	lastFreshId := 0
	for _, r := range ranges {
		start := max(r.start, lastFreshId+1)
		if start <= r.end {
			result += r.end - start + 1
			lastFreshId = r.end
		}

	}
	return result
}

func main() {
	ranges, ids := getInput()
	sortRanges(ranges)
	fmt.Println("part 1:", part1(ranges, ids))
	fmt.Println("part 2:", part2(ranges))
}
