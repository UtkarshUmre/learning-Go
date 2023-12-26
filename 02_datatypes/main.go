package main

import "fmt"

const LoginToken string = "gbddhjklm"

func main() {
	fmt.Println(LoginToken)
	fmt.Println("variables, constants and datatypes in go programming")

	//string datatype
	var username string = "utkarsh"
	//to print the variable
	fmt.Println(username)
	//to print the datatype of variable
	fmt.Printf("variable is of type: %T \n", username)

	//boolean datatype
	var isLoggedIn bool = true
	//to print the variable
	fmt.Println(isLoggedIn)
	//to print the datatype of variable
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	//small value
	var small uint8 = 255
	//printing the variable
	fmt.Println(small)
	//Printing the datatype of variable
	fmt.Printf("value is of type: %T \n", small)

	//small float vlue
	var smallFloat float64 = 255.4553242456209
	//printing the variable
	fmt.Println(smallFloat)
	//printing the datatype of variable
	fmt.Printf("value is of tyepe: %T \n", smallFloat)

	//default vlaues and some aliases
	// var anotherVariable int
	anotherVariable := 568
	fmt.Println(anotherVariable)
	fmt.Printf("anotherVariable is of tyep: %T \n", anotherVariable)

	//implicit type
	var website = "linktr.ee/utkarshumre"
	fmt.Println(website)
	fmt.Printf("type of website : %T \n", website)

	//no var style
	numberOfUsers := 43
	fmt.Println(numberOfUsers)
	fmt.Printf("type of numberOfUsers: %T \n", numberOfUsers)

}
