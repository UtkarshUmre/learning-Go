package main

import "fmt"

func main() {
	fmt.Println("functions in go lang")

	greet()
	add := adder(9, 1)
	fmt.Println(add)
}

func adder(x int, y int) int {
	return x + y
}

func greet() {
	fmt.Println("hello how are you ")
}
