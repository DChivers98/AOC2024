package main

import (
	"fmt"

	"advent2024/utils"
)

type Position struct {
	row, col int
}

type State struct {
	pos       Position
	direction string
}

func main() {
	lines := utils.ReadFileLines("../data/day6.txt")
	var grid [][]string
	for _, line := range lines {
		var chars []string
		for _, char := range line {
			chars = append(chars, string(char))
		}
		grid = append(grid, chars)
	}

	var start Position
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == "^" {
				start = Position{row, col}
			}
		}
	}

	partB(grid, start)
}

func partA(grid [][]string, start Position) map[Position]bool {
	visitedPositions := make(map[Position]bool)
	row, col := start.row, start.col
	direction := "UP"

	for {
		visitedPositions[Position{row, col}] = true

		nextRow, nextCol := row, col
		switch direction {
		case "UP":
			nextRow--
		case "RIGHT":
			nextCol++
		case "DOWN":
			nextRow++
		case "LEFT":
			nextCol--
		}

		if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
			break
		}

		if grid[nextRow][nextCol] == "#" {
			switch direction {
			case "UP":
				direction = "RIGHT"
			case "RIGHT":
				direction = "DOWN"
			case "DOWN":
				direction = "LEFT"
			case "LEFT":
				direction = "UP"
			}
		} else {
			row, col = nextRow, nextCol
		}
	}

	fmt.Printf("Part A result: %d\n", len(visitedPositions))

	return visitedPositions
}

func partB(grid [][]string, start Position) {
	visitedPositions := partA(grid, start)

	count := 0
	for obstacle := range visitedPositions {
		if obstacle == start {
			continue
		}

		visitedStates := make(map[State]bool)
		row, col := start.row, start.col
		direction := "UP"

		isLoop := false
		for {
			state := State{Position{row, col}, direction}
			if visitedStates[state] {
				isLoop = true
				break
			}

			visitedStates[state] = true
			nextRow, nextCol := row, col
			switch direction {
			case "UP":
				nextRow--
			case "RIGHT":
				nextCol++
			case "DOWN":
				nextRow++
			case "LEFT":
				nextCol--
			}

			// Check if guard leaves the grid (no loop)
			if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
				break
			}

			// Check if there's an obstruction (including our new obstacle)
			if grid[nextRow][nextCol] == "#" || (nextRow == obstacle.row && nextCol == obstacle.col) {
				// Turn right 90 degrees
				switch direction {
				case "UP":
					direction = "RIGHT"
				case "RIGHT":
					direction = "DOWN"
				case "DOWN":
					direction = "LEFT"
				case "LEFT":
					direction = "UP"
				}
			} else {
				// Move forward
				row, col = nextRow, nextCol
			}
		}
		if isLoop {
			count++
		}
	}

	fmt.Printf("Part B result: %d\n", count)
}
