package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main()  {
	
	file, _ := os.Open("input")
	defer file.Close()

	var elfs[] int

	scan := bufio.NewScanner(file)
	
	calories := 0
	for scan.Scan() {
		line := scan.Text()
		
		if line == "" {
			elfs = append(elfs, calories)
			calories = 0
		} else {
			n, _ := strconv.Atoi(line)
			calories += n
		}
	}

	sort.Ints(elfs)
	size:=len(elfs)

	i, total := 1,0
	for i < 4 {
		total = total + elfs[size-i]
		i++
	}

	fmt.Println(total)
}