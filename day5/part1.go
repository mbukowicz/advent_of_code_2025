package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Range struct {
	Start int
	End   int
}

func Part1() {
	f, err := os.Open("inputs/day5.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var ranges []Range
	readRanges := true
	freshCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("ranges", ranges)
			readRanges = false
			continue
		}
		if readRanges {
			var r Range
			_, err := fmt.Sscanf(line, "%d-%d", &r.Start, &r.End)
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, r)
		} else {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			isFresh := false
			for _, r := range ranges {
				if n >= r.Start && n <= r.End {
					isFresh = true
					break
				}
			}
			if isFresh {
				freshCount++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
	fmt.Println("Total fresh:", freshCount)
}
