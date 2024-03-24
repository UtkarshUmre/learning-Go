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
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./file.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("text data inside the file is: \n", string(databyte))

}

func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}

}
