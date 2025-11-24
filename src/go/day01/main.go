package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2024/utils"
)

func main() {
	lines := utils.ReadFileLines("../data/day1.txt")

	// Slices for left and right of string.
	var left []int
	var right []int

	// Iterate over each line and split into left and right, convert to int.
	for i := range len(lines) {
		result := strings.Split(lines[i], "   ")
		left = append(left, utils.ToInt(result[0]))
		right = append(right, utils.ToInt(result[1]))
	}

	partA(left, right)
	partB(left, right)
}

func partA(left []int, right []int) {
	// Sort the slices.
	slices.Sort(left)
	slices.Sort(right)

	// Iterate the slices and find the absolute difference, then append to total.
	total := 0
	for i := range len(left) {
		difference := utils.AbsVal(left[i] - right[i])
		total += difference
	}

	fmt.Printf("Part A result: %d\n", total)
}

func partB(left []int, right []int) {
	// Map for right totals, key is number, value is total in right list.
	totalsInRight := make(map[int]int)

	// Iterate the right list, if number already added to map add one to total, else add it to map.
	for i := range len(right) {
		currentNum := right[i]
		_, ok := totalsInRight[currentNum]
		if ok {
			totalsInRight[currentNum]++
		} else {
			totalsInRight[currentNum] = 1
		}
	}

	// Iterate left list, find corresponding occurences in right list and multiply. Add to total.
	total := 0
	for i := range len(left) {
		rightTotal, ok := totalsInRight[left[i]]
		if ok {
			total += (rightTotal * left[i])
		} else {
			continue
		}
	}

	fmt.Printf("Part B result: %d\n", total)
}
