package main

import "fmt"

func main() {

	fmt.Println("struct in go")
	Nathan := User{"Nathan", "Nathan@gmail", 23, true}
	fmt.Printf("nathans details are %+v\n", Nathan)
	fmt.Printf("name is %v and emai is %v \n", Nathan.Name, Nathan.Email)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
