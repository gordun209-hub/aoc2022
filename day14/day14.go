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

type Sands struct {
	droppinLocation Point
	droppedSands    []Point
	droppedNum      int
}

type Point struct {
	x, y int
}

type Points []Point

type Grid struct {
	data [][]string
}
type Directions struct {
	left, right, up, down, leftDiag, rightDiag string
}

func Maim() {
	trace := getTrace()
	grid := getGrid(&trace)

	sands := Sands{droppinLocation: Point{1, 500}, droppedSands: Points{}, droppedNum: 0}
	fillRocks(&grid, &trace)
	fmt.Println(sands)
}

func getDirections(p Point, g *Grid) Directions {
	left := g.data[p.x][p.y+1]
	right := g.data[p.x][p.y+1]
	up := g.data[p.x][p.y+1]
	down := g.data[p.x][p.y+1]
	leftDiag := g.data[p.x+1][p.y+1]
	rightDiag := g.data[p.x-1][p.y+1]
	return Directions{left: left, right: right, up: up, down: down, leftDiag: leftDiag, rightDiag: rightDiag}
}

// x sabit y azaltinca saga \\//
// x sabit y arttirinca sola \\//
// y sabit x arttirinca assa \\//
// y sabit x azaltinca yukari \\//
// dropping sand from {0,500}
// needs to drop while func comes false
// down one if possible
// attempt to down diagnoal left
// attempt to down diagnoal right

func needToStop(g *Grid, p Point) bool {
	dirs := getDirections(p, g)
	// TODO fix this shiet
	if dirs.leftDiag != AIR && dirs.rightDiag != AIR && dirs.down != AIR {
		return true
	}
	return false
}

func fillRocks(grid *Grid, trace *Trace) {
	for _, inst := range trace.ps {
		grid.fillBetween(inst, "#")
	}
	cunt := 0
	for i := 0; i <= len(grid.data[0])-1; i++ {
		for j := 0; j <= len(grid.data)-1; j++ {
			p := Point{j, i}
			if grid.data[j][i] != "#" {
				grid.set(p, AIR)
			} else {
				cunt++
			}
		}
	}
}

// 2. verdimiz yatay 1. verdimiz dikey aga
func getGrid(trace *Trace) Grid {
	y, x, _, _ := trace.MaxDist()
	grid := Grid{data: make([][]string, y)}

	for i := range grid.data {
		grid.data[i] = make([]string, x)
	}
	return grid
}

func getTrace() Trace {
	file, _ := os.ReadFile("input14.txt")

	lines := split(string(file), "\n")[:182]
	trace := Trace{}
	for _, line := range lines {
		points := split(line, " -> ")
		pointss := []Point{}
		for _, v := range points {

			k := split(v, ",")
			x := toInt(k[0])
			y := toInt(k[1])
			pointss = append(pointss, Point{x: y, y: x})

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

func (g *Grid) set(p Point, s string) {
	g.data[p.x][p.y] = s
}

func (g *Grid) fillBetween(p []Point, s string) {
	// Loop through the points and set the value in the grid
	for i := 0; i < len(p)-1; i++ {
		x1, y1 := p[i].x, p[i].y
		x2, y2 := p[i+1].x, p[i+1].y

		// If the x values are the same, fill the cells between the y values
		if x1 == x2 {
			if y1 <= y2 {
				for y := y1; y <= y2; y++ {
					g.set(Point{x1, y}, s)
				}
			} else {
				for y := y2; y <= y1; y++ {
					g.set(Point{x1, y}, s)
				}
			}
		} else if y1 == y2 {
			// If the y values are the same, fill the cells between the x values
			if x1 <= x2 {
				for x := x1; x <= x2; x++ {
					g.set(Point{x, y1}, s)
				}
			} else {
				for x := x2; x <= x1; x++ {
					g.set(Point{x, y1}, s)
				}
			}
		}
	}
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

	return maxX + 1, maxY + 1, minX, minY
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
