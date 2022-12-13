package day13

import (
	"fmt"
	"strings"
)

const input = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

type pair struct {
	first  string
	second string
	id     int
}

// need to create arrays or smth idk im 8
// ex
// takes
// [1, [2]] and []
// compares
// 1 and []
// result true

// compare [[4,4] ,4,4] vs [[4,4],4,4,4]

func (p *pair) compare(fst, scnd string) bool {
	trimmedright := replaceFirstBracket(fst)
	trimmedleft := replaceFirstBracket(scnd)
	fmt.Println(trimmedleft, "una", trimmedright)
	return true
}

func replaceFirstBracket(s string) string {
	return strings.Replace(s, "[", "", 1)
}

func isMixedType(s string) bool {
	return strings.Contains(s, "[")
}

var id int

func Maim() {
	// start from beginin remove starting [ bruh

	lines := lines()
	for _, line := range lines {
		pair := initializePair(line)
		pair.compare(pair.first, pair.second)

	}
}

func removeBrackets(s string) string {
	return strings.ReplaceAll(s, "[", "")
}

func lines() []string {
	return strings.Split(string(input), "\n\n")
}

func initializePair(s string) *pair {
	pairs := strings.Split(s, "\n")
	id++
	return &pair{pairs[0], pairs[1], id}
}
