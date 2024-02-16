package main

import "fmt"

func main() {

	loginCount := 15

	var result string

	if loginCount < 10 {
		result = " regular user"
	} else if loginCount > 20 {
		result = "non regular user"
	} else {
		result = "watch out "
	}

	fmt.Println(result)
}
