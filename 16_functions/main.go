package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	fmt.Print("enter daddy here: ")
	daddy()
}

func daddy() {
	var input string
	fmt.Scan(&input)
	fmt.Println("bloody", input)
}
