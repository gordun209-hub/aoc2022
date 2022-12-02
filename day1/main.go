package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// first elf total 1000 + 2000 + 3000 = 6k
// second elf total 4000
// third elf 6000 + 5000 = 11k
// so on ..
const sampleData = `1000
2000
3000

4000 

5000
6000

7000
8000
9000

10000`

type data []int

func (d data) Len() int {
	return len(d)
}

func (d data) Less(x, y int) bool {
	return x < y
}
func (a data) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Day1(path string) {
	input, err := getInput(path)
	if err != nil {
		fmt.Println(err)
	}
	max := 0
	for _, val := range input {
		if val > max {
			max = val
		}
	}
	sort.Ints(input)
	fmt.Println(input[len(input)-3:])
	total := 0
	for _, v := range input[len(input)-3:] {
		total += v
	}
	fmt.Println(total)
}

func getInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	total := 0
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			lines = append(lines, total)
			total = 0
		} else {
			val, _ := strconv.Atoi(scanner.Text())

			total += val
		}
	}
	return lines, scanner.Err()
}
