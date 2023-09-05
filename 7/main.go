package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name      string
	size      int
	directory bool
	parent    *node
	childs    []*node
}

func (n *node) AddChild(child *node) *node {
	n.childs = append(n.childs, child)
	return n
}

func (n *node) printnode(depth int) {
	spaces := strings.Repeat(" ", depth)
	fmt.Print(spaces, "- ", n.name)
	if n.directory {
		fmt.Println(" (dir)")
		for _, child := range n.childs {
			child.printnode(depth + 1)
		}
	} else {
		fmt.Println(" (file, size =", n.size, ")")
	}
}

func (n *node) calculateSizes() int {
	if n.directory == true { // calculate the total size of childs
		for _, child := range n.childs {
			n.size += child.calculateSizes()
		}
	}
	return n.size
}

func (n *node) FindSmallestToFree(smallest *int, necessary int) {
	if n.directory && n.size >= necessary && *smallest > n.size {
		*smallest = n.size
	}
	for _, child := range n.childs {
		child.FindSmallestToFree(smallest, necessary)
	}
}

func main() {
	total_disk_space := 70000000
	needed_space := 30000000

	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)

	root := node{
		name:      "/",
		directory: true,
		parent:    nil,
	}
	current := &root

	for scan.Scan() {
		line := scan.Text()
		tokens := strings.Split(line, " ")

		if tokens[0] == "$" { // command
			if tokens[1] == "cd" { // cd command
				if tokens[2] == ".." {
					current = current.parent
				} else {
					for _, child := range current.childs {
						if child.name == tokens[2] {
							current = child
							break
						}
					}
				}
			} else { // ls command
				//skip to next line
			}
		} else { // must be the output of an ls command
			var add node
			add.parent = current
			if tokens[0] == "dir" { // the output is a directory
				add.directory = true
				add.name = tokens[1]
			} else { // else is a file
				size, _ := strconv.Atoi(tokens[0])
				add.size = size
				add.name = tokens[1]
				add.directory = false
			}
			current = current.AddChild(&add)
		}
	}
	// root.printnode(0)

	total_used := root.calculateSizes()
	unused := total_disk_space - total_used
	size_to_free := needed_space - unused

	root.FindSmallestToFree(&total_used, size_to_free)
	fmt.Println(total_used)
}
