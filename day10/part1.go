package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MachineInPart1 struct {
	targetIndicatorLights int
	buttons               []int
}

func parseTargetIndicatorLights(indicatorLightsString string) int {
	// example: [.###.#]
	target := 0
	for i := len(indicatorLightsString) - 1; i >= 1; i-- {
		target *= 2
		if indicatorLightsString[i] == '#' {
			target += 1
		}
	}
	return target
}

func pow2(n int) int {
	result := 1
	for range n {
		result *= 2
	}
	return result
}

func toggleIndicatorLights(indicatorLights []int) int {
	result := 0
	for _, indicatorLight := range indicatorLights {
		result ^= indicatorLight
	}
	return result
}

func parseButtonsInPart1(buttonStrings []string) []int {
	buttons := []int{}
	for _, buttonString := range buttonStrings {
		indicatorLightStrings := strings.Split(buttonString[1:len(buttonString)-1], ",")
		toggledIndicatorLights := 0
		for _, indicatorLightString := range indicatorLightStrings {
			indicatorLight, err := strconv.Atoi(indicatorLightString)
			if err != nil {
				panic(err)
			}
			toggledIndicatorLights += pow2(indicatorLight)
		}
		buttons = append(buttons, toggledIndicatorLights)
	}
	return buttons
}

func combinationsWithRepetition[T any](arr []T, k int) [][]T {
	var result [][]T
	var combination []T

	var dfs func(start int)
	dfs = func(start int) {
		if len(combination) == k {
			tmp := make([]T, k)
			copy(tmp, combination)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(arr); i++ {
			combination = append(combination, arr[i])
			dfs(i)                                         // allow picking the same element again
			combination = combination[:len(combination)-1] // remove last element
		}
	}

	dfs(0)
	return result
}

func findMinPressesInPart1(machine MachineInPart1) int {
	fmt.Println("Machine:", machine)

	for k := 1; k < len(machine.buttons); k++ {
		for _, combination := range combinationsWithRepetition(machine.buttons, k) {
			attempt := toggleIndicatorLights(combination)
			fmt.Println("k:", k, "combination:", combination, "attempt:", attempt)
			if attempt == machine.targetIndicatorLights {
				fmt.Println("minPresses:", k)
				return k
			}
		}
	}

	return -1
}

func Part1() {
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
		targetIndicatorLights := parseTargetIndicatorLights(parts[0])
		buttons := parseButtonsInPart1(parts[1 : len(parts)-1])
		machine := MachineInPart1{targetIndicatorLights, buttons}
		result += findMinPressesInPart1(machine)
	}
	fmt.Println("Result:", result)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
