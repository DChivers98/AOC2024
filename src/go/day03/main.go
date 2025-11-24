package main

import (
	"fmt"
	"regexp"

	"advent2024/utils"
)

func main() {
	// Read the file int a single string.
	data := utils.ReadFile("../data/day3.txt")

	partA(data)
	partB(data)
}

func partA(data string) {
	// Check for all matches of mul(<number>, <number>) in the string.
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	total := 0
	for _, match := range matches {
		matchMultiply := utils.ToInt(match[1]) * utils.ToInt(match[2])
		total += matchMultiply
	}

	fmt.Printf("Part A result: %d\n", total)
}

func partB(data string) {
	// Check for all matches of mul(<number>, <number>) in the string.
	re := regexp.MustCompile(`(?:do\(\)|don\'t\(\)|mul\((\d+),(\d+)\))`)
	matches := re.FindAllStringSubmatch(data, -1)

	total := 0
	active := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			active = true
		case "don't()":
			active = false
		default:
			if active && len(match) > 2 {
				matchMultiply := utils.ToInt(match[1]) * utils.ToInt(match[2])
				total += matchMultiply
			}
		}
	}

	fmt.Printf("Part B result: %d\n", total)
}
