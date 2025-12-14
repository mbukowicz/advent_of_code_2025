package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	f, err := os.Open("inputs/day3.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	total_sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]int, 0, len(line))
		for _, r := range line {
			n := int(r - '0')
			digits = append(digits, n)
		}

		largestJoltage := findLargestJoltage(digits)
		fmt.Println("Largest joltage for line", line, "is", largestJoltage)

		total_sum += largestJoltage
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
	fmt.Println("Total sum:", total_sum)
}

func findLargestJoltage(digits []int) int {
	largest := 0
	k := 12
	start := 0
	for k > 0 {
		maxIdx := start
		for i := start + 1; i <= len(digits)-k; i++ {
			if digits[i] > digits[maxIdx] {
				maxIdx = i
			}
		}
		largest = largest*10 + digits[maxIdx]
		start = maxIdx + 1
		k--
	}
	return largest
}
