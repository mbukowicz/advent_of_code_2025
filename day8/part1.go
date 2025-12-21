package day8

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Part1() {
	stopAfter := 1000
	f, err := os.Open("inputs/day8.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	points := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			panic("invalid point format")
		}
		x, y, z := 0, 0, 0
		fmt.Sscanf(parts[0], "%d", &x)
		fmt.Sscanf(parts[1], "%d", &y)
		fmt.Sscanf(parts[2], "%d", &z)
		point := Point{X: x, Y: y, Z: z}
		points = append(points, point)
	}
	distances := []Distance{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := distanceSquared(points[i], points[j])
			d := Distance{A: &points[i], B: &points[j], dist: dist}
			distances = append(distances, d)
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	circuitID := 1
	circuits := map[Point]int{}
	connections := 0
	for _, d := range distances {
		fmt.Printf("Considering distance %d between %+v and %+v\n", d.dist, *d.A, *d.B)
		circuitA, circuitAExists := circuits[*d.A]
		circuitB, circuitBExists := circuits[*d.B]
		if circuitAExists && circuitBExists {
			if circuitA != circuitB {
				// merge B into A
				for p, c := range circuits {
					if c == circuitB {
						circuits[p] = circuitA
					}
				}
				fmt.Printf("Merging circuits %d and %d\n", circuitA, circuitB)
			}
		} else if circuitAExists && !circuitBExists {
			circuits[*d.B] = circuitA
			fmt.Printf("Adding point %+v to circuit %d\n", *d.B, circuitA)
		} else if !circuitAExists && circuitBExists {
			circuits[*d.A] = circuitB
			fmt.Printf("Adding point %+v to circuit %d\n", *d.A, circuitB)
		} else {
			circuits[*d.A] = circuitID
			circuits[*d.B] = circuitID
			fmt.Printf("Creating new circuit %d with points %+v and %+v\n", circuitID, *d.A, *d.B)
			circuitID++
		}
		fmt.Println()
		connections++
		if connections >= stopAfter {
			break
		}
	}

	countsByCircuitId := map[int]int{}
	for _, circuit := range circuits {
		countsByCircuitId[circuit]++
	}

	circuitCounts := []int{}
	for _, count := range countsByCircuitId {
		circuitCounts = append(circuitCounts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(circuitCounts)))

	result := 1
	for i := 0; i < 3 && i < len(circuitCounts); i++ {
		result *= circuitCounts[i]
	}

	fmt.Printf("Result: %d\n", result)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
