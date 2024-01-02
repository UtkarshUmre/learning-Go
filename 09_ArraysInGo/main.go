package main

import "fmt"

func main() {
	fmt.Println("arrays in go lang")

	var fruitList [5]string

	fruitList[0] = "mango"
	fruitList[1] = "apple"
	fruitList[2] = "banana"
	fruitList[3] = "guava"
	fruitList[4] = "jackfruit"

	fmt.Println("fruit list is", fruitList)
}
