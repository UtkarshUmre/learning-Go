# Accepting User Input in Go

User input is an essential aspect of many programs as it enables interaction and customization. In Go, handling user input is straightforward and can be done using the `fmt` package from the standard library.

## Using `fmt.Scanln()`

The `fmt.Scanln()` function allows us to accept input from the user via the console. It reads input until a newline (`Enter`) character is encountered.

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Println("Enter your name:")
    fmt.Scanln(&name)
    fmt.Printf("Hello, %s!\n", name)
}
```

Here, we prompt the user to enter their name. The fmt.Scanln(&name) line reads the user's input and stores it in the name variable.

## Reading Numbers with fmt.Scanln()

To read numeric input, the approach remains similar:

```
package main

import "fmt"

func main() {
    var num int
    fmt.Println("Enter a number:")
    fmt.Scanln(&num)
    fmt.Printf("You entered: %d\n", num)
}
```

This code snippet reads an integer input from the user and displays it back.

## Error Handling

While fmt.Scanln() is convenient, it's essential to handle potential errors, especially when dealing with user input. For instance:

```
package main

import "fmt"

func main() {
    var age int
    fmt.Println("Enter your age:")
    _, err := fmt.Scanln(&age)
    if err != nil {
        fmt.Println("Error occurred:", err)
        return
    }
    fmt.Printf("Your age is: %d\n", age)
}
```

## Conclusion

Handling user input in Go is straightforward using the fmt package. By using functions like fmt.Scanln(), we can interactively accept user input and build more interactive programs.
