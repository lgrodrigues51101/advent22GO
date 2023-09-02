package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)

	total := 0

	for scan.Scan() {
		line := scan.Text()

		elf_one := strings.Split(line, ",")[0]
		elf_two := strings.Split(line, ",")[1]

		section_one, _ := strconv.Atoi(strings.Split(elf_one, "-")[0])
		section_two, _ := strconv.Atoi(strings.Split(elf_one, "-")[1])

		section_three, _ := strconv.Atoi(strings.Split(elf_two, "-")[0])
		section_four, _ := strconv.Atoi(strings.Split(elf_two, "-")[1])

		if section_three <= section_two && section_four >= section_one {
			total += 1
		}
	}

	fmt.Println(total)
}
