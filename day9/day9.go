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

// ......
// ......
// ......
// ......
// H..... // also initial H = s

type TwoDVector = [][]string

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

func Maim() {
	lines := getInput()

	for _, line := range lines {
		// fmt.Println(getDirection(line))
		fmt.Println(getNumberOfMove(line))
	}
	vec := create2DVector(5, 6)
	printVector(vec)
}

func getDirection(s string) string {
	return s[0:1]
}

func getNumberOfMove(s string) string {
	return s[2:3]
}

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
