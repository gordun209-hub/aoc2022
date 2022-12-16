package day13

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var p Packets

func removeEmptyLinesAndSPaces(lines []string) []string {
	var output []string
	for _, line := range lines {
		if line != "" {
			output = append(output, line)
		}
	}
	return output
}

type Packets struct {
	Val []interface{}
}

func lines() []string {
	f, _ := os.ReadFile("./day13/input.txt")

	return strings.Split(string(f), "\n\r")
}

func Maim() {
	lines := lines()
	fmt.Println(lines)
	for _, line := range lines {
		p.Val = append(p.Val, buildSlice(line))
	}
	fmt.Println(p.Val)
}

func buildSlice(str string) []interface{} {
	var output []interface{}
	err := json.Unmarshal([]byte(str), &output)
	if err != nil {
		fmt.Println(err)
	}
	return output
}
