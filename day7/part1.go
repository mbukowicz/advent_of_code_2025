package day7

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
	f, err := os.Open("inputs/day7.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}

	beams := make(map[int]bool)
	beamCount := 0
	for row, line := range lines {
		if row == 0 {
			for col := 0; col < len(line); col++ {
				if line[col] == 'S' {
					beams[col] = true
				}
			}
		} else {
			newBeams := make(map[int]bool)
			for beamCol, _ := range beams {
				switch line[beamCol] {
				case '.':
					newBeams[beamCol] = true
				case '^':
					beamCount++
					newBeams[beamCol-1] = true
					newBeams[beamCol+1] = true
				}
			}
			beams = newBeams
		}
	}
	fmt.Println("Total beams:", beamCount)
}
