package day8

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Part2() {
	f, err := os.Open("inputs/day8.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	points := []Point{}
	pointsNotInCircuits := map[Point]bool{}
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
		pointsNotInCircuits[point] = true
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
	for _, d := range distances {
		fmt.Printf("Considering distance %d between %+v and %+v\n", d.dist, *d.A, *d.B)
		circuitA, circuitAExists := circuits[*d.A]
		circuitB, circuitBExists := circuits[*d.B]

		// no need to merge circuits because we're only interested in when all points get connected

		if circuitAExists && !circuitBExists {
			circuits[*d.B] = circuitA
			fmt.Printf("Adding point %+v to circuit %d\n", *d.B, circuitA)
			delete(pointsNotInCircuits, *d.B)
		} else if !circuitAExists && circuitBExists {
			circuits[*d.A] = circuitB
			fmt.Printf("Adding point %+v to circuit %d\n", *d.A, circuitB)
			delete(pointsNotInCircuits, *d.A)
		} else if !circuitAExists && !circuitBExists {
			circuits[*d.A] = circuitID
			circuits[*d.B] = circuitID
			fmt.Printf("Creating new circuit %d with points %+v and %+v\n", circuitID, *d.A, *d.B)
			circuitID++
			delete(pointsNotInCircuits, *d.A)
			delete(pointsNotInCircuits, *d.B)
		}
		fmt.Println()

		if len(pointsNotInCircuits) == 0 {
			result := d.A.X * d.B.X
			fmt.Printf("All points are now in circuits. Final result: %d\n", result)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
