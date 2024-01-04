Understanding Structs in Go

Structs are a fundamental building block in Go, enabling the creation of complex data structures that can hold different types of data.
What is a Struct?

A struct is a composite data type that groups together variables of different types under a single name. It's a way to define custom data types by combining individual data items into a larger structure.
Creating a Struct

To define a struct in Go, use the type keyword followed by the struct's name and its fields:

go

type Person struct {
    FirstName string
    LastName  string
    Age       int
}

This Person struct has fields for FirstName, LastName, and Age.
Initializing a Struct

You can create instances of a struct by assigning values to its fields:

go

person := Person{
    FirstName: "John",
    LastName:  "Doe",
    Age:       30,
}

Accessing Struct Fields

You can access individual fields of a struct using dot notation:

go

fmt.Println("First Name:", person.FirstName)
fmt.Println("Last Name:", person.LastName)
fmt.Println("Age:", person.Age)

Updating Struct Fields

Struct fields can be updated by reassigning new values:

go

person.Age = 31

Embedding Structs

Go supports embedding one struct within another to create more complex structures:

go

type Employee struct {
    Person // Embedded Person struct
    JobTitle string
    Salary   float64
}

Structs and Methods

Go allows defining methods for structs:

go

func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}

Conclusion

Structs are a versatile feature in Go, allowing developers to create custom data types with various fields and methods. They provide a way to organize and manipulate data efficiently, contributing to the language's simplicity and power.
