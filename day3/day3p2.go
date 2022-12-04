package day3

import (
	"fmt"
	"os"
	"strings"
)

var letters = ""

const sampleData = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func Main2() {
	content, err := os.ReadFile("input3.txt")
	if err != nil {
		fmt.Println(err)
	}
	w := strings.Split(string(content), "\n")
	g := removeLastElement(w)

	for i := 0; i < len(g); i += 3 {
		group := getThreeLineGroup(g[i : i+3])
		fmt.Println(group)
		fmt.Println("-0-----")
		letters += findCommonLetters(group)
	}
	fmt.Println(getSumOfChars(letters))
}

func removeLastElement(s []string) []string {
	return s[:len(s)-1]
}

// create group for each three line
func getThreeLineGroup(s []string) []string {
	return []string{s[0], s[1], s[2]}
}

func findCommonLetters(s []string) string {
	var common string
	for _, v := range s[0] {
		for _, v2 := range s[1] {
			if v == v2 {
				for _, v3 := range s[2] {
					if v == v3 {
						return string(v)
					}
				}
			}
		}
	}
	return common
}
