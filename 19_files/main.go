package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("lets look how to handle files in go lang")
	content := "the content written in this string needs to go in the file in order to learn file handling in go language"
	file, err := os.Create("./sampletxtfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	file.Close()

}
