package main

import "fmt"

func main() {
	fmt.Println("welcome to structs in go lang")
	// no inheretance in go lang;no super or parent no such concepts in go lang

	utkarsh := User{"utkarsh", "utkarsh@mlh.io", true, 21}
	fmt.Println(utkarsh)
	fmt.Printf("utkarsh details are %+v\n", utkarsh)
	fmt.Printf("name is %v and email is %v\n", utkarsh.Name, utkarsh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
