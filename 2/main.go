package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)
	total := 0
	
	for scan.Scan() {
		line := scan.Text()
		
		res := strings.Split(line, " ")

		x, y := res[0][0], res[1][0]
		score := 0

		// used ascii value to simplify
		switch y {
		case 'X':
			score = (int(x) - 63) % 3 + 1 // (v - 1 - 65 + 3) % 3 + 1
		case 'Y':
			// draw
			score = int(x) - 61 // v - 65 + 1 + 3 
		case 'Z':
			score = (int(x) - 64) % 3 + 7 // (v + 1 - 65) % 3 + 7
		}
		total += score
	}
	fmt.Println(total)
}