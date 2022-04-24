package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("hello , arnab")
	// PerformGet()
	// PerformPost()
	PerformPostForm()
}

func PerformGet() {
	const url = "http://localhost:8080"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("status :", res.StatusCode)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("byte count is : ", byteCount)

	fmt.Println(responseString.String())

	fmt.Println("my content :: ", string(content))
}

func PerformPost() {

	const myURl = "http://localhost:8080/post"

	requestBody := strings.NewReader(`
		{
			"name" : "arnab",
			"age" : "24"
		}
	`)

	response, err := http.Post(myURl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	resContent, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(resContent))
}

func PerformPostForm() {
	const myURl = "http://localhost:8080/postform"

	// formdata
	data := url.Values{}

	data.Add("Firstname", "arnab")
	data.Add("Lastname", "saha")

	res, err := http.PostForm(myURl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
}
