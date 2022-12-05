package day5

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// need this to be grouped by numbers

type stackData map[int][]string

// func Main() {
// 	formatdata := splitByEmptyNewline(sampleinput)
// 	stacks := getStacks(formatdata)
// 	rearrangements := getRearrengements(formatdata)
// 	groups := map[int]string{}
// 	parseInstruction(rearrangements[0])
// 	// strr := Rearrenge(stacks, rearrangements[0])
// 	// s := strings.SplitN(stacks, "\n", 3)
// 	// for i := 0; i < len(rearrangements); i++ {
// 	// 	fmt.Println(rearrangements[i])
// 	// }
// 	stacks = formatStack(stacks)
//
// 	groupindex := 1
// 	stackss := strings.Split(stacks, "\n")
// 	stacksss := stackss[0:3]
// 	for _, v := range stacksss {
// 		ww := strings.Split(v, " ")
// 		for i, w := range ww {
// 			if w != "" {
// 				groups[groupindex] += w
// 				fmt.Println(groupindex, i, "ww", w)
// 				groupindex++
// 			}
// 			if groupindex%3 == 1 {
// 				groupindex = 1
// 			}
// 		}
// 	}
// 	fmt.Println(groups)
// }

func Maim() {
	file, err := os.ReadFile("input5.txt")
	if err != nil {
		fmt.Println(err)
	}
	s := string(file)
	format := splitByEmptyNewline(s)

	rearrangements := getRearrengements(format)
	// stackdata := stackData{
	// 	1: {"Z", "N"},
	// 	2: {"M", "C", "D"},
	// 	3: {"P"},
	// }

	stackdata2 := stackData{
		1: {"B", "S", "W", "Z", "G", "P", "W"},
		2: {"J", "V", "B", "C", "Z", "F"},
		3: {"V", "L", "M", "H", "N", "Z", "D", "C"},
		4: {"L", "D", "M", "Z", "P", "F", "J", "B"},
		5: {"V", "F", "C", "G", "J", "B", "Q", "H"},
		6: {"G", "F", "Q", "T", "S", "L", "B"},
		7: {"L", "G", "C", "Z", "V"},
		8: {"N", "G", "L"},
		9: {"J", "F", "H", "C"},
	}
	st := applyInstructions2(stackdata2, rearrangements)
	fmt.Println(st)

	// fmt.Println(stackdata)
	// iterate by row
}

// apply instructions
func applyInstructions(stackdata stackData, rearrangements []string) stackData {
	for i := 0; i < len(rearrangements)-1; i++ {
		move, from, to := parseInstruction(rearrangements[i])
		for j := 0; j < move; j++ {
			stackdata[to] = append(stackdata[to], stackdata[from][len(stackdata[from])-1])
			stackdata[from] = stackdata[from][:len(stackdata[from])-1]
		}
	}
	return stackdata
}

func applyInstructions2(stackdata stackData, rearrangements []string) stackData {
	for i := 0; i < len(rearrangements)-1; i++ {
		move, from, to := parseInstruction(rearrangements[i])
		if move > len(stackdata[from]) {
			move = len(stackdata[from])
		}
		stackdata[to] = append(stackdata[to], stackdata[from][len(stackdata[from])-move:]...)
		stackdata[from] = stackdata[from][:len(stackdata[from])-move]

	}
	return stackdata
}

func formatData(str string) []string {
	strr := strings.ReplaceAll(str, " ", "_")

	return strings.Split(strr, "\n")
}

// TODO
func Rearrenge(stacks string, rearrangements string) string {
	move, from, to := parseInstruction(rearrangements)
	fmt.Println(move, from, to)
	return "todo"
}

func getRearrengements(input []string) []string {
	return strings.Split(input[1], "\n")
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)
}

func parseInstruction(s string) (int, int, int) {
	splitt := strings.Split(s, " ")
	move, _ := strconv.Atoi(splitt[1])
	from, _ := strconv.Atoi(splitt[3])
	to, _ := strconv.Atoi(splitt[5])

	return move, from, to
}
