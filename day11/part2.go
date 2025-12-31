package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part2() {
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
	countFromSvrToOutThroughFftAndDac := countConnections("svr", "fft", connections, cache) * countConnections("fft", "dac", connections, cache) * countConnections("dac", "out", connections, cache)
	countFromSvrToOutThroughDacAndFft := countConnections("svr", "dac", connections, cache) * countConnections("dac", "fft", connections, cache) * countConnections("fft", "out", connections, cache)
	totalCount := countFromSvrToOutThroughFftAndDac + countFromSvrToOutThroughDacAndFft
	fmt.Println(totalCount)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
