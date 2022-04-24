package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/20"

func main() {

	fmt.Println("hi go")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Res is of type %T\n", res)

	defer res.Body.Close()

	fmt.Println("status code ", res.Status, res.StatusCode)
	dataBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)
}
