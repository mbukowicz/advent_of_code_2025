package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	f, err := os.Open("inputs/day6.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}

	totalSum := 0
	nums := make([]int, 0)
	for col := len(lines[0]) - 1; col >= 0; col-- {
		num := 0
		for row := 0; row < len(lines)-1; row++ {
			char := lines[row][col]
			if char != ' ' {
				digit := int(char - '0')
				num = num*10 + digit
			}
		}
		nums = append(nums, num)

		char := lines[len(lines)-1][col]
		switch char {
		case '+':
			sum := 0
			for _, n := range nums {
				sum += n
			}
			fmt.Println("Adding:", nums, "=", sum)
			totalSum += sum
			nums = make([]int, 0)
			col--
		case '*':
			prod := 1
			for _, n := range nums {
				prod *= n
			}
			fmt.Println("Multiplying:", nums, "=", prod)
			totalSum += prod
			nums = make([]int, 0)
			col--
		}
	}
	fmt.Println("Total sum:", totalSum)
}
