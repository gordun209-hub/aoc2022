package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TwoDVector = [][]string

type Rope struct {
	TwoDVector             TwoDVector
	StartingPoint          []int
	HeadLocation           []int
	TailLocation           []int
	VisitedNumberByTail    *int
	visitedLocationsByTail [][]int
}

func (r *Rope) clear() {
}

var (
	visitedPoints [][]int
	totalVisited  = 1
)

func (r Rope) Move(direction string, numverOfMove int) {
	for i := 0; i <= numverOfMove-1; i++ {
		switch direction {

		case "U":
			r.clear()
			r.HeadLocation[0] = r.HeadLocation[0] - 1
			r.MoveTail()
		case "D":
			r.clear()
			r.HeadLocation[0] = r.HeadLocation[0] + 1
			r.MoveTail()
		case "L":
			r.clear()
			r.HeadLocation[1] = r.HeadLocation[1] - 1
			r.MoveTail()
		case "R":
			r.clear()
			r.HeadLocation[1] = r.HeadLocation[1] + 1
			r.MoveTail()

		}
	}
}

func (r *Rope) DoesTailNeedToMove() bool {
	if !r.isDiagonalAdj() && !r.isVerticalOrHorizontalAdj() && !r.isOverlappingWithHead() {
		return true
	}
	return false
}

func (r *Rope) isDiagonalAdj() bool {
	// Get the x and y coordinates of the head and tail locations
	headX, headY := r.HeadLocation[0], r.HeadLocation[1]
	tailX, tailY := r.TailLocation[0], r.TailLocation[1]

	// Check if the tail is in any of the four diagonal locations adjacent to the head
	if (headX == tailX+1 && headY == tailY+1) || (headX == tailX+1 && headY == tailY-1) || (headX == tailX-1 && headY == tailY+1) || (headX == tailX-1 && headY == tailY-1) {
		return true
	}

	return false
}

func (r *Rope) isVerticalOrHorizontalAdj() bool {
	// Get the x and y coordinates of the head and tail locations
	headX, headY := r.HeadLocation[0], r.HeadLocation[1]
	tailX, tailY := r.TailLocation[0], r.TailLocation[1]

	// check if horizontally or vertically adjacent
	if (headX == tailX && headY == tailY+1) || (headX == tailX && headY == tailY-1) || (headX == tailX+1 && headY == tailY) || (headX == tailX-1 && headY == tailY) {
		return true
	}

	return false
}

func (r *Rope) MoveTail() {
	// push to visitedPoints if point change, if it is first time visit, increase number
	tailx, taily := r.TailLocation[0], r.TailLocation[1]

	if r.DoesTailNeedToMove() {

		if r.HeadLocation[0] == r.TailLocation[0] {
			if r.HeadLocation[1] > r.TailLocation[1] {
				r.TailLocation[1]++
			} else if r.HeadLocation[1] < r.TailLocation[1] {
				r.TailLocation[1]--
			}
		}
		if r.HeadLocation[1] == r.TailLocation[1] {
			if r.HeadLocation[0] > r.TailLocation[0] {
				r.TailLocation[0]++
			} else if r.HeadLocation[0] < r.TailLocation[0] {
				r.TailLocation[0]--
			}
		}
		if r.HeadLocation[0] > r.TailLocation[0] && r.HeadLocation[1] > r.TailLocation[1] {
			r.TailLocation[0]++
			r.TailLocation[1]++
		}
		if r.HeadLocation[0] > r.TailLocation[0] && r.HeadLocation[1] < r.TailLocation[1] {
			r.TailLocation[0]++
			r.TailLocation[1]--
		}
		if r.HeadLocation[0] < r.TailLocation[0] && r.HeadLocation[1] > r.TailLocation[1] {
			r.TailLocation[0]--
			r.TailLocation[1]++
		}
		if r.HeadLocation[0] < r.TailLocation[0] && r.HeadLocation[1] < r.TailLocation[1] {
			r.TailLocation[0]--
			r.TailLocation[1]--
		}

	}

	if tailx != r.TailLocation[0] || taily != r.TailLocation[1] {
		r.addVisitedLocation()
	}
	// If the new tail location has not been visited, add it to the visitedPoints slice and increment totalVisited
}

var visitedLocationsByTail [][]int

func (r *Rope) addVisitedLocation() {
	copy := make([]int, len(r.TailLocation))
	copy[0] = r.TailLocation[0]
	copy[1] = r.TailLocation[1]
	visitedLocationsByTail = append(visitedLocationsByTail, copy)
}

func (r *Rope) isOverlappingWithHead() bool {
	// Get the x and y coordinates of the head and tail locations
	headX, headY := r.HeadLocation[0], r.HeadLocation[1]
	tailX, tailY := r.TailLocation[0], r.TailLocation[1]

	// Check if the tail is in any of the four diagonal locations adjacent to the head
	if headX == tailX && headY == tailY {
		return true
	}
	return false
}

func Maim() {
	V := generateVector(5, 6)
	fillVectorWithDots(V)
	R := &Rope{
		TwoDVector:             V,
		StartingPoint:          []int{4, 0},
		HeadLocation:           []int{4, 0},
		TailLocation:           []int{4, 0},
		visitedLocationsByTail: [][]int{},
		VisitedNumberByTail:    &totalVisited,
	}

	lines := getInput()
	for _, line := range lines {
		direction, numberOfMove := getDirection(line), getNumberOfMove(line)
		num, _ := strconv.Atoi(numberOfMove)
		R.Move(direction, num)
	}
	total := 0
	for range visitedLocationsByTail {
		total++
	}
	fmt.Println(total)
}

func (r *Rope) Print() {
	for i := range r.TwoDVector {
		for j := range r.TwoDVector[i] {
			fmt.Print(r.TwoDVector[i][j])
		}
		fmt.Println()
	}

	fmt.Println("-----------------------1-")
}

func generateVector(x int, y int) TwoDVector {
	vector := make(TwoDVector, x)
	for i := range vector {
		vector[i] = make([]string, y)
	}
	return vector
}

func fillVectorWithDots(vector TwoDVector) {
	for i := range vector {
		for j := range vector[i] {
			vector[i][j] = "."
		}
	}
}

// /*  Util Functions */
func getDirection(s string) string {
	return s[0:1]
}

func getNumberOfMove(s string) string {
	return s[2:3]
}

//	func printVector(s TwoDVector) {
//		for _, line := range s {
//			fmt.Println(line)
//		}
//	}
func getInput() []string {
	input, _ := os.ReadFile("input9.txt")
	lines := strings.Split(string(input), "\n")[:2000]
	fmt.Println(lines)
	return lines
}

//
// /* Generating 2d vector */
// func create2DVector(vs, hs int) TwoDVector {
// 	s := make(TwoDVector, vs)
//
// 	// fill 2d with dots .
// 	for i := range s {
// 		s[i] = make([]string, hs)
// 		for j := range s[i] {
// 			s[i][j] = "."
// 		}
// 	}
//
// 	s[4][0] = "s"
//
// 	return s
// }
//
// func (r *Rope) printRope() {
// 	for _, line := range r.TwoDVector {
// 		fmt.Println(line)
// 	}
// }
