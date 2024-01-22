package main

import (
	"fmt"
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

}
