package day10

//
// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )
//
// // instructions
// // addx V -> takes two cycles to complete, after two cycles
// // x register is increased by value V
// // noop takes one cycle to complete it has no other effect
//
// const sampleInput = `noop
// addx 3
// addx -5
// `
//
// const ex = `addx 15
// addx -11
// addx 6
// addx -3
// addx 5
// addx -1
// addx -8
// addx 13
// addx 4
// noop
// addx -1
// addx 5
// addx -1
// addx 5
// addx -1
// addx 5
// addx -1
// addx 5
// addx -1
// addx -35
// addx 1
// addx 24
// addx -19
// addx 1
// addx 16
// addx -11
// noop
// noop
// addx 21
// addx -15
// noop
// noop
// addx -3
// addx 9
// addx 1
// addx -3
// addx 8
// addx 1
// addx 5
// noop
// noop
// noop
// noop
// noop
// addx -36
// noop
// addx 1
// addx 7
// noop
// noop
// noop
// addx 2
// addx 6
// noop
// noop
// noop
// noop
// noop
// addx 1
// noop
// noop
// addx 7
// addx 1
// noop
// addx -13
// addx 13
// addx 7
// noop
// addx 1
// addx -33
// noop
// noop
// noop
// addx 2
// noop
// noop
// noop
// addx 8
// noop
// addx -1
// addx 2
// addx 1
// noop
// addx 17
// addx -9
// addx 1
// addx 1
// addx -3
// addx 11
// noop
// noop
// addx 1
// noop
// addx 1
// noop
// noop
// addx -13
// addx -19
// addx 1
// addx 3
// addx 26
// addx -30
// addx 12
// addx -1
// addx 3
// addx 1
// noop
// noop
// noop
// addx -9
// addx 18
// addx 1
// addx 2
// noop
// noop
// addx 9
// noop
// noop
// noop
// addx -1
// addx 2
// addx -37
// addx 1
// addx 3
// noop
// addx 15
// addx -21
// addx 22
// addx -6
// addx 1
// noop
// addx 2
// addx 1
// noop
// addx -10
// noop
// noop
// addx 20
// addx 1
// addx 2
// addx 2
// addx -6
// addx -11
// noop
// noop
// noop`
//
// var (
// 	register int
// 	cycle    int
// )
//
// type Machine struct {
// 	Reg, Cyc, tot int
// 	draw          string
// }
//
// // Part1
// func (m *Machine) findSums() {
// 	if m.Cyc == 40 || m.Cyc == 80 || m.Cyc == 140 || m.Cyc == 180 || m.Cyc == 220 {
// 		m.tot += m.Cyc * m.Reg
// 	}
// }
//
// func (m *Machine) increaseCycle() {
// 	m.Cyc++
// }
//
// func (m *Machine) drawDot() {
// 	m.draw += "."
// }
//
// func (m *Machine) drawOteki() {
// 	m.draw += "#"
// }
//
// func (m *Machine) drawNewLine() {
// 	m.draw += "\n"
// }
//
// func (m *Machine) Draw() {
// 	if m.Overlaps() {
// 		fmt.Println(m.Cyc, m.Reg)
// 		m.draw += "#"
// 	} else {
// 		fmt.Println(m.Cyc, m.Reg)
// 		m.draw += "."
// 	}
// }
//
// func (m *Machine) Overlaps() bool {
// 	if m.Reg == m.Cyc || m.Reg-1 == m.Cyc || m.Reg-2 == m.Cyc {
// 		fmt.Println(m.Reg, m.Cyc)
// 		return true
// 	}
// 	return false
// }
//
// func enumeratePosition(s int) []int {
// 	var pos []int
// 	for i := s - 3; i < s; i++ {
// 		pos = append(pos, i)
// 	}
// 	return pos
// }
//
// func (m *Machine) incCycle(i int) {
// 	if m.Cyc == 40 && m.Cyc >= 0 {
// 		if m.Cyc == 40 {
// 			m.Cyc++
// 		}
// 		if m.Overlaps() {
// 			m.draw += "#"
// 		} else if !m.Overlaps() {
// 			m.draw += "."
// 		}
// 	}
// 	if m.Cyc >= 40 && m.Cyc < 80 {
// 		if m.Cyc == 80 {
// 			m.Cyc++
// 		}
// 		if m.Overlaps() {
// 			m.draw += "#"
// 		} else if !m.Overlaps() {
// 			m.draw += "."
// 		} else {
// 			m.draw += "\n"
// 		}
// 	}
//
// 	if m.Cyc >= 80 && m.Cyc < 120 {
// 		if m.Overlaps() {
// 			m.draw += "#"
// 		} else if !m.Overlaps() {
// 			m.draw += "."
// 		} else {
// 			m.draw += "\n"
// 		}
// 	}
//
// 	if m.Cyc >= 120 && m.Cyc < 160 {
// 		if m.Overlaps() {
// 			m.draw += "#"
// 		} else if !m.Overlaps() {
// 			m.draw += "."
// 		}
// 	}
//
// 	if m.Cyc >= 160 && m.Cyc < 180 {
// 		if m.Cyc == 160 {
// 			m.draw += "\n "
// 			m.Cyc++
// 		}
// 		if m.Overlaps() {
// 			m.draw += "#"
// 		} else if !m.Overlaps() {
// 			m.draw += "."
// 		}
// 	}
// 	m.Cyc += i
// }
//
// func (m *Machine) increaseReg() {
// 	m.Reg++
// }
//
// func (m *Machine) incReg(i int) {
// 	m.Reg += i
// }
//
// func Maim() {
// 	file, _ := os.ReadFile("input10.txt")
//
// 	input := strings.Split(string(file), "\n")[:140]
//
// 	M := &Machine{
// 		Reg: 3,
// 		Cyc: 1,
// 		tot: 1,
// 		// fil with #
// 		draw: string(""),
// 	}
//
// 	fmt.Println(enumeratePosition(5))
// 	for _, line := range input {
//
// 		instructions := strings.Fields(line)
// 		if isNoop(instructions[0]) {
// 			M.incCycle(1)
// 		} else {
// 			amount := getAmount(instructions[1])
// 			M.incCycle(1)
// 			M.incCycle(1)
// 			M.incReg(amount)
//
// 		}
//
// 	}
// 	fmt.Println(M.draw)
// }
//
// func isNoop(s string) bool {
// 	return strings.HasPrefix(s, "n")
// }
//
// func getAmount(s string) int {
// 	if strings.HasPrefix(s, "-") {
// 		in, _ := strconv.Atoi(s[1:])
// 		return -in
// 	} else {
// 		in, _ := strconv.Atoi(s)
// 		return in
// 	}
// }
//
// func getCommand(s []string) (string, string) {
// 	if len(s) == 2 {
// 		return s[0], s[1]
// 	} else {
// 		return s[0], ""
// 	}
// }
