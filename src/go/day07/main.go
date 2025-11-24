package main

import (
	"advent2024/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.ReadFileLines("../data/day7.txt")

	var results []int
	var numbers [][]int
	for _, line := range lines {
		split := strings.Split(line, ":")

		results = append(results, utils.ToInt(split[0]))

		strNumbers := strings.Fields(split[1])

		var intNumbers []int
		for _, number := range strNumbers {
			intNumbers = append(intNumbers, utils.ToInt(number))
		}

		numbers = append(numbers, intNumbers)
	}

	partA(results, numbers)
	partB(results, numbers)
}

func partA(results []int, numbers [][]int) {
	total := 0
	for i := 0; i < len(results) && i < len(numbers); i++ {
		var check func(int, int) bool
		check = func(currentTotal int, idx int) bool {
			return (
			// Check if we are at the end of the numbers and the total is what we are looking for.
			idx == len(numbers[i]) && currentTotal == results[i] ||
				// Check if adding the next number makes the result we are looking for.
				idx < len(numbers[i]) && (check(currentTotal+numbers[i][idx], idx+1) ||
					// Check if multiplying by the next number makes the result we are looking for.
					check(currentTotal*numbers[i][idx], idx+1)))
		}
		if check(numbers[i][0], 1) {
			total += results[i]
		}
	}

	fmt.Printf("Part A result: %d\n", total)
}

func partB(results []int, numbers [][]int) {
	total := 0
	for i := 0; i < len(results) && i < len(numbers); i++ {
		var check func(int, int) bool
		check = func(currentTotal int, idx int) bool {
			return (
			// Check if we are at the end of the numbers and the total is what we are looking for.
			idx == len(numbers[i]) && currentTotal == results[i] ||
				// Check if adding the next number makes the result we are looking for.
				idx < len(numbers[i]) && (check(currentTotal+numbers[i][idx], idx+1) ||
					// Check if multiplying by the next number makes the result we are looking for.
					check(currentTotal*numbers[i][idx], idx+1) ||
					// Check if concatinating the numbers produces the result we are looking for.
					check(utils.ToInt(fmt.Sprintf("%d%d", currentTotal, numbers[i][idx])), idx+1)))
		}
		if check(numbers[i][0], 1) {
			total += results[i]
		}
	}

	fmt.Printf("Part B result: %d\n", total)
}
