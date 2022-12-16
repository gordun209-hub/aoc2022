package utils

import "strings"

func GetLines(input string) []string {
	return strings.Split(input, "\n\n")
}

type Ordered interface {
	~float64 | ~int | ~string
}

// usage: numbers:= []float64{3.5,-2.4,12.8,9.1}
// bubblesort[float64](numbers)
func BubbleSort[T Ordered](data []T) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func Quicksort[T Ordered](data []T, low, high int) {
	if low < high {
		pivot := patrition(data, low, high)
		Quicksort(data, low, pivot)
		Quicksort(data, pivot+1, high)
	}
}

func patrition[T Ordered](data []T, low, high int) int {
	// Pick lowest bound element as pivot
	pivot := data[low]

	i := low
	j := high
	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}
		for data[j] > pivot && j > low {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[low] = data[j]
	data[j] = pivot
	return j
}

// Concurret quicksort

func InsertSort[T Ordered](data []T) {
	i := 1
	for i < len(data) {
		h := data[i]
		j := i - 1
		for j >= 0 && h < data[j] {
			data[j+1] = data[j]
			j -= 1
		}
		data[j+1] = h
		i += 1
	}
}

func IsSorted[T Ordered](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func Merge[T Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

func MergeSort[T Ordered](data []T) []T {
	if len(data) > 100 {
		middle := len(data) / 2
		left := data[:middle]
		right := data[middle:]
		data = Merge(MergeSort(right), MergeSort(left))
	} else {
		InsertSort(data)
	}
	return data
}

// Search

const size = 100_000_000

func LinearSearch[T Ordered](slice []T, target T) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false
}

func BinarySearch[T Ordered](slice []T, target T) bool {
	low := 0
	high := len(slice) - 1
	for low <= high {
		median := (low + high) / 2
		if slice[median] < target {
			low = median + 1
		} else {
			high = median + 1
		}
	}
	if low == len(slice) || slice[low] != target {
		return false
	}
	return true
}

type Stack[T Ordered] struct {
	items []T
}

func getZero[T Ordered]() T {
	var result T
	return result
}

func (s *Stack[T]) Push(item T) {
	if item != getZero[T]() {
		s.items = append(s.items, item)
	}
}

func (s *Stack[T]) Pop() T {
	length := len(s.items)
	if length > 0 {
		returnValue := s.items[length-1]
		s.items = s.items[:(length - 1)]
		return returnValue
	} else {
		return getZero[T]()
	}
}

func (s Stack[T]) Top() T {
	length := len(s.items)
	if length > 0 {
		return s.items[length-1]
	} else {
		return getZero[T]()
	}
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type NodeStack[T any] struct {
	first *Node[T]
}

func (s *NodeStack[T]) Push(item T) {
	newNode := Node[T]{item, nil}
	newNode.next = s.first
	s.first = &newNode
}

func (s *NodeStack[T]) Top() T {
	return s.first.value
}

func (s *NodeStack[T]) Pop() T {
	result := s.first.value
	s.first = s.first.next
	return result
}

func (s *NodeStack[T]) IsEmpty() bool {
	return s.first == nil
}

type BinaryTree struct {
	Root     *BNode
	NumNodes int
}
type BNode struct {
	Value string
	Left  *BNode
	Right *BNode
}

type nodePair struct {
	Val1, Val2 string
}

type nodePos struct {
	Val  string
	YPos int
	XPos int
}

var (
	data   []nodePos
	endPos []nodePair
)
