package day9

import (
	"fmt"
	"os"
	"strings"
)

// ......
// ......
// ......
// ......
// ......

type (
	TwoDVector = [][]string
	Rope       struct {
		TwoDVector
		StartingPoint []int
		HeadLocation  []int
		TailLocation  []int
	}
)

func (r Rope) Print() {
	for _, line := range r.TwoDVector {
		fmt.Println(line)
	}
}

func (r Rope) Move(direction string, numverOfMove int) {
	for i := 0; i <= numverOfMove-1; i++ {
		switch direction {
		case "U":
			// fill wector with $
			r.HeadLocation[0] = r.HeadLocation[0] - 1
			r.TwoDVector[r.HeadLocation[0]][r.HeadLocation[1]] = "H"
		case "D":
			// fill wector with $
			r.HeadLocation[0] = r.HeadLocation[0] + 1
			r.TwoDVector[r.HeadLocation[0]][r.HeadLocation[1]] = "H"
		case "L":
			// fill wector with $
			r.HeadLocation[1] = r.HeadLocation[1] - 1
			r.TwoDVector[r.HeadLocation[0]][r.HeadLocation[1]] = "H"
			fmt.Println(r.HeadLocation)
		case "R":

			// fill wector with $
			r.HeadLocation[1] = r.HeadLocation[1] + 1
			r.TwoDVector[r.HeadLocation[0]][r.HeadLocation[1]] = "H"
		}
	}
}

func Maim() {
	// lines := getInput()

	vec := create2DVector(5, 6)
	rope := Rope{
		TwoDVector:    vec,
		StartingPoint: []int{4, 0},
		HeadLocation:  []int{4, 0},
		TailLocation:  []int{4, 0},
	}
	rope.Print()

	// for _, line := range lines {
	// 	// fmt.Println(getDirection(line))
	// 	// direction := getDirection(line)
	// 	// movenum := getNumberOfMove(line)
	// 	// fmt.Println(direction, movenum)
	// }

	rope.Move("R", 4)
	rope.Print()
	fmt.Println("-----------------------")
	rope.Move("U", 4)
	rope.Print()
	rope.Move("L", 3)
	rope.Print()
	rope.Move("D", 1)
	rope.Print()
	rope.Move("R", 4)
	rope.Print()
	rope.Move("D", 1)
	rope.Print()
	rope.Move("L", 5)
	rope.Print()
	rope.Move("R", 2)
}

/*  Util Functions */
func getDirection(s string) string {
	return s[0:1]
}

func getNumberOfMove(s string) string {
	return s[2:3]
}

func printVector(s TwoDVector) {
	for _, line := range s {
		fmt.Println(line)
	}
}

func getInput() []string {
	input, _ := os.ReadFile("sample9.txt")
	lines := strings.Split(string(input), "\n")[:8]
	return lines
}

/* Generating 2d vector */
func create2DVector(vs, hs int) TwoDVector {
	s := make(TwoDVector, vs)

	// fill 2d with dots .
	for i := range s {
		s[i] = make([]string, hs)
		for j := range s[i] {
			s[i][j] = "."
		}
	}

	s[4][0] = "s"

	return s
}
