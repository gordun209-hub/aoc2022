package main

import (
	_ "github.com/gordun209-hub/aoc2022/day1"
	_ "github.com/gordun209-hub/aoc2022/day10"
	_ "github.com/gordun209-hub/aoc2022/day11"
	_ "github.com/gordun209-hub/aoc2022/day12"
	_ "github.com/gordun209-hub/aoc2022/day13"
	d14 "github.com/gordun209-hub/aoc2022/day14"
	_ "github.com/gordun209-hub/aoc2022/day2"
	_ "github.com/gordun209-hub/aoc2022/day3"
	_ "github.com/gordun209-hub/aoc2022/day4"
	_ "github.com/gordun209-hub/aoc2022/day5"
	_ "github.com/gordun209-hub/aoc2022/day6"
	_ "github.com/gordun209-hub/aoc2022/day7"
	_ "github.com/gordun209-hub/aoc2022/day8"
	_ "github.com/gordun209-hub/aoc2022/day9"
)

func main() {
	d14.Maim()
	// [x+1, y] = down
	// [x-1,y] = up
	// [x,y+1] = right
	// [x,y-1] = left
}
