package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var total int
	for _, line := range lines {
		var lineDirection string
		levels := strings.Split(line, " ")

		for i, level := range levels {
			if i == len(levels)-1 {
				total += 1
				break
			}
			// gap same or geater than 3 FAIL
			curr, _ := strconv.Atoi(level)
			next, _ := strconv.Atoi(levels[i+1])
			gap := abs(curr - next)
			if gap == 0 || gap > 3 {
				break
			}

			// no direction yet SET
			isDesc := curr > next
			if lineDirection == "" {
				if isDesc == true {
					lineDirection = "desc"
					continue
				} else {
					lineDirection = "asc"
					continue
				}
			}

			// wrong direction FAIL
			if (lineDirection == "desc" && isDesc) || (lineDirection == "asc" && !isDesc) {
				continue
			} else {
				break
			}
		}
	}
	fmt.Println(total)
}

func part2() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var total int
	for _, line := range lines {
		var lineDirection string
		var badLevelCount int
		levels := strings.Split(line, " ")

		for i, level := range levels {
			if i == len(levels)-1 {
				total++
				break
			}
			// gap same or geater than 3 FAIL
			curr, _ := strconv.Atoi(level)
			next, _ := strconv.Atoi(levels[i+1])
			gap := abs(curr - next)
			if gap == 0 || gap > 3 {
				badLevelCount++
				if badLevelCount > 1 {
					break
				}
				continue
			}

			// no direction yet SET
			isDesc := curr > next
			if lineDirection == "" {
				if isDesc == true {
					lineDirection = "desc"
					continue
				} else {
					lineDirection = "asc"
					continue
				}
			}

			// wrong direction FAIL
			if (lineDirection == "desc" && isDesc) || (lineDirection == "asc" && !isDesc) {
				continue
			} else {
				badLevelCount++
				if badLevelCount > 1 {
					break
				}
				continue
			}
		}
	}
	fmt.Println(total)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
