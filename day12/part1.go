package day12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const patternsCount = 6
const linesPerPattern = 5
const patternSize = 3

func canFitNaively(width int, height int, patternCounts []int) bool {
	totalPresents := 0
	for _, count := range patternCounts {
		totalPresents += count
	}

	for row := 0; row < height && totalPresents > 0; row += patternSize {
		for col := 0; col < width && totalPresents > 0; col += patternSize {
			totalPresents--
		}
	}
	return totalPresents == 0
}

func Part1() {
	f, err := os.Open("inputs/day12.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	patterns := [][]int{}
	for i := 0; i < patternsCount; i++ {
		pattern := []int{}
		for j := 0; j < linesPerPattern; j++ {
			if !scanner.Scan() {
				panic("Not enough input while still parsing patterns")
			}
			if j >= 1 && j <= 3 {
				line := scanner.Text()
				for k := 0; k < patternSize; k++ {
					if line[k] == '#' {
						pattern = append(pattern, 1)
					} else {
						pattern = append(pattern, 0)
					}
				}
			}
		}
		patterns = append(patterns, pattern)
	}

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		sizes := strings.Split(parts[0], "x")
		width, _ := strconv.Atoi(sizes[0])
		height, _ := strconv.Atoi(sizes[1])
		patternCounts := []int{}
		for _, patternCountString := range strings.Split(parts[1], " ") {
			patternCount, _ := strconv.Atoi(patternCountString)
			patternCounts = append(patternCounts, patternCount)
		}
		if canFitNaively(width, height, patternCounts) {
			count++
		}
	}
	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
