package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidID2(id int64) bool {
	idStr := strconv.FormatInt(id, 10)
	for i := 1; i < len(idStr); i++ {
		if len(idStr)%i != 0 {
			continue
		}
		allRepeated := true
		firstSubstring := idStr[0:i]
		for j := i; j+i <= len(idStr); j += i {
			nextSubstring := idStr[j : j+i]
			if nextSubstring != firstSubstring {
				allRepeated = false
				break
			}
		}
		if allRepeated {
			return true
		}
	}
	return false
}

func sumInvalidIDs2(from int64, to int64) int64 {
	sum := int64(0)
	for i := from; i <= to; i++ {
		if isInvalidID2(i) {
			fmt.Println("Invalid: ", i)
			sum += i
		}
	}
	return sum
}

func Part2() {
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
			result += sumInvalidIDs2(start, end)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
	fmt.Println("Result: ", result)
}
