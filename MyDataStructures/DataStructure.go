package datastructure

import "fmt"

/*
0 1 2 3 4

	|---------|
	|         |
	|         |
	|         |
	|         |
	|---------|
*/

type TwoDV struct {
	topLeft     int
	topRight    int
	bottomRight int
	bottomLeft  int
	middle      int
	size        int
	row         map[int]int
	col         map[int]int
}

func newVec(row, col map[int]int, size int) TwoDV {
	return *&TwoDV{
		row:         row,
		col:         col,
		size:        size,
		topLeft:     -0,
		topRight:    -0,
		bottomRight: -1,
		bottomLeft:  -1,
		middle:      -1,
	}
}

func Maimmm() {
	Vec := newVec(make(map[int]int), make(map[int]int), 100)

	Vec.String()
	// Vec.Append(0, 0, 1)
	// // Vec.Append(0, 0, 2)
	// fmt.Println(Vec.row)
}

func (T TwoDV) Append(x, y int, el int) {
	// TODO
}

// func (T *TwoDV) Pop(x, y int) {
// 	// TODO
// }
//
// func (T *TwoDV) Find(x int) {
// 	// TODO
// }
//
// func (T *TwoDV) Visualize() {
// 	// TODO
// }
//
// func (T *TwoDV) Filter(func() bool) {
// 	// TODO
// }
//
// func (T *TwoDV) Map(func()) {
// 	// TODO
// }
//

func (T *TwoDV) String() {
	fmt.Printf("size %d\nBL: %d, BR: %d\nTL: %d TR: %d Mid:%d\nsize: %d\n", T.size, T.bottomLeft, T.bottomRight, T.topLeft, T.topRight, T.size, T.middle)
}
