package day7

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
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

	beams := make(map[int]int)
	for row, line := range lines {
		if row == 0 {
			for col := 0; col < len(line); col++ {
				if line[col] == 'S' {
					beams[col] = 1
				}
			}
		} else {
			newBeams := make(map[int]int)
			for beamCol, incomingBeams := range beams {
				switch line[beamCol] {
				case '.':
					outgoingBeams, exists := newBeams[beamCol]
					if exists {
						newBeams[beamCol] = outgoingBeams + incomingBeams
					} else {
						newBeams[beamCol] = incomingBeams
					}
				case '^':
					outgoingBeamsLeft, existsLeft := newBeams[beamCol-1]
					if existsLeft {
						newBeams[beamCol-1] = outgoingBeamsLeft + incomingBeams
					} else {
						newBeams[beamCol-1] = incomingBeams
					}
					outgoingBeamsRight, existsRight := newBeams[beamCol+1]
					if existsRight {
						newBeams[beamCol+1] = outgoingBeamsRight + incomingBeams
					} else {
						newBeams[beamCol+1] = incomingBeams
					}
				}
			}
			beams = newBeams
		}
	}
	beamCount := 0
	for _, beams := range beams {
		beamCount += beams
	}
	fmt.Println("Total beams:", beamCount)
}
