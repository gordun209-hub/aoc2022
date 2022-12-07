package day7

import (
	"fmt"
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

type Folder struct {
	Name     string
	Size     int
	Children []*Folder // nil if not a directory
}

type File struct {
	Name string
	Size int
}

type Tree struct {
	file File
	size int
	leaf []*Tree
}

func parseData(s []string) {
	for _, line := range s {
		switch {
		case strings.HasPrefix(line, "$ cd"):
			dirToMove := parseDirCommand(line)
			fmt.Println("move to", dirToMove)
			// fmt.Println("move to ", strings.Fields(line))
		case strings.HasPrefix(line, "dir"): // TODO append dir to cd
			// fmt.Println(line, "this is dir")
		case strings.HasPrefix(line, "$ ls"): // TODO list dirs and files
			// fmt.Println(line, " is a command ls")
		case checkIfNumber(line):
			size, name := getFileInfo(line)
			fmt.Println("size is", size, "name is ", name)
		default:
			fmt.Println("error occured please cover all switch cases")
		}
	}
}

func enumerateTree(tree *Tree) {
	fmt.Println(tree.file.Name, tree.size)
	for _, leaf := range tree.leaf {
		enumerateTree(leaf)
	}
}

const sizeLimit = 100000

func Maim() {
	splitted := strings.Split(sampleInput, "\n")
	parseData(splitted)
}

func checkIfNumber(s string) bool {
	_, e := strconv.Atoi(strings.Fields(s)[0])
	return e == nil
}

func parseDirCommand(s string) string {
	ss := strings.Split(s, " ")
	return ss[2]
}

func getFileInfo(s string) (string, string) {
	return strings.Fields(s)[0], strings.Fields(s)[1]
}
