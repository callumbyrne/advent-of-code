package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// part1()
	part2()
}

func part1() {
	input, _ := os.ReadFile("input.txt")

	pattern := `mul\([0-9]{1,3},[0-9]{1,3}\)`
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return
	}

	muls := r.FindAllString(string(input), -1)

	pattern = `[0-9]{1,3}`
	r, err = regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return
	}

	var total int
	for _, mul := range muls {
		nums := r.FindAllString(mul, 2)
		one, _ := strconv.Atoi(nums[0])
		two, _ := strconv.Atoi(nums[1])

		total = total + (one * two)
	}

	fmt.Println(total)
}

func part2() {
	input, _ := os.ReadFile("input.txt")

	pattern := `mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don\'t\(\)`
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return
	}

	matches := r.FindAllString(string(input), -1)

	active := true
	var total int
	for _, match := range matches {
		if match == "do()" {
			active = true
			continue
		} else if match == "don't()" {
			active = false
			continue
		}

		if active == true {
			pattern = `[0-9]{1,3}`
			r, _ = regexp.Compile(pattern)

			nums := r.FindAllString(match, 2)
			one, _ := strconv.Atoi(nums[0])
			two, _ := strconv.Atoi(nums[1])

			total = total + (one * two)
		}
	}
	fmt.Println(total)
}
