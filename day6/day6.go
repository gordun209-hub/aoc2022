package day6

import (
	"fmt"
	"os"
)

var (
	sample1 = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	sample2 = `bvwbjplbgvbhsrlpgdmjqwftvncz`
	sample3 = `nppdvjthqldpwncqszvftbrmjlhg`
	sample4 = `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`
)

func Maim() {
	content, err := os.ReadFile("input6.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(findStartOfPacket(string(content)))
}

func findStartOfPacket(s string) int {
	for i := 0; i < len(s)-14; i++ {
		if checkAllDifferent(takeFirstFourteen(s[i : i+14])) {
			return i + 14
		}
	}
	return 0
}

func checkAllDifferent(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func takeFirstFourteen(s string) string {
	if len(s) < 14 {
		return s[0 : len(s)-1]
	}

	return s[0:14]
}
