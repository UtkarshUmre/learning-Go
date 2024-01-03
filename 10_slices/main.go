package main

import "fmt"

func main() {

	fmt.Println("slices in go lang")

	//initialization and declaration of slices
	var fruitList = []string{"tomato", "apple", "mango", "guava", "banana", "strawberry", "kiwi", "papaya"}
	fmt.Println(fruitList)
	fmt.Println("fruitlist length is", len(fruitList))
	fmt.Printf("fruitlist is of type %T\n", fruitList)

	fruitList = append(fruitList, "pomegranate", "dragonfruit", "coconut")
	fmt.Println(fruitList)

}
