package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CheckVisibility(grid [][]int, line int, column int) bool {

	visible := true
	for i := 0; i < column; i++ { // check left
		if grid[line][i] >= grid[line][column] {
			visible = false
		}
	}

	if !visible { // if not visible from left check right
		visible = true
		for i := column + 1; i < len(grid); i++ {
			if grid[line][i] >= grid[line][column] {
				visible = false
			}
		}
	}

	if !visible { // if not visible from right check top
		visible = true
		for i := 0; i < line; i++ {
			if grid[i][column] >= grid[line][column] {
				visible = false
			}
		}
	}

	if !visible { // if not visible from top check bottom
		visible = true
		for i := line + 1; i < len(grid); i++ { // left
			if grid[i][column] >= grid[line][column] {
				visible = false
			}
		}
	}

	return visible
}

func ScenicScore(grid [][]int, line int, column int) int {
	ret := 1

	sum := 0
	for i := column - 1; i >= 0; i-- { // check left
		sum++
		if grid[line][i] >= grid[line][column] {
			break
		}
	}
	ret *= sum

	// fmt.Printf("v=%d(%d,%d) %d-", grid[line][column], line, column, sum)

	sum = 0
	for i := column + 1; i < len(grid); i++ { // check right
		sum++
		if grid[line][i] >= grid[line][column] {
			break
		}
	}
	ret *= sum

	// fmt.Print(sum, "-")

	sum = 0
	for i := line - 1; i >= 0; i-- { // check top
		sum++
		if grid[i][column] >= grid[line][column] {
			break
		}

	}
	ret *= sum

	// fmt.Print(sum, "-")

	sum = 0
	for i := line + 1; i < len(grid); i++ { // check bottom
		sum++
		if grid[i][column] >= grid[line][column] {
			break
		}
	}
	ret *= sum

	// fmt.Println(sum, ret)

	return ret
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	var grid [][]int

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		line := scan.Text()

		newRow := make([]int, len(line))
		grid = append(grid, newRow)

		for i := 0; i < len(line); i++ {
			grid[len(grid)-1][i], _ = strconv.Atoi(string(line[i]))
		}

	}

	grid_size := len(grid)
	solve := 0

	for i := 1; i < grid_size-1; i++ {
		for j := 1; j < grid_size-1; j++ {
			if aux := ScenicScore(grid, i, j); aux > solve {
				solve = aux
			}
		}
	}
	fmt.Println(solve)
}
