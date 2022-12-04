package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// .........
// .234..... 2-4
// .....678. 6-8
// .23...... 2-3

// input (2,5) output [2 3 4 5]

func Maim2() {
	count := 0
	content, err := os.ReadFile("input4.txt")
	if err != nil {
		fmt.Println(err)
	}
	contentt := strings.Split(string(content), "\n")
	formatted := removeLastElement(contentt)

	for _, v := range formatted {
		splitted := strings.Split(v, ",")
		ww := strings.Trim(splitted[0], " ")
		ww2 := strings.Trim(splitted[1], " ")

		// bura hatali
		firstpairx, _ := strconv.Atoi(string(strings.Split(ww, "-")[0]))
		firstpairy, _ := strconv.Atoi(string(strings.Split(ww, "-")[1]))
		secondpairx, _ := strconv.Atoi(string(strings.Split(ww2, "-")[0]))
		secondpairy, _ := strconv.Atoi(string(strings.Split(ww2, "-")[1]))
		enumerated := enumerateInterval(firstpairx, firstpairy)
		enumerated2 := enumerateInterval(secondpairx, secondpairy)
		if containsAny(enumerated, enumerated2) || containsAny(enumerated2, enumerated) {
			count++
		}

	}
	fmt.Println(count)
}

func containsAny(needle []int, haystack []int) bool {
	for _, v := range needle {
		if contains(v, haystack) {
			return true
		}
	}
	return false
}

