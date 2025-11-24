package main

import (
	"advent2024/utils"
	"fmt"
)

type Position struct {
	row, col int
}

func main() {
	lines := utils.ReadFileLines("../data/day8.txt")
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
	width := len(grid)
	height := len(grid)
	// Stores the unique antinode positions found.
	antinodes := make(map[Position]bool)

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// Iterate over all positions in the grid.
	for y1 := range height {
		for x1 := range width {
			// Skip empty grid positions.
			if grid[y1][x1] == "." {
				continue
			}
			// Iterate over all positions again th find pairs of antennas.
			for y2 := range height {
				for x2 := range width {
					// Skip empty cells or different frequencies or same cell.
					if grid[y2][x2] == "." || grid[y1][x1] != grid[y2][x2] || (x1 == x2 && y1 == y2) {
						continue
					}

					// Calculate vector difference between two antennas.
					dx, dy := x2-x1, y2-y1

					// Noramlise vector to smallest step by step by dividing by gcd(dx, dy).
					greatestCommonDenominator := gcd(dx, dy)
					dx /= greatestCommonDenominator
					dy /= greatestCommonDenominator

					// Calculate antinode positions one step beyond each antenna along the line.
					a1 := Position{x2 + dx, y2 + dy} // Beyond second antenna
					a2 := Position{x1 - dx, y1 - dy} // Beyond first antenna

					// Check if antinode a1 is inside grid bounds; if yes, add to map
					if a1.row >= 0 && a1.row < width && a1.col >= 0 && a1.col < height {
						antinodes[a1] = true
					}

					// Check if antinode a2 is inside grid bounds; if yes, add to map
					if a2.row >= 0 && a2.row < width && a2.col >= 0 && a2.col < height {
						antinodes[a2] = true
					}
				}
			}
		}
	}
	fmt.Printf("Part A result: %d\n", len(antinodes))
}

func partB(grid [][]string) {
	width := len(grid)
	height := len(grid)
	// Stores the unique antinode positions found.
	antinodes := make(map[Position]bool)

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// Iterate over all positions in the grid.
	for y1 := range height {
		for x1 := range width {
			// Skip empty grid positions.
			if grid[y1][x1] == "." {
				continue
			}
			// Iterate over all positions again th find pairs of antennas.
			for y2 := range height {
				for x2 := range width {
					// Skip empty cells or different frequencies or same cell.
					if grid[y2][x2] == "." || grid[y1][x1] != grid[y2][x2] || (x1 == x2 && y1 == y2) {
						continue
					}

					// Calculate vector difference between two antennas.
					dx, dy := x2-x1, y2-y1

					// Noramlise vector to smallest step by step by dividing by gcd(dx, dy).
					greatestCommonDenominator := gcd(dx, dy)
					dx /= greatestCommonDenominator
					dy /= greatestCommonDenominator

					// From antenna 2, walk in the direction (dx, dy)
					for x, y := x2, y2; x >= 0 && x < width && y >= 0 && y < height; x, y = x+dx, y+dy {
						antinodes[Position{x, y}] = true
					}

					// From antenna 1, walk in the opposite direction (-dx, -dy)
					for x, y := x1, y1; x >= 0 && x < width && y >= 0 && y < height; x, y = x-dx, y-dy {
						antinodes[Position{x, y}] = true
					}
				}
			}
		}
	}
	fmt.Printf("Part B result: %d\n", len(antinodes))
}
