package day13

//
// import (
// 	"bufio"
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"math"
// 	"os"
// 	"sort"
// )
//
// type (
// 	count                  int
// 	SortableInterfaceSlice []interface{}
// )
//
// func Maim() {
// 	content, _ := os.ReadFile("input13.txt")
// 	contentReader := bytes.NewReader(content)
// 	scanner := bufio.NewScanner(contentReader)
// 	scanner.Split(bufio.ScanLines)
//
// 	sum := 0
// 	index := 1
// 	for scanner.Scan() {
// 		Frow := scanner.Text()
//
// 		if len(Frow) == 0 {
// 			continue
// 		}
//
// 		if !scanner.Scan() {
// 			fmt.Println("zmm")
// 		}
// 		Srow := scanner.Text()
//
// 		left := buildSlice(Frow)
// 		right := buildSlice(Srow)
//
// 		if CompareSlice(left, right) == 0 {
// 			sum += index
// 		}
//
// 		index++
// 	}
//
// 	fmt.Println(sum)
// }
//
// func sortSlice(input SortableInterfaceSlice) {
// 	sort.Slice(input, func(i, j int) bool {
// 		return input[i].(float64) < input[j].(float64)
// 	})
// }
//
// func CompareSlice(L, R interface{}) count {
// 	isLeftSlice, isRightSlice := isSlice(L), isSlice(R)
//
// 	// make it slice if itsnt
// 	if !isLeftSlice && isRightSlice {
// 		return CompareSlice([]interface{}{L}, R)
// 	} else if isLeftSlice && !isRightSlice {
// 		return CompareSlice(L, []interface{}{R})
// 	}
//
// 	leftSlice, rightSlice := L.([]interface{}), R.([]interface{})
// 	minLen := int(math.Min(float64(len(leftSlice)), float64(len(rightSlice))))
//
// 	for i := 0; i < minLen; i++ {
// 		L, R := leftSlice[i], rightSlice[i]
//
// 		if isSlice(L) || isSlice(R) {
// 			if result := CompareSlice(L, R); result != 2 {
// 				return result
// 			}
// 			continue
// 		}
//
// 		if L.(float64) > R.(float64) {
// 			return 1
// 		}
//
// 		if L.(float64) < R.(float64) {
// 			return 0
// 		}
// 	}
//
// 	if len(rightSlice) == len(leftSlice) {
// 		return 2
// 	}
//
// 	if len(rightSlice) == minLen && len(leftSlice) != minLen {
// 		return 1
// 	}
//
// 	return 0
// }
//
// func buildSlice(str string) []interface{} {
// 	var output []interface{}
// 	err := json.Unmarshal([]byte(str), &output)
// 	if err != nil {
// 		fmt.Println("una")
// 	}
// 	return output
// }
//
// func isSlice(input interface{}) bool {
// 	switch input.(type) {
// 	case []interface{}:
// 		return true
// 	}
// 	return false
// }
