package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sampleinput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

// .........
// .234..... 2-4
// .....678. 6-8
// .23...... 2-3

var print = fmt.Println

// input (2,5) output [2 3 4 5]
func enumerateInterval(start, end int) []int {
	var out []int
	for i := start; i <= end; i++ {
		out = append(out, i)
	}
	return out
}

func Maim() {
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
		fmt.Println(ww, ww2)

		// bura hatali
		firstpairx, _ := strconv.Atoi(string(ww[0]))
		firstpairy, _ := strconv.Atoi(string(ww[2]))
		secondpairx, _ := strconv.Atoi(string(ww2[0]))
		secondpairy, _ := strconv.Atoi(string(ww2[2]))
		enumerated := enumerateInterval(firstpairx, firstpairy)
		enumerated2 := enumerateInterval(secondpairx, secondpairy)
		if containsall(enumerated, enumerated2) || containsall(enumerated2, enumerated) {
			count++
		}

	}
	fmt.Println(count)
}

func contains(needle int, haystack []int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func containsall(needle []int, haystack []int) bool {
	for _, v := range needle {
		if !contains(v, haystack) {
			return false
		}
	}
	return true
}

func removeLastElement(s []string) []string {
	return s[:len(s)-1]
}
