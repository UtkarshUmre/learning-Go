package main

import "fmt"

func main() {
	fmt.Println("function in go lang")

	greet() //calling the function

	result := add(9, 1)
	fmt.Println(result)
}

func greet() { //initializing the funcion in go programming language
	fmt.Println("namaste from go lang")
}

func add(x int, y int) int {
	return x + y

}
