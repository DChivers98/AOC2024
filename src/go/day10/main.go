package main

import (
	"advent2024/utils"
	"fmt"
)

type Position struct {
	row, col int
}

func main() {
	lines := utils.ReadFileLines("../data/day10.txt")
	var grid [][]int
	for _, line := range lines {
		var nums []int
		for _, num := range line {
			nums = append(nums, int(num-'0'))
		}
		grid = append(grid, nums)
	}

	var startPositions []Position
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				startPositions = append(startPositions, Position{row, col})
			}
		}
	}

	partA(grid, startPositions)
	partB(grid, startPositions)
}

func partA(grid [][]int, startPositions []Position) {
	target := 9
	total := 0

	var walkForward func(row, col, currentNumber int, reached map[Position]bool)
	walkForward = func(row, col, currentNumber int, reached map[Position]bool) {
		// If reached end of path (9) append that 9 location to the reached map for that start position.
		if currentNumber == target {
			reached[Position{row, col}] = true
			return
		}

		// Up, down, left and right directions in row, col format.
		for _, direction := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			// Check each direction.
			nextRow, nextCol := row+direction[0], col+direction[1]

			// If not out of bounds and one of the directions is the next number in the sequence walk forward again.
			if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) && grid[nextRow][nextCol] == currentNumber+1 {
				walkForward(nextRow, nextCol, currentNumber+1, reached)
			}
		}
	}

	for _, startPosition := range startPositions {
		// Map to store reached path ends (9).
		reached := make(map[Position]bool)
		// Start walking forward at each starting position (0).
		walkForward(startPosition.row, startPosition.col, 0, reached)
		// Add the reached path ends for each start position to the total.
		total += len(reached)
	}

	fmt.Println("Part A result:", total)
}

func partB(grid [][]int, startPositions []Position) {
	target := 9
	total := 0

	var walkForward func(row, col, currentNumber int)
	walkForward = func(row, col, currentNumber int) {
		// If reached end of path (9) append that 9 location to the reached map for that start position.
		if currentNumber == target {
			total++
			return
		}

		// Up, down, left and right directions in row, col format.
		for _, direction := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			// Check each direction.
			nextRow, nextCol := row+direction[0], col+direction[1]

			// If not out of bounds and one of the directions is the next number in the sequence walk forward again.
			if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) && grid[nextRow][nextCol] == currentNumber+1 {
				walkForward(nextRow, nextCol, currentNumber+1)
			}
		}
	}

	for _, startPosition := range startPositions {
		// Start walking forward at each starting position (0).
		walkForward(startPosition.row, startPosition.col, 0)
	}

	fmt.Println("Part B result:", total)
}
