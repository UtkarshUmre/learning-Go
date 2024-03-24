package main

import "fmt"

func main() {

	fmt.Println("functions in go")
	greeter()

	result := adder(2, 7)
	fmt.Println("Result is:", result)

}

func greeter() {
	fmt.Println("hello all")
}

func adder(valOne int, valTwo int) int {

	return valOne + valTwo

}

func proAdder(values ...int) {

}
