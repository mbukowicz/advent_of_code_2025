package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	cache := map[string]int{}
	fmt.Println(countConnections("you", "out", connections, cache))

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
