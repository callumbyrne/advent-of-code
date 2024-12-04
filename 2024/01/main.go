package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var l []int
	var r []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "   ")

		num1, err1 := strconv.Atoi(s[0])
		if err1 != nil {
			log.Fatal(err)
		}
		num2, err2 := strconv.Atoi(s[1])
		if err2 != nil {
			log.Fatal(err)
		}
		l = append(l, num1)
		r = append(r, num2)
	}

	sort.Ints(l)
	sort.Ints(r)

	var total int

	for i, n := range l {
		num1 := n
		num2 := r[i]

		if num1 == 0 || num2 == 0 {
			return
		}

		total = total + abs(num1-num2)
	}

	fmt.Println(total)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var l []int
	var r []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "   ")

		num1, err1 := strconv.Atoi(s[0])
		if err1 != nil {
			log.Fatal(err)
		}
		num2, err2 := strconv.Atoi(s[1])
		if err2 != nil {
			log.Fatal(err)
		}
		l = append(l, num1)
		r = append(r, num2)
	}

	var total int

	for _, n1 := range l {
		var o int

		for _, n2 := range r {
			if n2 == n1 {
				o = o + 1
			}
		}

		total = total + (n1 * o)
	}

	fmt.Println(total)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
