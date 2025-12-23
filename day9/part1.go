package day9

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1() {
	f, err := os.Open("inputs/day9.txt")
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
		if len(parts) != 2 {
			panic("invalid point format")
		}
		x, y := 0, 0
		fmt.Sscanf(parts[0], "%d", &x)
		fmt.Sscanf(parts[1], "%d", &y)
		point := Point{Y: y, X: x}
		points = append(points, point)
	}
	maxSize := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			width := abs(points[i].X-points[j].X) + 1
			height := abs(points[i].Y-points[j].Y) + 1
			size := width * height
			if size > maxSize {
				fmt.Println("New max size found:", size, "with points", points[i], "and", points[j])
				maxSize = size
			}
		}
	}

	fmt.Printf("Result: %d\n", maxSize)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
