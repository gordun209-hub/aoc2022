package day9

import (
	"fmt"
)

type TwoDVector = [][]string

type Rope struct {
	TwoDVector    TwoDVector
	StartingPoint []int
	HeadLocation  []int
	TailLocation  []int
	twodvector    TwoDVector
}

func (r *Rope) clear() {
	for i := range r.TwoDVector {
		for j := range r.TwoDVector[i] {
			r.TwoDVector[i][j] = "."
		}
	}
	r.TwoDVector[r.TailLocation[0]][r.TailLocation[1]] = "T"
	r.TwoDVector[r.HeadLocation[0]][r.HeadLocation[1]] = "H"
	fmt.Println(r.TailLocation)
	fmt.Println(r.HeadLocation)
}

const Edges = 4

func (r Rope) Move(direction string, numverOfMove int) {
	for i := 0; i <= numverOfMove-1; i++ {
		switch direction {

		case "U":
			r.clear()
			r.MoveTail()
			r.HeadLocation[0] = r.HeadLocation[0] - 1
			r.Print()
		case "D":
			r.clear()
			r.HeadLocation[0] = r.HeadLocation[0] + 1
			r.MoveTail()
			r.Print()
		case "L":
			r.clear()
			r.HeadLocation[1] = r.HeadLocation[1] - 1
			r.MoveTail()
			r.Print()
		case "R":
			r.clear()
			r.HeadLocation[1] = r.HeadLocation[1] + 1
			r.MoveTail()
			r.Print()

		}

		r.TwoDVector[4][0] = "s"
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
	// same row
	if r.DoesTailNeedToMove() {

		fmt.Println(r.HeadLocation, r.TailLocation)
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
	}
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
	TailVisiteds := generateVector(5, 6)
	fillVectorWithDots(TailVisiteds)
	fillVectorWithDots(V)
	R := Rope{
		TwoDVector:    V,
		StartingPoint: []int{4, 0},
		HeadLocation:  []int{4, 0},
		TailLocation:  []int{4, 0},
	}

	R.Move("R", 4)
	R.Move("U", 1)
	// R.Move("L", 0)
	R.TwoDVector[R.TailLocation[0]][R.TailLocation[1]] = "T"
	R.TwoDVector[R.HeadLocation[0]][R.HeadLocation[1]] = "H"
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

// func (r *Rope) AdjacentPointsAreaToHead() [][]int {
// 	head := r.HeadLocation
// 	adjacentPoints := [][]int{}
// 	// check if the head is not on the edge of the vector
// 	if head[0] != 0 {
// 		adjacentPoints = append(adjacentPoints, []int{head[0] - 1, head[1]})
// 	}
// 	if head[0] != len(r.TwoDVector)-1 {
// 		adjacentPoints = append(adjacentPoints, []int{head[0] + 1, head[1]})
// 	}
// 	if head[1] != 0 {
// 		adjacentPoints = append(adjacentPoints, []int{head[0], head[1] - 1})
// 	}
// 	if head[1] != len(r.TwoDVector[0])-1 {
// 		adjacentPoints = append(adjacentPoints, []int{head[0], head[1] + 1})
// 	}
//
// 	if head[0] != 0 && head[1] != 0 {
// 		adjacentPoints = append(adjacentPoints, []int{head[0] - 1, head[1] - 1})
// 	}
//
// 	if head[0] != 0 && head[1] != len(r.TwoDVector[0])-1 {
// 		adjacentPoints = append(adjacentPoints, []int{head[0] - 1, head[1] + 1})
// 	}
//
// 	return adjacentPoints
// }
//
// func isAdjacentPoint(point, tail []int) bool {
// 	if (point[0] == tail[0] &&
// 		(point[1] == tail[1]+1 || point[1] == tail[1]-1)) ||
// 		(point[1] == tail[1] && (point[0] == tail[0]+1 ||
// 			point[0] == tail[0]-1)) {
// 		return true
// 	}
// 	if point[0] == tail[0] && point[1] == tail[1] {
// 		return true
// 	}
// 	return false
// }
//
// func (r *Rope) Move(direction string, numberOfMove int) {
// 	// fmt.Println("Moving ", direction, numberOfMove)
// 	switch direction {
// 	case "R":
// 		r.moveRight(numberOfMove)
// 	}
// }
//
// func (r *Rope) MoveTailToAdjacentPoint() {
// 	adjacentPoints := r.AdjacentPointsAreaToHead()
// 	for _, point := range adjacentPoints {
// 		if r.TailLocation[0] == point[0] && r.TailLocation[1] == point[1] {
// 			continue
// 		}
// 	}
// }
//
// func (r *Rope) moveRight(numberOfMove int) {
// 	for i := 0; i < numberOfMove; i++ {
// 		r.HeadLocation[1]++
//
// 		adjpoints := r.AdjacentPointsAreaToHead()
// 		// if tail inside adjacent points dont move
// 		// else move tail to adjacent point
// 		for _, point := range adjpoints {
// 			if isAdjacentPoint(point, r.TailLocation) {
// 				continue
// 			}
// 		}
//
// 	}
// }
//
// func Maim() {
// 	// lines := getInput()
// 	// AdjacentPointsAreaToHead
// 	// adjacentPoints
// 	// isAdjacentPoint
// 	// MoveTailToAdjacentPoint
// 	// moveRight
// 	vec := create2DVector(5, 6)
// 	vec2 := create2DVector(5, 6)
// 	rope := Rope{
// 		TwoDVector:             vec,
// 		StartingPoint:          []int{4, 0},
// 		HeadLocation:           []int{4, 0},
// 		TailLocation:           []int{4, 0},
// 		visitedLocationsByTail: vec2,
// 	}
// 	rope.HeadLocation = []int{4, 0}
// 	fmt.Println(rope.HeadLocation, rope.TailLocation)
//
// 	rope.Move("R", 1)
// 	fmt.Println(rope.AdjacentPointsAreaToHead())
// 	// rope.UpdateTailLocation()
//
// 	// rope.Move("U", 3)
// 	// rope.Move("D", 1)
// 	// rope.Move("R", 4)
// 	// rope.Move("D", 1)
// 	// rope.Move("L", 5)
// 	// rope.Move("R", 2)
//
// 	rope.TwoDVector[4][0] = "s"
// }
//
// /*  Util Functions */
// func getDirection(s string) string {
// 	return s[0:1]
// }
//
// func getNumberOfMove(s string) string {
// 	return s[2:3]
// }
//
// func printVector(s TwoDVector) {
// 	for _, line := range s {
// 		fmt.Println(line)
// 	}
// }
//
// func getInput() []string {
// 	input, _ := os.ReadFile("sample9.txt")
// 	lines := strings.Split(string(input), "\n")[:8]
// 	return lines
// }
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
