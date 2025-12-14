package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1() {
	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	dial := 50
	zeros := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		fmt.Println(direction, clicks) // R 60
		switch direction {
		case "R":
			dial += clicks
		case "L":
			dial -= clicks
		}

		for dial < 0 {
			dial += 100
		}
		for dial >= 100 {
			dial -= 100
		}
		if dial == 0 {
			zeros++
		}
		fmt.Println("Dial is now at:", dial)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
	fmt.Println("Zeros:", zeros)
}
