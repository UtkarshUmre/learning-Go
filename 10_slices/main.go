package main

import "fmt"

func main() {
	fmt.Println("slices in go lang")

	var groceryList = []string{"oats", "ketchup", "vegetables", "creatine"}
	groceryList = append(groceryList, "ice-cream", "bread", "juice")
	fmt.Println("grocery list is", groceryList)
	fmt.Printf("type of groceryList is %T\n", groceryList)
	fmt.Println("length of grocery list is ", len(groceryList))

	// fmt.Println("appended grocery list is ", groceryList)

}
