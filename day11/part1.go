package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countInPart1(from, to string, connections map[string][]string) int {
	if from == to {
		return 1
	}
	sum := 0
	for _, device := range connections[from] {
		sum += countInPart1(device, to, connections)
	}
	return sum
}

func Part1() {
	f, err := os.Open("inputs/day11.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	connections := map[string][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		from := parts[0]
		for _, to := range strings.Split(parts[1], " ") {
			devices, exists := connections[from]
			if !exists {
				devices = []string{}
			}
			devices = append(devices, to)
			connections[from] = devices
		}
	}
	fmt.Println(connections)
	fmt.Println(countInPart1("you", "out", connections))

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
