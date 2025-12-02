package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type idRange struct {
	start int
	end   int
}

func getRanges() []idRange {
	var reader = bufio.NewReader(os.Stdin)
	var ranges []idRange
	for {
		str, err := reader.ReadString(',')

		eofReached := false
		if err == io.EOF {
			eofReached = true
		} else if err != io.EOF && err != nil {
			panic(err)
		}

		var idRange idRange
		num, err := fmt.Sscanf(str, "%d-%d", &idRange.start, &idRange.end)
		if num != 2 {
			panic("Could not scan start and end of range")
		}
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, idRange)

		if eofReached {
			break
		}
	}

	return ranges
}

func isInvalid(id int) bool {
	idStr := strconv.Itoa(id)
	oddLength := len(idStr)%2 == 1
	if oddLength {
		return false
	}

	middle := len(idStr) / 2
	firstHalf := idStr[:middle]
	secondHalf := idStr[middle:]
	return firstHalf == secondHalf
}

func part1(ranges []idRange) int {
	result := 0
	for _, idRange := range ranges {
		diff := idRange.end - idRange.start + 1
		for i := range diff {
			id := idRange.start + i
			if isInvalid(id) {
				result += id
			}
		}
	}
	return result
}

func part2(ranges []idRange) int {

	return 0
}

func main() {
	ranges := getRanges()
	fmt.Println("part 1:", part1(ranges))
	fmt.Println("part 2:", part2(ranges))
}
