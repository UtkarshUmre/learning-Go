package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("lets see how to handle files in go lang")
	content := "this needs to go in a file"

	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()
}
