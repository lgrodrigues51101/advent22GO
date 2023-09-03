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

	scan.Scan()
	line := scan.Text()

	const size = 14
	var buffer [size]rune

	for index, c := range line {

		for i := 1; i < size; i++ {
			buffer[i-1] = buffer[i]
		}

		buffer[size-1] = c

		if index >= size-1 {
			alldiff := true

			for i := 0; i < size; i++ {
				for j := i + 1; j < size; j++ {
					if buffer[i] == buffer[j] {
						alldiff = false
					}
				}
			}

			if alldiff {
				fmt.Println(buffer)
				fmt.Println(index + 1)
				break
			}
		}
	}
}
