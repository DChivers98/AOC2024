package main

import (
	"advent2024/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFileLines("../data/day12.txt")

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

// Location is a position in the grid (garden).
type Location struct {
	row, col int
}

func partA(garden [][]string) {
	// Track visited locations to avoid reprocessing.
	visited := make(map[Location]bool)
	totalPrice := 0

	// Recursive to find region and calculate area and perimeter.
	var findRegion func(int, int, string) (int, int)
	findRegion = func(row, col int, plant string) (int, int) {
		// If already visited location, skip.
		if visited[Location{row, col}] {
			return 0, 0
		}

		// Mark location as visited.
		visited[Location{row, col}] = true
		plotArea, plotPerimeter := 1, 0

		// Check all 4 directions (up, down, left, right)
		for _, direction := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextRow, nextCol := row+direction[0], col+direction[1]

			// If neighbor is out of bounds or a different plant type, add to the perimeter.
			if nextRow < 0 || nextRow >= len(garden) || nextCol < 0 || nextCol >= len(garden[0]) || garden[nextRow][nextCol] != plant {
				plotPerimeter++
			} else {
				// Recursively process the neighbors of the same plant type.
				area, perimeter := findRegion(nextRow, nextCol, plant)
				plotArea += area
				plotPerimeter += perimeter
			}
		}

		return plotArea, plotPerimeter
	}

	// Iterate thorugh each cell in the garden.
	for row := range len(garden) {
		for col := range len(garden[0]) {
			// If the cell hasn't been visited, process the entire region.
			if !visited[Location{row, col}] {
				area, perimeter := findRegion(row, col, garden[row][col])
				// Add the price to the total.
				totalPrice += area * perimeter
			}
		}
	}

	fmt.Println("Part A result:", totalPrice)
}

func partB(garden [][]string) {
	// Track visited locations to avoid reprocessing.
	visited := make(map[Location]bool)
	totalPrice := 0

	// Recursive to find region.
	var findRegion func(int, int, string, map[Location]bool)
	findRegion = func(row, col int, plant string, region map[Location]bool) {
		// If out of bounds, already visited, or a different plant type exit.
		if row < 0 || row >= len(garden) || col < 0 || col >= len(garden[0]) || visited[Location{row, col}] || garden[row][col] != plant {
			return
		}

		// Mark location as visited and add to current region.
		visited[Location{row, col}] = true
		region[Location{row, col}] = true

		// Recursively process all 4 adjacent cells of the same plant type.
		for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			findRegion(row+d[0], col+d[1], plant, region)
		}
	}

	// Iterate through each cell in the garden.
	for row := range len(garden) {
		for col := range len(garden[0]) {
			// If the cell hasn't been visited, process its entire region.
			if !visited[Location{row, col}] {
				region := make(map[Location]bool)
				findRegion(row, col, garden[row][col], region)

				// Calculate the area.
				area := len(region)
				sides := 0

				// For each cell in the region, check its 4 corners.
				for loc := range region {
					// Check all 4 corners: top-left, top-right, bottom-right, bottom-left.
					for _, corner := range [][2]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}} {
						// Check if neighbors exist in the region.
						verticalNeighbor := region[Location{loc.row + corner[0], loc.col}]
						horizontalNeighbor := region[Location{loc.row, loc.col + corner[1]}]
						diagonalNeighbor := region[Location{loc.row + corner[0], loc.col + corner[1]}]

						// Count a side if
						// Both vertical and horizontal neighbors are missing (outer corner).
						// OR
						// Both neighbors exist but diagonal doesn't (inner corner).
						if (!verticalNeighbor && !horizontalNeighbor) || (verticalNeighbor && horizontalNeighbor && !diagonalNeighbor) {
							sides++
						}
					}
				}
				// Add the price to the total.
				totalPrice += area * sides
			}
		}
	}

	fmt.Println("Part B result:", totalPrice)
}
