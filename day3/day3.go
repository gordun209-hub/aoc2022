package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var priorities string = ""

func Main() {
	file, err := os.Open("input3.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		firsthalf, secondhalf := getHalf(scanner.Text())
		priorities += getCommon(firsthalf, secondhalf)
	}

	fmt.Println(priorities)
}

var ornekdata = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func createGroup(s string) []string {
	return strings.Split(s, "")
}

func takeThreeLine(s []string) []string {
	return s[:3]
}

func getSumOfChars(c string) int {
	sum := 0
	for _, v := range c {
		sum += getLetterIndex(string(v))
	}
	return sum
}

func removeLastIndex(c []string) []string {
	return c[:len(c)-1]
}

// // retunr half of given input
func getHalf(data string) (a, b string) {
	half := len(data) / 2
	return data[:half], data[half:]
}

// returns item that appears in both
func getCommon(a, b string) string {
	var common string
	for _, v := range a {
		if strings.Contains(b, string(v)) {
			return string(v)
		}
	}
	return common
}

func getLetterIndex(c string) int {
	if c == strings.ToUpper(c) {
		return int(c[0]) - 64 + 26
	}
	return int(c[0]) - 96
}
