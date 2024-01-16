package main

import "fmt"

func main() {
	// fmt.Println("pointers in go lang")

	// 	var ptr *int
	// 	fmt.Println(ptr)
	// 	fmt.Printf("type of ptr is: %T\n", ptr)

	myNumber := 45
	var ptr = &myNumber
	fmt.Println("value of actual pointer is ", ptr)
	fmt.Printf("type of ptr is: %T\n", ptr)

	*ptr = *ptr + 2
	fmt.Println(*ptr)

}
