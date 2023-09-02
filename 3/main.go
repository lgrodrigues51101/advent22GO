package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)
	total, priority := 0, 0

	for scan.Scan() {

		var sets = [3]map[int]string{}

		for i := 0; i < 3; i++ {
			line := scan.Text()

			sets[i] = map[int]string{}

			for _, c := range line {
				sets[i][int(c)] = string(c)
			}

			if i < 2 {
				scan.Scan()
			}
		}

		for key := range sets[0] {
			_, e1 := sets[1][key]
			_, e2 := sets[2][key]
			if e1 && e2 {
				// fmt.Println("Key := " + string(key))
				if key <= 90 { // uppercase
					priority = key - 64 + 26
				} else { // lowercase
					priority = key - 96
				}
				total += priority
			}
		}
	}

	fmt.Println(total)
}
