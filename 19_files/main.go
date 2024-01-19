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

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./sampletxtfile.txt")

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
