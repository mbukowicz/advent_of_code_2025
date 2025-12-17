package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() {
	f, err := os.Open("inputs/day4.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var rows [][]bool
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, 0, len(line))
		for _, ch := range line {
			col := ch == '@'
			row = append(row, col)
		}
		rows = append(rows, row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
	total_rolls := 0
	rolls_removed := true
	for rolls_removed {
		rolls_removed = false
		for row := 0; row < len(rows); row++ {
			for col := 0; col < len(rows[0]); col++ {
				if rows[row][col] {
					adjacentOccupied := countAdjacentOccupied(rows, row, col)
					if adjacentOccupied < 4 {
						rolls_removed = true
						total_rolls++
						rows[row][col] = false
					}
				}
			}
		}
	}
	fmt.Println("Total rolls:", total_rolls)
}
