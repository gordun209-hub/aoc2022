package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const data = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

var priorities string = ""

func Main() {
	fmt.Println("day3")
	//
	// lines := strings.Split(data, "\n")
	//
	// for _, line := range lines {
	// 	firsthalf, secondhalf := getHalf(line)
	// 	fmt.Println(getCommon(firsthalf, secondhalf))
	// 	priorities += getCommon(firsthalf, secondhalf)
	// }
	//
	// fmt.Println(priorities)
	// fmt.Println(getSumOfChars(priorities))

	file, err := os.Open("input3.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		firsthalf, secondhalf := getHalf(scanner.Text())
		priorities += getCommon(firsthalf, secondhalf)

	}
	fmt.Println(getSumOfChars(priorities))
}

func getSumOfChars(c string) int {
	sum := 0
	for _, v := range c {
		sum += getLetterIndex(string(v))
	}
	return sum
}

// retunr half of given input
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
