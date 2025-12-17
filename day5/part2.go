package day5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Part2() {
	f, err := os.Open("inputs/day5.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var ranges []Range
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("ranges", ranges)
			break
		}
		var r Range
		_, err := fmt.Sscanf(line, "%d-%d", &r.Start, &r.End)
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, r)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}

	fmt.Println("Original ranges:", ranges)
	mergedRanges := MergeRanges(ranges)
	fmt.Println("Merged ranges:", mergedRanges)

	freshCount := 0
	for _, r := range mergedRanges {
		freshCount += r.End - r.Start + 1
	}
	fmt.Println("Total fresh:", freshCount)
}

func MergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return nil
	}

	// Sort by Start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if r.Start <= last.End { // overlap or touching
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}
