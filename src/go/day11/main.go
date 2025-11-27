package main

import (
	"advent2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := utils.ReadFile("../data/day11.txt")
	stones := strings.Fields(data)

	partA(stones)
	partB(stones)
}

func trimLeadingZeros(s string) string {
	if trimmedString := strings.TrimLeft(s, "0"); trimmedString != "" {
		return trimmedString
	}
	return "0"
}

func partA(stones []string) {
	for range 25 {
		orderedStones := make([]string, 0, len(stones)*2)

		for _, stone := range stones {
			// Convert stone string to an int.
			intStone := utils.ToInt(stone)
			switch {
			// Val is 0.
			case intStone == 0:
				// Value becomes 1.
				orderedStones = append(orderedStones, "1")
			// Stone length is even, split in half.
			case len(stone)%2 == 0:
				// Trim leading zeros and append both halfs to the ordered stones.
				orderedStones = append(
					orderedStones,
					trimLeadingZeros(stone[:len(stone)/2]),
					trimLeadingZeros(stone[len(stone)/2:]),
				)
			// Stone length is odd, multiply by 2024 and append to ordered stones.
			default:
				orderedStones = append(orderedStones, strconv.Itoa(intStone*2024))
			}
		}

		// Make stones new ordered stones for next iteration.
		stones = orderedStones
	}

	fmt.Println("Part A result:", len(stones))
}

func partB(initialStones []string) {
	// Count how many times each stone value occurs.
	stoneCounts := make(map[string]int)
	for _, stone := range initialStones {
		stoneCounts[stone]++
	}

	for range 75 {
		newCounts := make(map[string]int)

		for stone, count := range stoneCounts {
			// Convert stone string to an int.
			intStone := utils.ToInt(stone)

			switch {
			// Val is 0.
			case intStone == 0:
				// Value becomes 1.
				newCounts["1"] += count
			// Stone length is even, split in half.
			case len(stone)%2 == 0:
				// Trim leading zeros for both halves and add them to the new counts.
				mid := len(stone) / 2
				newCounts[trimLeadingZeros(stone[:mid])] += count
				newCounts[trimLeadingZeros(stone[mid:])] += count
			// Stone length is odd, multiply by 2024.
			default:
				newCounts[strconv.Itoa(intStone*2024)] += count
			}
		}

		// Update stone counts for next iteration.
		stoneCounts = newCounts
	}

	// Sum total number of stones.
	total := 0
	for _, count := range stoneCounts {
		total += count
	}

	fmt.Println("Part B result:", total)
}
