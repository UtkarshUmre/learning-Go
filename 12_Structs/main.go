package main

import "fmt"

func main() {
	fmt.Println("welcome to structs in go lang")
	// no inheretance in go lang;no super or parent no such concepts in go lang

	utkarsh := User{"utkarsh", "utkarsh@mlh.io", true, 21}
	fmt.Println(utkarsh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
