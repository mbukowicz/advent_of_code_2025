package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("inputs/day2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	result := int64(0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			bounds := strings.Split(r, "-")
			start, _ := strconv.ParseInt(bounds[0], 10, 64)
			end, _ := strconv.ParseInt(bounds[1], 10, 64)
			result += sumInvalidIDs(start, end)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
	fmt.Println("Result: ", result)
}

func isInvalidID(id int64) bool {
	idStr := strconv.FormatInt(id, 10)
	if len(idStr)%2 == 0 {
		firstHalf := idStr[:len(idStr)/2]
		secondHalf := idStr[len(idStr)/2:]
		if firstHalf == secondHalf {
			return true
		}
	}

	return false
}

func sumInvalidIDs(from int64, to int64) int64 {
	sum := int64(0)
	for i := from; i <= to; i++ {
		if isInvalidID(i) {
			sum += i
		}
	}
	return sum
}
