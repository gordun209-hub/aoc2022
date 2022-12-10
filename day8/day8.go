package day8 //
// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )
//
// const input = `30373
// 25512
// 65332
// 33549
// 35390`
//
// func Maim() {
// 	// Split the input into lines
// 	// line, _ := os.ReadFile("input8.txt")
// 	lines := strings.Split(string(input), "\n")[:5]
// 	// Count the number of visible trees
// 	count := 0
// 	for i := 0; i < len(lines); i++ {
// 		for j := 0; j < len(lines[i]); j++ {
// 			// Convert the character to an integer
// 			num, _ := strconv.Atoi(string(lines[i][j]))
//
// 			// If the number is visible, increment the count
// 			if isVisiblee(lines, num, i, j) {
// 				count++
// 			}
//
// 		}
// 	}
//
// 	// Print the result
// 	fmt.Println(count)
// }
//
// func visibleFromLeft(lines []string, number int, i, j int) bool {
// 	count := 0
// 	for k := j - 1; k >= 0; k-- {
//
// 		num, _ := strconv.Atoi(string(lines[i][k]))
// 		if num >= number {
// 			count++
// 			return false
// 		}
// 	}
// 	fmt.Println(number, i, j)
// 	return true
// }
//
// func visibleFromRight(lines []string, number int, i, j int) bool {
// 	for k := j + 1; k < len(lines[i]); k++ {
// 		num, _ := strconv.Atoi(string(lines[i][k]))
// 		if num >= number {
// 			return false
// 		}
// 	}
//
// 	fmt.Println(number, i, j)
// 	return true
// }
//
// func visibleFromTop(lines []string, number int, i, j int) bool {
// 	for k := i - 1; k >= 0; k-- {
// 		num, _ := strconv.Atoi(string(lines[k][j]))
// 		if num >= number {
// 			return false
// 		}
// 	}
// 	fmt.Println(number, i, j)
// 	return true
// }
//
// func visibleFromBottom(lines []string, number int, i, j int) bool {
// 	for k := i + 1; k < len(lines); k++ {
// 		num, _ := strconv.Atoi(string(lines[k][j]))
// 		if num >= number {
// 			return false
// 		}
// 	}
// 	fmt.Println(number, i, j)
// 	return true
// }
//
// func isVisiblee(lines []string, number int, i, j int) bool {
// 	if i == 0 || i == len(lines)-1 || j == 0 || j == len(lines[i])-1 {
// 		return true
// 	}
//
// 	if !visibleFromLeft(lines, number, i, j) {
// 		if !visibleFromRight(lines, number, i, j) {
// 			if !visibleFromTop(lines, number, i, j) {
// 				if !visibleFromBottom(lines, number, i, j) {
// 					return false
// 				}
// 			}
// 		}
// 	}
// 	return true
// }
//
// // part2
// // multiply the number of trees not blocked by other trees
