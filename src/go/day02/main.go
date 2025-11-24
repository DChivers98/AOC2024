package main

import (
	"fmt"
	"strings"

	"advent2024/utils"
)

func main() {
	// Read the file.
	lines := utils.ReadFileLines("../data/day2.txt")

	// Read each sting, split into int strings, convert to ints,
	// put int in nested slice for that report.
	intReports := make([][]int, len(lines))
	for i, reportString := range lines {
		numString := strings.Fields(reportString)
		reportNums := make([]int, len(numString))

		for i, reportNum := range numString {
			num := utils.ToInt(reportNum)
			reportNums[i] = num
		}
		intReports[i] = reportNums
	}

	partA(intReports)
	partB(intReports)
}

func safeReport(report []int) bool {
	// Check whether the trend is increasing or decreasing.
	isIncreasing := report[1] > report[0]

	// Iterate through report items, get diff value and abs diff value
	// if the trend is one way at the start and the diff is the other,
	// they contradict so the report is not safe.
	for i := range len(report) - 1 {
		diff := report[i+1] - report[i]
		absDiff := utils.AbsVal(diff)

		if isIncreasing && diff <= 0 ||
			!isIncreasing && diff >= 0 ||
			absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}

func partA(reports [][]int) {
	// Check each report to see if they are safe.
	total := 0
	for _, report := range reports {
		if safeReport(report) {
			total++
		}
	}

	fmt.Printf("Part A result: %d\n", total)
}

func partB(reports [][]int) {
	// Check if report already safe, else go thorugh each reports
	// removing each item and checking if it is safe.
	total := 0
	for _, report := range reports {
		if safeReport(report) {
			total++
			continue
		}

		for i := range len(report) {
			modifiedReport := append(append([]int{}, report[:i]...), report[i+1:]...)

			if safeReport(modifiedReport) {
				total++
				break
			}
		}
	}

	fmt.Printf("Part B result: %d\n", total)
}
