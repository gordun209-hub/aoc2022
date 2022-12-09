package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sampleInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

const LIMIT = 100000

var (
	FS      = make(map[string]int)
	path    []string
	maxUsed = 70000000 - 30000000
)

func ReadInput() *os.File {
	readfile, err := os.Open("in7.txt")
	if err != nil {
		fmt.Println(err)
	}
	return readfile
}

func fields(s string) []string {
	words := strings.Fields(s)
	return words
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func Maim() {
	scanner := bufio.NewScanner(strings.NewReader(sampleInput))
	for scanner.Scan() {
		line := scanner.Text()
		words := fields(line)
		if words[1] == "cd" {
			if words[2] == ".." {
				// go back
				path = path[:len(path)-1]
			} else {
				// append file name to path
				path = append(path, words[2])
			}
		} else if words[1] == "ls" || words[0] == "dir" {
			// list files
			continue
		} else {
			sz := words[0]
			fmt.Println(path)
			// Add this file's size to the current directory size *and* the size of all parents
			for i := 1; i <= len(path); i++ {
				w := convertToInt(sz)
				FS[strings.Join(path[:i], "/")] += w
			}
		}
	}

	totalUsed := FS["/"]
	NeedFree := totalUsed - maxUsed

	p1 := 0
	p2 := 1e9
	for _, v := range FS {
		if v <= 100000 {
			p1 += v
		}
		if v >= NeedFree {
			p2 = min(p2, v)
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

func min(a float64, b int) float64 {
	if int(a) < b {
		return a
	}
	return float64(b)
}
