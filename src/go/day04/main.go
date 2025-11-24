package main

import (
	"fmt"

	"advent2024/utils"
)

func main() {
	lines := utils.ReadFileLines("../data/day4.txt")
	var grid [][]string
	for _, line := range lines {
		var chars []string
		for _, char := range line {
			chars = append(chars, string(char))
		}
		grid = append(grid, chars)
	}

	partA(grid)
	partB(grid)
}

func partA(grid [][]string) {
	directions := [][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up Left
		{-1, 1},  // Up Right
		{1, -1},  // Down Left
		{1, 1},   // Down Right
	}

	count := 0
	// Check every position in the grid.
	for currentRowPosition := range len(grid) {
		for currentColPostion := range len(grid[0]) {
			// If the position is an 'X' (the start of XMAS).
			if grid[currentRowPosition][currentColPostion] == "X" {
				// Try all 8 directions to see if they are the next letter.
				for _, direction := range directions {
					directionRow, directionCol := direction[0], direction[1]
					found := true
					// Check each letter of XMAS.
					for stepsInDirection := range 4 {
						// Calculate where the letter should be.
						nextLetterExpectedRow := (currentRowPosition + stepsInDirection*directionRow)
						nextLetterExpectedCol := (currentColPostion + stepsInDirection*directionCol)

						// If out of bounds or wrong letter stop.
						if nextLetterExpectedRow < 0 ||
							nextLetterExpectedRow >= len(grid) ||
							nextLetterExpectedCol < 0 ||
							nextLetterExpectedCol >= len(grid[0]) ||
							grid[nextLetterExpectedRow][nextLetterExpectedCol] != "XMAS"[stepsInDirection:stepsInDirection+1] {
							found = false
							break
						}
					}
					if found {
						count++
					}
				}
			}
		}
	}

	fmt.Printf("Part A result: %d\n", count)
}

func partB(grid [][]string) {
	count := 0
	// Loop the grid.
	for currentRow := 1; currentRow < len(grid)-1; currentRow++ {
		for currentCol := 1; currentCol < len(grid[0])-1; currentCol++ {

			// Looks for A's as that is the center.
			if grid[currentRow][currentCol] != "A" {
				continue
			}

			// Get the 4 corners around the 'A'.
			topLeft := grid[currentRow-1][currentCol-1]
			topRight := grid[currentRow-1][currentCol+1]
			bottomRight := grid[currentRow+1][currentCol+1]
			bottomLeft := grid[currentRow+1][currentCol-1]

			// Check if the other letters around the 'A' are 'M' or 'S'.
			if ((topLeft == "M" && bottomRight == "S") || (topLeft == "S" && bottomRight == "M")) &&
				((topRight == "M" && bottomLeft == "S") || (topRight == "S" && bottomLeft == "M")) {
				count++
			}
		}
	}

	fmt.Printf("Part A result: %d\n", count)
}
