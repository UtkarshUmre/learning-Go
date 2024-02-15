package main

import "fmt"

func main() {

	fmt.Println("struct in go")
	Nathan := User{"Nathan", "Nathan@gmail", 23, true}
	fmt.Println(Nathan)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
