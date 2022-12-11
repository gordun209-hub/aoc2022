package day11

import "fmt"

type Monke struct {
	startingItems []int
	name          string
	inspectTime   int
}

func (M *Monke) getFirstItem() int {
	if len(M.startingItems) == 0 {
		return 0
	}
	return M.startingItems[0]
}

func (m *Monke) dropItem() {
	m.startingItems = m.startingItems[:len(m.startingItems)-1]
}

func (m *Monke) appendItem(i int) {
	newItems := make([]int, len(m.startingItems)+1)
	newItems[0] = i
	copy(newItems[1:], m.startingItems)
	m.startingItems = newItems
}

func (m *Monke) throwItem(toMonke *Monke, item int) {
	m.dropItem()
	toMonke.appendItem(item)
}

// func (m *Monke) throwItem(i int) {
// 	m.startingItems = append(m.startingItems, i)
// }

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func Maim() {
	monke0 := Monke{
		startingItems: []int{63, 84, 80, 83, 84, 53, 88, 72},
		name:          string("monke0"),
		inspectTime:   0,
	}
	monke1 := Monke{
		startingItems: []int{67, 56, 92, 88, 84},
		name:          string("monke1"),
		inspectTime:   0,
	}
	monke2 := Monke{
		startingItems: []int{52},
		name:          string("monke2"),
		inspectTime:   0,
	}
	monke3 := Monke{
		startingItems: []int{59, 53, 60, 92, 69, 72},
		name:          string("monke3"),
		inspectTime:   0,
	}

	monke4 := Monke{
		startingItems: []int{61, 52, 55, 61},
		name:          string("monke4"),
		inspectTime:   0,
	}

	monke5 := Monke{
		startingItems: []int{79, 53},
		name:          string("monke5"),
		inspectTime:   0,
	}
	monke6 := Monke{
		startingItems: []int{59, 86, 67, 95, 92, 77, 91},
		name:          string("monke6"),
		inspectTime:   0,
	}
	monke7 := Monke{
		startingItems: []int{58, 83, 89},
		name:          string("monke7"),
		inspectTime:   0,
	}
	monkaes := []*Monke{
		&monke0, &monke1, &monke2, &monke3,
		&monke4, &monke5, &monke6, &monke7,
	}

	const la = 13 * 11 * 2 * 5 * 7 * 3 * 19 * 17

	for i := 0; i < 10000; i++ {
		for _, monke := range monkaes {
			if monke.name == monke0.name {
				for _, item := range monke0.startingItems {
					monke0.inspectTime++
					worryLevel := item * 11

					if worryLevel%13 == 0 {
						monke0.throwItem(&monke4, worryLevel%la)
					} else {
						monke0.throwItem(&monke7, worryLevel%la)
					}
				}
			} else if monke.name == monke1.name {
				for _, item := range monke1.startingItems {
					monke1.inspectTime++
					worryLevel := item + 4

					if worryLevel%11 == 0 {
						monke1.throwItem(&monke5, worryLevel%la)
					} else {
						monke1.throwItem(&monke3, worryLevel%la)
					}
				}
			} else if monke.name == monke2.name {
				for _, item := range monke2.startingItems {
					monke2.inspectTime++
					worryLevel := item * item

					if worryLevel%2 == 0 {
						monke2.throwItem(&monke3, worryLevel%la)
					} else {
						monke2.throwItem(&monke1, worryLevel%la)
					}
				}
			} else if monke.name == monke3.name {
				for _, item := range monke3.startingItems {
					monke3.inspectTime++
					worryLevel := item + 2

					if worryLevel%5 == 0 {
						monke3.throwItem(&monke5, worryLevel%la)
					} else {
						monke3.throwItem(&monke6, worryLevel%la)
					}
				}
			} else if monke.name == monke4.name {
				for _, item := range monke4.startingItems {
					monke4.inspectTime++
					worryLevel := item + 3

					if worryLevel%7 == 0 {
						monke4.throwItem(&monke7, worryLevel%la)
					} else {
						monke4.throwItem(&monke2, worryLevel%la)
					}
				}
			} else if monke.name == monke5.name {
				for _, item := range monke5.startingItems {
					monke5.inspectTime++
					worryLevel := item + 1

					if worryLevel%3 == 0 {
						monke5.throwItem(&monke0, worryLevel%la)
					} else {
						monke5.throwItem(&monke6, worryLevel%la)
					}
				}
			} else if monke.name == monke6.name {
				for _, item := range monke6.startingItems {
					monke6.inspectTime++
					worryLevel := item + 5

					if worryLevel%19 == 0 {
						monke6.throwItem(&monke4, worryLevel%la)
					} else {
						monke6.throwItem(&monke0, worryLevel%la)
					}
				}
			} else if monke.name == monke7.name {
				for _, item := range monke7.startingItems {
					monke7.inspectTime++
					worryLevel := item * 19

					if worryLevel%17 == 0 {
						monke7.throwItem(&monke2, worryLevel%la)
					} else {
						monke7.throwItem(&monke1, worryLevel%la)
					}
				}
			}
		}
	}
	for _, monke := range monkaes {
		fmt.Println(monke.inspectTime, monke.name, 174976*174975)
	}
}

// for i := 0; i < 10000; i++ {
// 	for _, monke := range monkaes {
// 		if monke.name == monke0.name {
// 			for _, item := range monke0.startingItems {
// 				monke0.inspectTime++
// 				worryLevel := item * 11
//
// 				if worryLevel%13 == 0 {
// 					monke0.throwItem(&monke4, worryLevel)
// 				} else {
// 					monke0.throwItem(&monke7, worryLevel)
// 				}
// 			}
// 		} else if monke.name == monke1.name {
// 			for _, item := range monke1.startingItems {
// 				monke1.inspectTime++
// 				worryLevel := item + 4
//
// 				if worryLevel%11 == 0 {
// 					monke1.throwItem(&monke5, worryLevel)
// 				} else {
// 					monke1.throwItem(&monke3, worryLevel)
// 				}
// 			}
// 		} else if monke.name == monke2.name {
// 			for _, item := range monke2.startingItems {
// 				monke2.inspectTime++
// 				worryLevel := item * item
//
// 				if worryLevel%2 == 0 {
// 					monke2.throwItem(&monke3, worryLevel)
// 				} else {
// 					monke2.throwItem(&monke1, worryLevel)
// 				}
// 			}
// 		} else if monke.name == monke3.name {
// 			for _, item := range monke3.startingItems {
// 				monke3.inspectTime++
// 				worryLevel := item + 2
//
// 				if worryLevel%5 == 0 {
// 					monke3.throwItem(&monke5, worryLevel)
// 				} else {
// 					monke3.throwItem(&monke6, worryLevel)
// 				}
// 			}
// 		} else if monke.name == monke4.name {
// 			for _, item := range monke4.startingItems {
// 				monke4.inspectTime++
// 				worryLevel := item + 3
//
// 				if worryLevel%7 == 0 {
// 					monke4.throwItem(&monke7, worryLevel)
// 				} else {
// 					monke4.throwItem(&monke2, worryLevel)
// 				}
// 			}
// 		} else if monke.name == monke5.name {
// 			for _, item := range monke5.startingItems {
// 				monke5.inspectTime++
// 				worryLevel := item + 1
//
// 				if worryLevel%3 == 0 {
// 					monke5.throwItem(&monke0, worryLevel)
// 				} else {
// 					monke5.throwItem(&monke6, worryLevel)
// 				}
// 			}
// 		} else if monke.name == monke6.name {
// 			for _, item := range monke6.startingItems {
// 				monke6.inspectTime++
// 				worryLevel := item + 5
//
// 				if worryLevel%19 == 0 {
// 					monke6.throwItem(&monke4, worryLevel)
// 				} else {
// 					monke6.throwItem(&monke0, worryLevel)
// 				}
// 			}
// 		} else if monke.name == monke7.name {
// 			for _, item := range monke7.startingItems {
// 				monke7.inspectTime++
// 				worryLevel := item * 19
//
// 				if worryLevel%17 == 0 {
// 					monke7.throwItem(&monke2, worryLevel)
// 				} else {
// 					monke7.throwItem(&monke1, worryLevel)
// 				}
// 			}
// 		}
// 	}

// monke0 := Monke{
// 	startingItems: []int{63, 84, 80, 83, 84, 53, 88, 72},
// 	name:          string("monke0"),
// 	inspectTime:   0,
// }
// monke1 := Monke{
// 	startingItems: []int{67, 56, 92, 88, 84},
// 	name:          string("monke1"),
// 	inspectTime:   0,
// }
// monke2 := Monke{
// 	startingItems: []int{52},
// 	name:          string("monke2"),
// 	inspectTime:   0,
// }
// monke3 := Monke{
// 	startingItems: []int{59, 53, 60, 92, 69, 72},
// 	name:          string("monke3"),
// 	inspectTime:   0,
// }
//
// monke4 := Monke{
// 	startingItems: []int{61, 52, 55, 61},
// 	name:          string("monke4"),
// 	inspectTime:   0,
// }
//
// monke5 := Monke{
// 	startingItems: []int{79, 53},
// 	name:          string("monke5"),
// 	inspectTime:   0,
// }
// monke6 := Monke{
// 	startingItems: []int{59, 86, 67, 95, 92, 77, 91},
// 	name:          string("monke6"),
// 	inspectTime:   0,
// }
// monke7 := Monke{
// 	startingItems: []int{58, 83, 89},
// 	name:          string("monke7"),
// 	inspectTime:   0,
// }
// monkaes := []*Monke{
// 	&monke0, &monke1, &monke2, &monke3,
// 	&monke4, &monke5, &monke6, &monke7,
// }
