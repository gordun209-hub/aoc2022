package day9

// ......
// ......
// ......
// ......
// ......

// type (
// 	TwoDVector = [][]string
// 	Rope       struct {
// 		TwoDVector
// 		StartingPoint []int
// 		HeadLocation  []int
// 		TailLocation  []int
// 	}
// )
//
// func (r Rope) Print() {
// 	for _, line := range r.TwoDVector {
// 		fmt.Println(line)
// 	}
// }
//
// func (r *Rope) UpdateTailLocation() {
// 	head := r.HeadLocation
// 	tail := r.TailLocation
//
// 	// Check if the tail is not adjacent to the head
// 	if (head[0] != tail[0] || (head[1] != tail[1]+1 && head[1] != tail[1]-1)) && (head[1] != tail[1] || (head[0] != tail[0]+1 && head[0] != tail[0]-1)) {
// 		// Move the tail towards the head until they are adjacent
// 		if head[0] > tail[0] {
// 			tail[0]++
// 		} else if head[0] < tail[0] {
// 			tail[0]--
// 		}
// 		if head[1] > tail[1] {
// 			tail[1]++
// 		} else if head[1] < tail[1] {
// 			tail[1]--
// 		}
//
// 		// Mark the tail's new location as visited
// 		r.TwoDVector[tail[0]][tail[1]] = "T"
// 	}
// }
//
// // need to follow up tail
// func (r Rope) Move(direction string, numverOfMove int) {
// 	for i := 0; i <= numverOfMove-1; i++ {
// 		switch direction {
//
// 		case "U":
// 			// fill wector with $
// 			r.UpdateTailLocation()
// 			r.HeadLocation[0] = r.HeadLocation[0] - 1
//
// 			r.UpdateTailLocation()
// 			// set previous tail to $
//
// 		case "D":
// 			// fill wector with $
// 			r.UpdateTailLocation()
// 			r.HeadLocation[0] = r.HeadLocation[0] + 1
//
// 			r.UpdateTailLocation()
// 			// set previous tail to $
//
// 		case "L":
// 			// fill wector with $
// 			r.UpdateTailLocation()
// 			r.HeadLocation[1] = r.HeadLocation[1] - 1
//
// 			r.UpdateTailLocation()
// 			// set previous tail to $
//
// 		case "R":
//
// 			// fill wector with $
// 			r.UpdateTailLocation()
// 			r.HeadLocation[1] = r.HeadLocation[1] + 1
//
// 			r.UpdateTailLocation()
// 			// set previous tail to $
//
// 		}
// 	}
// }
//
// func Maim() {
// 	// lines := getInput()
//
// 	vec := create2DVector(5, 6)
// 	rope := Rope{
// 		TwoDVector:    vec,
// 		StartingPoint: []int{4, 0},
// 		HeadLocation:  []int{4, 0},
// 		TailLocation:  []int{4, 0},
// 	}
// 	fmt.Println(rope.HeadLocation)
// 	fmt.Println(rope.TailLocation)
// 	rope.Move("R", 4)
// 	rope.Move("U", 4)
// 	rope.Move("L", 3)
// 	rope.Move("D", 1)
// 	rope.Move("R", 4)
// 	// rope.Move("D", 1)
// 	// rope.Move("L", 5)
// 	// rope.Move("R", 2)
//
// 	rope.TwoDVector[4][0] = "s"
// 	rope.Print()
// }
//
// // /*  Util Functions */
// // func getDirection(s string) string {
// // 	return s[0:1]
// // }
// //
// // func getNumberOfMove(s string) string {
// // 	return s[2:3]
// // }
// //
// // func printVector(s TwoDVector) {
// // 	for _, line := range s {
// // 		fmt.Println(line)
// // 	}
// // }
// //
// // func getInput() []string {
// // 	input, _ := os.ReadFile("sample9.txt")
// // 	lines := strings.Split(string(input), "\n")[:8]
// // 	return lines
// // }
// //
// // /* Generating 2d vector */
// // func create2DVector(vs, hs int) TwoDVector {
// // 	s := make(TwoDVector, vs)
// //
// // 	// fill 2d with dots .
// // 	for i := range s {
// // 		s[i] = make([]string, hs)
// // 		for j := range s[i] {
// // 			s[i][j] = "."
// // 		}
// // 	}
// //
// // 	s[4][0] = "s"
// //
// // 	return s
// // }
