package main

import (
	"advent2024/utils"
	"fmt"
	"sort"
)

type Memory struct {
	id     int
	length int
	isFile bool
}

func main() {
	diskMap := utils.ReadFile("../data/day9.txt")

	// Convert string to slice of digits
	diskMapDigits := make([]int, len(diskMap))
	for i, r := range diskMap {
		diskMapDigits[i] = int(r - '0')
	}

	// Flatten the diskMap to a slice of files or free space.
	var flattenedMap []int
	currentFileID := 0
	for i, memory := range diskMapDigits {
		if i%2 == 0 {
			for range memory {
				flattenedMap = append(flattenedMap, currentFileID)
			}
			currentFileID++
		} else {
			for range memory {
				flattenedMap = append(flattenedMap, -1)
			}
		}
	}

	// Pass a copy to partA
	partAMap := make([]int, len(flattenedMap))
	copy(partAMap, flattenedMap)
	partA(partAMap)

	// Pass a copy to partB
	partBMap := make([]int, len(flattenedMap))
	copy(partBMap, flattenedMap)
	partB(partBMap)
}

func partA(flattenedMap []int) {
	startIndex := 0                   // Left pointer.
	endIndex := len(flattenedMap) - 1 // Right pointer.

	for startIndex < endIndex { // Continue while the two pointer have not met.
		// Move the start index right until we find a empty -1 gap.
		for startIndex < len(flattenedMap) && flattenedMap[startIndex] != -1 {
			startIndex++
		}
		// Move the end index left until we find a non-zero gap (file).
		for endIndex >= 0 && flattenedMap[endIndex] == -1 {
			endIndex--
		}
		// Ensure the index have not overlapped.
		if startIndex >= endIndex {
			break
		}
		// Swap the right file into the left free space.
		flattenedMap[startIndex], flattenedMap[endIndex] = flattenedMap[endIndex], flattenedMap[startIndex]
	}

	// For each memory slot which contains a file add index * file id to the total.
	total := 0
	for i, value := range flattenedMap {
		if value != -1 {
			total += i * value
		}
	}

	fmt.Println("Part A result:", total)
}

type File struct {
	id     int
	start  int
	length int
	moved  bool
}

func partB(flattenedMap []int) {
	// Build a slice of files
	var files []File
	for i := 0; i < len(flattenedMap); {
		if flattenedMap[i] != -1 {
			id := flattenedMap[i]
			start := i
			for i < len(flattenedMap) && flattenedMap[i] == id {
				i++
			}
			files = append(files, File{id: id, start: start, length: i - start})
		} else {
			i++
		}
	}

	// Sort files by descending ID so we move largest file first
	sort.Slice(files, func(i, j int) bool {
		return files[i].id > files[j].id
	})

	for _, f := range files {
		bestStart := -1
		freeCount := 0

		// Find leftmost free span large enough
		for i := 0; i < f.start; i++ {
			if flattenedMap[i] == -1 {
				freeCount++
				if freeCount >= f.length {
					bestStart = i - f.length + 1
					break
				}
			} else {
				freeCount = 0
			}
		}

		// Move the file if a span is found
		if bestStart != -1 {
			// Clear old position
			for i := f.start; i < f.start+f.length; i++ {
				flattenedMap[i] = -1
			}
			// Place file in new position
			for i := bestStart; i < bestStart+f.length; i++ {
				flattenedMap[i] = f.id
			}
		}
	}

	// Compute checksum
	total := 0
	for i, value := range flattenedMap {
		if value != -1 {
			total += i * value
		}
	}

	fmt.Println("Part B result:", total)
}
