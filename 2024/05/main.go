package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	incorrect, rulesMap := part1()
	part2(incorrect, rulesMap)
}

func part1() ([][]string, map[string][]string) {
	input, _ := os.ReadFile("input.txt")
	parts := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rules := strings.Split(strings.TrimSpace(parts[0]), "\n")
	updates := strings.Split(strings.TrimSpace(parts[1]), "\n")

	rulesMap := make(map[string][]string)
	for _, rule := range rules {
		r := strings.Split(rule, "|")
		rulesMap[r[0]] = append(rulesMap[r[0]], r[1])
	}

	var total int
	var incorrect [][]string
outer:
	for _, update := range updates {
		nums := strings.Split(update, ",")
		for i, num := range nums {
			if i == 0 {
				continue
			}
			if val, exists := rulesMap[num]; exists {
				for j := 0; j < i; j++ {
					if slices.Contains(val, nums[j]) {
						incorrect = append(incorrect, nums)
						continue outer
					}
				}
			}
		}
		midIdx := (len(nums) - 1) / 2
		midVal, _ := strconv.Atoi(nums[midIdx])
		total = total + midVal
	}

	// fmt.Println(total)
	return incorrect, rulesMap
}

func part2(incorrect [][]string, rulesMap map[string][]string) {
	var total int
	for _, update := range incorrect {
		fixed := fix(update, rulesMap)
		midIdx := (len(fixed) - 1) / 2
		midVal, _ := strconv.Atoi(fixed[midIdx])
		total = total + midVal
	}
	fmt.Println(total)
}

func fix(nums []string, rulesMap map[string][]string) []string {
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if val, exists := rulesMap[num]; exists {
			for j := 0; j < i; j++ {
				if slices.Contains(val, nums[j]) {
					// remove j from slice
					valueToMove := nums[j]
					updated := append(nums[:j], nums[j+1:]...)
					// nums[j] needs to be moved infront of i
					updated = append(updated[:i], append([]string{valueToMove}, updated[i:]...)...)
					return fix(updated, rulesMap)
				}
			}
		}
	}
	return nums
}
