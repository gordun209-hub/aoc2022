package day2

import (
	"bufio"
	"fmt"
	"os"
)

// Points for choose
// 1 Rock
// 2 Paper
// 3 Scissors

// example guide
// A Y
// B X
// C Z

type (
	shape string
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

var point = 0

// Opponent
// A = Rock
// B = paper
// C = scissors

// Me
// X = Rock
// Y = Paper
// Z = Scissors

func getPoint1(me, opponent string) int {
	if me == "X" && opponent == "A" {
		return 1 + 3
	} else if me == "Y" && opponent == "B" {
		return 2 + 3
	} else if me == "Z" && opponent == "C" {
		return 3 + 3
	} else if me == "X" && opponent == "C" {
		return 1 + 6
	} else if me == "X" && opponent == "B" {
		return 1 + 0
	} else if me == "Y" && opponent == "A" {
		return 2 + 6
	} else if me == "Y" && opponent == "C" {
		return 2 + 0
	} else if me == "Z" && opponent == "A" {
		return 3 + 0
	} else if me == "Z" && opponent == "B" {
		return 3 + 6
	}
	return 0
}

// Opponent
// A = Rock
// B = paper
// C = scissors

// Me
// X = lose
// Y = draw
// Z = win
func getPoint2(me, opponent string) int {
	// me scissors enemy rock
	if me == "X" && opponent == "A" {
		return 3 + 0
		// opponent paper me paper
	} else if me == "Y" && opponent == "B" {
		return 2 + 3
		// Z win enemy chose scissors me rock
	} else if me == "Z" && opponent == "C" {
		return 1 + 6
	} else if me == "X" && opponent == "C" {
		return 2 + 0
		// opponent paper me rock because X means lose
	} else if me == "X" && opponent == "B" {
		return 1 + 0
		// opponent chose rock me want draw so chose rock
	} else if me == "Y" && opponent == "A" {
		return 1 + 3
	} else if me == "Y" && opponent == "C" {
		return 3 + 3
	} else if me == "Z" && opponent == "A" {
		return 2 + 6
	} else if me == "Z" && opponent == "B" {
		return 3 + 6
	}
	return 0
}

const input = `A Y
B X
C Z`

type Hand struct {
	opponent string
	me       string
}

func placeHands(opponent, me string) Hand {
	return Hand{opponent, me}
}

func parseInput() []Hand {
	hand := []Hand{}
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hand = append(hand, placeHands(string(scanner.Text()[0]), string(scanner.Text()[2])))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return hand
}

func Day2() {
	handss := parseInput()
	for _, v := range handss {
		point += getPoint2(v.me, v.opponent)
	}
	fmt.Println(point)
	// parseInput()
}
