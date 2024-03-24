package main

import "fmt"

func main() {

	fmt.Println("functions in go")
	greeter()

	result := adder(2, 7)
	fmt.Println("Result is:", result)

	proResult := proAdder(1, 2, 5, 3, 9, 4)
	fmt.Println("result is:", proResult)

}

func greeter() {
	fmt.Println("hello all")
}

func adder(valOne int, valTwo int) int {

	return valOne + valTwo

}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total

}
