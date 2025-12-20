package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.Open("inputs/day6.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numsRows := [][]int{}
	operations := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if fields[0] != "+" && fields[0] != "*" {
			nums := []int{}
			for _, strNum := range fields {
				num, err := strconv.Atoi(strNum)
				if err != nil {
					fmt.Fprintf(os.Stderr, "invalid number %q: %v\n", strNum, err)
					return
				}
				nums = append(nums, num)
			}
			numsRows = append(numsRows, nums)
		} else {
			operations = append(operations, fields...)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}

	// fmt.Println("Numbers:", numsRows)
	// fmt.Println("Operations:", operations)
	totalSum := 0
	for i, op := range operations {
		result := 0
		switch op {
		case "+":
			for _, nums := range numsRows {
				result += nums[i]
			}
		case "*":
			result = 1
			for _, nums := range numsRows {
				result *= nums[i]
			}
		default:
			fmt.Fprintf(os.Stderr, "unknown operation %q\n", op)
			return
		}
		totalSum += result
		fmt.Println(result)
	}
	fmt.Println("Total sum:", totalSum)
}
