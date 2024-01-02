package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	if len(s.items) == 0 {
		return 0 // Returning 0 for an empty stack (you can change this handling)
	}
	topIndex := len(s.items) - 1
	item := s.items[topIndex]
	s.items = s.items[:topIndex]
	return item
}

func reverseString(input string) string {
	stack := Stack{}
	for _, char := range input {
		stack.Push(char)
	}

	var reversed strings.Builder
	for len(stack.items) > 0 {
		reversed.WriteRune(stack.Pop())
	}

	return reversed.String()
}

func main() {
	input := "edoCfOsyaD563relacs"
	reversed := reverseString(input)
	fmt.Println("Reversed:", reversed)
}
