package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack represents a stack data structure.
type Stack struct {
	items []interface{}
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the item from the top of the stack without removing it.
func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Reverse reverses the order of elements in the stack.
func (s *Stack) Reverse() {
	reversedStack := Stack{}
	for !s.IsEmpty() {
		item, _ := s.Pop()
		reversedStack.Push(item)
	}
	s.items = reversedStack.items
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scan := bufio.NewScanner(file)

	quant := 9

	stacks := make([]Stack, quant)

	for scan.Scan() {
		line := scan.Text()

		if line == "" { // moves will start
			break
		}

		for i := 0; i < quant; i++ {
			if line[i*3+i] == ' ' {
				continue
			}
			letter := string(line[i*3+i+1])

			stacks[i].Push(letter)
		}
	}

	for i := 0; i < quant; i++ {
		stacks[i].Reverse()
	}

	for scan.Scan() {
		line := scan.Text()

		tokens := strings.Split(line, " ")
		from, _ := strconv.Atoi(tokens[3])
		quant, _ := strconv.Atoi(tokens[1])
		to, _ := strconv.Atoi(tokens[5])

		temp := Stack{}

		for i := 0; i < quant; i++ {
			elem, _ := stacks[from-1].Pop()
			temp.Push(elem)
		}

		for i := 0; i < quant; i++ {
			elem, _ := temp.Pop()
			stacks[to-1].Push(elem)
		}

	}

	for i := 0; i < quant; i++ {
		letter, _ := stacks[i].Peek()
		fmt.Print(letter)
	}
}
