package main

import (
	"fmt"
)

func main() {
	fmt.Print("please enter your name: ")
	greet()

	Result := adder(4, 7)
	fmt.Println("added value is:", Result)

}

func adder(valOne int, valTwo int) int {
	add := valOne + valTwo

	return add
}

func greet() {
	var input string
	fmt.Scan(&input)
	fmt.Println("hello", input)
}
