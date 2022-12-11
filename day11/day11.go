package day11

import (
	"fmt"
	"os"
	"strings"
)

const MonkeyCount = 7

type Monke struct {
	items     []int
	operation func(x int) int
	test      func(x int) bool
}

func ReadFile() []string {
	file, _ := os.ReadFile("./day11/sample")

	lines := strings.Split(string(file), "\n")
	return lines
}

func Maim() {
	input := ReadFile()
	fmt.Println(input)
}
