package day10

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/bartolsthoorn/gohighs/highs"
)

type MachineInPart2 struct {
	buttons  [][]int
	joltages []int
}

func parseButtonsInPart2(buttonStrings []string) [][]int {
	buttons := [][]int{}
	for _, buttonString := range buttonStrings {
		indicatorLightStrings := strings.Split(buttonString[1:len(buttonString)-1], ",")
		toggledIndicatorLights := []int{}
		for _, indicatorLightString := range indicatorLightStrings {
			indicatorLight, err := strconv.Atoi(indicatorLightString)
			if err != nil {
				panic(err)
			}
			toggledIndicatorLights = append(toggledIndicatorLights, indicatorLight)
		}
		buttons = append(buttons, toggledIndicatorLights)
	}
	return buttons
}

func parseJoltages(joltagesInputString string) []int {
	joltageStrings := strings.Split(joltagesInputString[1:len(joltagesInputString)-1], ",")
	joltages := []int{}
	for _, joltageString := range joltageStrings {
		joltage, err := strconv.Atoi(joltageString)
		if err != nil {
			panic(err)
		}
		joltages = append(joltages, joltage)
	}
	return joltages
}

func findMinPressesInPart2(machine MachineInPart2) int {
	fmt.Println("Machine:", machine)

	buttonsCount := len(machine.buttons)

	colCosts := make([]float64, buttonsCount)
	for i := range colCosts {
		colCosts[i] = 1
	}

	colLower := make([]float64, buttonsCount)

	colUpper := make([]float64, buttonsCount)
	for i := range colUpper {
		colUpper[i] = highs.Inf()
	}

	varTypes := make([]highs.VariableType, buttonsCount)
	for i := range varTypes {
		varTypes[i] = highs.Integer
	}

	model := highs.Model{
		ColCosts: colCosts, // Objective coefficients
		ColLower: colLower, // All buttons cannot be negative
		ColUpper: colUpper, // No upper bound
		VarTypes: varTypes, // All integer variables (ILP)
	}

	for joltageIndicatorLight, joltage := range machine.joltages {
		coefficients := make([]float64, buttonsCount)
		for i := range coefficients {
			coefficients[i] = 0
		}
		for i, indicatorLights := range machine.buttons {
			if slices.Contains(indicatorLights, joltageIndicatorLight) {
				coefficients[i] = 1
			}
		}
		model.AddDenseRow(float64(joltage), coefficients, float64(joltage))
	}

	solution, err := model.Solve(highs.WithOutput(false))
	if err != nil {
		panic(err)
	}

	if solution.IsOptimal() {
		return int(solution.Objective)
	} else {
		panic("Could not find an optimal solution")
	}

	return -1
}

func Part2() {
	f, err := os.Open("inputs/day10.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		fmt.Println(parts)
		buttons := parseButtonsInPart2(parts[1 : len(parts)-1])
		joltages := parseJoltages(parts[len(parts)-1])
		machine := MachineInPart2{buttons, joltages}
		minPresses := findMinPressesInPart2(machine)
		fmt.Println("Min presses:", minPresses)
		result += minPresses
	}
	fmt.Println("Result:", result)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}

// Here are examples how to solve the first 2 test cases using gohigh

// import (
// 	"fmt"
// 	"github.com/bartolsthoorn/gohighs/highs"
// )

// Machine #1
// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
// b1 = 1.00, b2 = 2.00, b3 = 0.00, b4 = 4.00, b5 = 0.00, b6 = 3.00
// Objective = 10.00

// model := highs.Model{
// 	ColCosts: []float64{1, 1, 1, 1, 1, 1},                                                                                    // Objective coefficients
// 	ColLower: []float64{0, 0, 0, 0, 0, 0},                                                                                    // All buttons cannot be negative
// 	ColUpper: []float64{highs.Inf(), highs.Inf(), highs.Inf(), highs.Inf(), highs.Inf(), highs.Inf()},                        // No upper bound
// 	VarTypes: []highs.VariableType{highs.Integer, highs.Integer, highs.Integer, highs.Integer, highs.Integer, highs.Integer}, // All integer variables (ILP)
// }

// model.AddDenseRow(3, []float64{0, 0, 0, 0, 1, 1}, 3)
// model.AddDenseRow(5, []float64{0, 1, 0, 0, 0, 1}, 5)
// model.AddDenseRow(4, []float64{0, 0, 1, 1, 1, 0}, 4)
// model.AddDenseRow(7, []float64{1, 1, 0, 1, 0, 0}, 7)

// solution, err := model.Solve(highs.WithOutput(false))
// if err != nil {
// 	panic(err)
// }

// if solution.IsOptimal() {
// 	fmt.Printf("b1 = %.2f, b2 = %.2f, b3 = %.2f, b4 = %.2f, b5 = %.2f, b6 = %.2f\n", solution.ColValues[0], solution.ColValues[1], solution.ColValues[2], solution.ColValues[3], solution.ColValues[4], solution.ColValues[5])
// 	fmt.Printf("Objective = %.2f\n", solution.Objective)
// }

// Machine #2
// [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
// b1 = 2.00, b2 = 5.00, b3 = 0.00, b4 = 5.00, b5 = 0.00
// Objective = 12.00

// model := highs.Model{
// 	ColCosts: []float64{1, 1, 1, 1, 1},                                                                        // Objective coefficients
// 	ColLower: []float64{0, 0, 0, 0, 0},                                                                        // All buttons cannot be negative
// 	ColUpper: []float64{highs.Inf(), highs.Inf(), highs.Inf(), highs.Inf(), highs.Inf()},                      // No upper bound
// 	VarTypes: []highs.VariableType{highs.Integer, highs.Integer, highs.Integer, highs.Integer, highs.Integer}, // All integer variables (ILP)
// }

// model.AddDenseRow(7, []float64{1, 0, 1, 1, 0}, 7)
// model.AddDenseRow(5, []float64{0, 0, 0, 1, 1}, 5)
// model.AddDenseRow(12, []float64{1, 1, 0, 1, 1}, 12)
// model.AddDenseRow(7, []float64{1, 1, 0, 0, 1}, 7)
// model.AddDenseRow(2, []float64{1, 0, 1, 0, 1}, 2)

// solution, err := model.Solve(highs.WithOutput(false))
// if err != nil {
// 	panic(err)
// }

// if solution.IsOptimal() {
// 	fmt.Printf("b1 = %.2f, b2 = %.2f, b3 = %.2f, b4 = %.2f, b5 = %.2f\n", solution.ColValues[0], solution.ColValues[1], solution.ColValues[2], solution.ColValues[3], solution.ColValues[4])
// 	fmt.Printf("Objective = %.2f\n", solution.Objective)
// }
