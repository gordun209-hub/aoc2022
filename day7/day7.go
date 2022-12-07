package day7

import (
	"fmt"
)

// first parse input
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

type File struct {
	Parent *Dir
	Name   string
	Size   int
}

func (f *File) String()
func (f *File) GetSize()
func (f *File) New()

type Dir struct {
	Name     string
	Children []*Dir // nil if not a directory
}

func (d *Dir) String()
func (d *Dir) GetSize()
func (d *Dir) New()

func Maim() {
	root := Dir{Name: "/"}
	fmt.Println(root)
}
