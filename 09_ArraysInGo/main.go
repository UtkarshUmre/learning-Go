package main

import "fmt"

func main() {
	fmt.Println("array in go lang")

	// var furnitureList [4]string

	// furnitureList[0] = "table"
	// furnitureList[1] = "chair"
	// furnitureList[2] = "sofa"
	// furnitureList[3] = "cupboard"

	var furnitureList = [3]string{"sofa", "chair", "talbe"}

	fmt.Println("furniture  list is ", furnitureList)
	fmt.Printf("furniture list is of type %T\n", furnitureList)
	fmt.Println("furniture list if of lenght:", len(furnitureList))

}
