package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("webrequest in go lang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)

	}

	fmt.Printf("type of response is %T\n", response)
	// fmt.Println(response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}
