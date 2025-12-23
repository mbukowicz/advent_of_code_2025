package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countInPart2(from, to string, connections map[string][]string, cache map[string]int) int {
	if from == to {
		return 1
	}

	cacheKey := from + ":" + to
	valueFromCache, existsInCache := cache[cacheKey]
	if existsInCache {
		return valueFromCache
	}

	sum := 0
	for _, device := range connections[from] {
		sum += countInPart2(device, to, connections, cache)
	}

	cache[cacheKey] = sum

	return sum
}

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
	countFromSvrToOutThroughFftAndDac := countInPart2("svr", "fft", connections, cache) * countInPart2("fft", "dac", connections, cache) * countInPart2("dac", "out", connections, cache)
	countFromSvrToOutThroughDacAndFft := countInPart2("svr", "dac", connections, cache) * countInPart2("dac", "fft", connections, cache) * countInPart2("fft", "out", connections, cache)
	totalCount := countFromSvrToOutThroughFftAndDac + countFromSvrToOutThroughDacAndFft
	fmt.Println(totalCount)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
