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
	fmt.Println("fruit list is of length", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)
	fmt.Println("vegList length is", len(vegList))

}
