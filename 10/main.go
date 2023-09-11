package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycle int
var X int

func update() {
	if cycle%40 >= X-1 && cycle%40 <= X+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	cycle++
	if cycle%40 == 0 {
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)

	X, cycle = 1, 0

	for scan.Scan() {

		line := scan.Text()
		token := strings.Split(line, " ")

		if token[0] == "addx" { // add
			update()
			update()
			value, _ := strconv.Atoi(token[1])
			X += value
		} else { // noop
			update()
		}
	}
}
