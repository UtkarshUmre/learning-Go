package main

import "fmt"

func main() {

	fmt.Println("welcome to loops in golang")

	days := []string{"sunday", "tuesday", "thursday", "saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	for i := range days {
		fmt.Println(days[i])
	}

}
