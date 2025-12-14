package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
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

		largestFirstDigit, largestFirstDigitIndex := largestFirstDigitWithIndex(digits)
		if largestFirstDigitIndex == -1 {
			panic("no largest number found")
		}
		largestSecondDigit := largestSecondDigit(digits, largestFirstDigitIndex+1)
		// fmt.Println("Largest first num:", largestFirstNum, "at index", largestFirstNumIndex, "largest second num:", largestSecondNum)

		largestJoltage := largestFirstDigit*10 + largestSecondDigit
		fmt.Println("Largest joltage for line", line, "is", largestJoltage)

		total_sum += largestJoltage
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
	fmt.Println("Total sum:", total_sum)
}

func largestFirstDigitWithIndex(digits []int) (int, int) {
	if len(digits) == 0 {
		return -1, -1
	}
	largest := digits[0]
	largetsIndex := 0
	for j := 1; j < len(digits)-1; j++ {
		if largest < digits[j] {
			largest = digits[j]
			largetsIndex = j
		}
		if largest == 9 {
			break
		}
	}
	return largest, largetsIndex
}

func largestSecondDigit(digits []int, startIndex int) int {
	largest := digits[startIndex]
	for j := startIndex; j < len(digits); j++ {
		if largest < digits[j] {
			largest = digits[j]
		}
	}
	return largest
}
