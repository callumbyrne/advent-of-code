package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// part1()
	part2()
}

type Coord struct {
	x int
	y int
}

type Point struct {
	c Coord
	v string
	d Coord
}

func part1() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	maxWidth := len(lines[0])
	maxHeight := len(lines)

	var total int
	for y := 0; y < maxHeight; y++ {
		for x := 0; x < maxWidth; x++ {
			currentPoint := Point{
				c: Coord{
					x: x,
					y: y,
				},
				v: string(lines[y][x]),
			}

			if currentPoint.v == "X" {
				neighbours := getNeighbours(currentPoint, lines, maxHeight, maxWidth)
				for _, n := range neighbours {
					if n.v == "M" {
						n1 := getNextValid(n, lines, maxHeight, maxWidth)
						if n1.v == "A" {
							n2 := getNextValid(n1, lines, maxHeight, maxWidth)
							if n2.v == "S" {
								total++
								continue
							}
						}
						continue
					}
				}
			}
		}
	}
	fmt.Println(total)
}

func part2() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	maxWidth := len(lines[0])
	maxHeight := len(lines)

	var total int
	for y := 0; y < maxHeight; y++ {
		for x := 0; x < maxWidth; x++ {
			currentPoint := Point{
				c: Coord{
					x: x,
					y: y,
				},
				v: string(lines[y][x]),
			}

			if currentPoint.v == "A" {
				neighbours := getNeighbours(currentPoint, lines, maxHeight, maxWidth)
				r := checkValid(neighbours)
				if r == true {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func getNeighbours(point Point, lines []string, maxHeight int, maxWidth int) []Point {
	var neighbours []Point

	for ay := -1; ay <= 1; ay++ {
		if point.c.y+ay == -1 || point.c.y+ay == maxHeight {
			continue
		}
		for ax := -1; ax <= 1; ax++ {
			if point.c.x+ax == -1 || point.c.x+ax == maxWidth {
				continue
			}

			neighbour := Point{
				c: Coord{
					x: point.c.x + ax,
					y: point.c.y + ay,
				},
				v: string(lines[point.c.y+ay][point.c.x+ax]),
				d: Coord{
					x: ax,
					y: ay,
				},
			}

			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}

func getNextValid(point Point, lines []string, maxHeight int, maxWidth int) Point {
	if point.c.y+point.d.y == -1 || point.c.y+point.d.y == maxHeight {
		return Point{}
	} else if point.c.x+point.d.x == -1 || point.c.x+point.d.x == maxWidth {
		return Point{}
	}

	next := Point{
		c: Coord{
			x: point.c.x + point.d.x,
			y: point.c.y + point.d.y,
		},
		v: string(lines[point.c.y+point.d.y][point.c.x+point.d.x]),
		d: point.d,
	}

	return next
}

func checkValid(neighbours []Point) bool {
	if len(neighbours) != 9 {
		return false
	}

	var rl bool
	var lr bool

	if neighbours[0].v == "M" && neighbours[8].v == "S" {
		rl = true
	} else if neighbours[0].v == "S" && neighbours[8].v == "M" {
		rl = true
	}

	if neighbours[2].v == "M" && neighbours[6].v == "S" {
		lr = true
	} else if neighbours[2].v == "S" && neighbours[6].v == "M" {
		lr = true
	}

	return rl == true && lr == true
}
