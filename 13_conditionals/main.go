package main

import "fmt"

func main() {
	fmt.Println("conditionals in go")

	loginCount := 25
	var result string

	if loginCount < 10 {
		result = "regular user "
	} else {
		result = "not regular user"
	}

	fmt.Println(result)
}
