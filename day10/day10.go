package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// instructions
// addx V -> takes two cycles to complete, after two cycles
// x register is increased by value V
// noop takes one cycle to complete it has no other effect

const sampleInput = `noop
addx 3
addx -5
`

const ex = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

var (
	register int
	cycle    int
)

type Machine struct {
	Reg, Cyc, tot int
}

func (m *Machine) findSums() {
	if m.Cyc == 20 || m.Cyc == 60 || m.Cyc == 100 || m.Cyc == 140 || m.Cyc == 180 || m.Cyc == 220 {
		m.tot += m.Cyc * m.Reg
	}
}

func (m *Machine) increaseCycle() {
	m.Cyc++
}

func (m *Machine) decreaseCycle() {
	m.Cyc--
}

func (m *Machine) incCycle(i int) {
	m.Cyc += i
}

func (m *Machine) increaseReg() {
	m.Reg++
}

func (m *Machine) incReg(i int) {
	m.Reg += i
}

func Maim() {
	file, _ := os.ReadFile("input10.txt")

	input := strings.Split(string(file), "\n")[:141]

	M := &Machine{
		Reg: 1,
		Cyc: 0,
	}
	fmt.Println(M.Cyc)

	fmt.Println(input)
	for _, line := range input {

		instructions := strings.Fields(line)
		if isNoop(instructions[0]) {
			M.increaseCycle()
			M.findSums()
			if M.Cyc == 140 {
				fmt.Println("wartdik", M.Cyc, M.Reg)
			}
		} else {
			amount := getAmount(instructions[1])
			// that is like -5 or 5 and takes  2 cycle to complete

			M.incCycle(1)
			M.findSums()
			if M.Cyc == 140 {
				fmt.Println("wartdik", M.Cyc, M.Reg)
			}
			M.incCycle(1)
			M.findSums()
			M.incReg(amount)

		}

	}
	fmt.Println(M.tot)
}

func isNoop(s string) bool {
	return strings.HasPrefix(s, "n")
}

func getAmount(s string) int {
	if strings.HasPrefix(s, "-") {
		in, _ := strconv.Atoi(s[1:])
		return -in
	} else {
		in, _ := strconv.Atoi(s)
		return in
	}
}

func getCommand(s []string) (string, string) {
	if len(s) == 2 {
		return s[0], s[1]
	} else {
		return s[0], ""
	}
}
