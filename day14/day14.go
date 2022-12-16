package day14

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ROCK    string = "#"
	AIR     string = "."
	SOURCE  string = "+"
	SAND    string = "o"
	FALLING string = "~"
)

const input = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

type Trace struct {
	ps []Points
}

func (t Trace) MaxDist() (int, int, int, int) {
	minX := 999999999999
	minY := 999999999999
	var maxX, maxY int
	for _, v := range t.ps {
		if maxX < v.getMaxX() {
			maxX = v.getMaxX()
		}
		if maxY < v.getMaxY() {
			maxY = v.getMaxY()
		}
		if minY > v.getMinY() {
			minY = v.getMinY()
		}

		if minX > v.getMinX() {
			minX = v.getMinX()
		}

	}

	return maxX, maxY, minX, minY
}

func (t Points) getMaxX() int {
	max := 0
	for _, point := range t {
		if point.x > max {
			max = point.x
		}
	}
	return max
}

func (t Points) getMaxY() int {
	max := 0
	for _, point := range t {
		if point.y >= max {
			max = point.y
		}
	}

	return max
}

func (t Points) getMinX() int {
	min := 999999999999999999
	for _, point := range t {
		if point.x < min {
			min = point.x
		}
	}

	return min
}

func (t Points) getMinY() int {
	min := 999999999999999999
	for _, point := range t {
		if point.y < min {
			min = point.y
		}
	}
	return min
}

type Point struct {
	x, y int
}

type Points []Point

type Grid struct {
	data         [][]string
	sandLocation Point
}

func (g *Grid) OneDown(c Point) Point {
	return Point{x: c.x - 1, y: c.y}
}

func (g *Grid) print() {
	for _, v := range g.data {
		fmt.Println(v)
	}
}

func (g *Grid) OneUp(c Point) Point {
	return Point{x: c.x + 1, y: c.y}
}

func (g *Grid) dropSand() {
	dropPoint := Point{500, 0}

	g.set(dropPoint, SAND)
	g.set((Point{500, 2}), SAND)
}

func (g *Grid) set(p Point, s string) {
	g.data[p.x][p.y] = s
}

func Maim() {
	trace := getTrace()
	maxX, maxY, _, _ := trace.MaxDist()
	grid := getGrid(maxY+1, maxX+1)

	for i, inst := range grid.data {
		for j := range inst {
			p := Point{j, i}
			grid.set(p, AIR)
		}
	}
	for _, inst := range trace.ps {
		for _, l := range inst {
			grid.set(l, ROCK)
		}
	}
	start := Point{500, 0}
	grid.set(start, "+")
	grid.dropSand()
	grid.print()
}

func getGrid(x, y int) Grid {
	grid := Grid{data: make([][]string, x)}
	for i := range grid.data {
		grid.data[i] = make([]string, y)
	}
	return grid
}

func getTrace() Trace {
	file, _ := os.ReadFile("input14.txt")

	lines := split(string(file), "\n")
	trace := Trace{}
	for _, line := range lines[:182] {
		points := split(line, " -> ")
		pointss := []Point{}
		for _, v := range points {

			k := split(v, ",")
			x := toInt(k[0])
			y := toInt(k[1])
			pointss = append(pointss, Point{x: x, y: y})

		}
		trace.ps = append(trace.ps, pointss)
	}

	return trace
}

func split(s string, perim string) []string {
	return strings.Split(s, perim)
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return v
}
