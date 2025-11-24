package main

import (
	"fmt"
	"strings"

	"advent2024/utils"
)

func main() {
	lines := utils.ReadFileLines("../data/day5.txt")

	rules := [][]int{}
	updates := [][]int{}

	section := 0
	for _, line := range lines {
		if line == "" {
			section = 1
			continue
		}

		parts := strings.Split(line, map[int]string{0: "|", 1: ","}[section])

		ints := []int{}
		for _, part := range parts {
			ints = append(ints, utils.ToInt(part))
		}

		if section == 0 {
			rules = append(rules, ints)
		} else {
			updates = append(updates, ints)
		}
	}

	partA(rules, updates)
	partB(rules, updates)
}

func partA(rules [][]int, updates [][]int) {
	total := 0

	for _, update := range updates {
		rightOrder := true
		for i := range len(update) - 1 {
			for j := range len(rules) {
				if update[i] == rules[j][1] && update[i+1] == rules[j][0] {
					rightOrder = false
				}
			}
		}

		if rightOrder {
			total += update[len(update)/2]
		}
	}

	fmt.Printf("Part A result: %d\n", total)
}

func partB(rules [][]int, updates [][]int) {
	total := 0

	for _, update := range updates {
		anySwap := false
		for {
			swapped := false
			for i := 0; i < len(update)-1; i++ {
				for _, rule := range rules {
					if update[i] == rule[1] && update[i+1] == rule[0] {
						update[i], update[i+1] = update[i+1], update[i]
						swapped = true
						anySwap = true
						break
					}
				}
			}
			if !swapped {
				break
			}
		}
		if anySwap {
			total += update[len(update)/2]
		}
	}

	fmt.Printf("Part B result: %d\n", total)
}
