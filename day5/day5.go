package day5

import (
	"fmt"
	"regexp"
)

var sampleinput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func Main() {
	formatdata := splitByEmptyNewline(sampleinput)
	stacks := getStacks(formatdata)
	rearrangements := getRearrengements(formatdata)
	fmt.Println("Stacks: => ", stacks)
	fmt.Println("rearrengements => ", rearrangements)
}

func getStacks(input []string) string {
	return input[0]
}

func getRearrengements(input []string) string {
	return input[1]
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)
}
